package main

import (
	"log"

	"github.com/jason2071/pets/internal/config"
	"github.com/jason2071/pets/internal/database"
	"github.com/jason2071/pets/internal/domain"
	"github.com/jason2071/pets/internal/handler"
	"github.com/jason2071/pets/internal/repository"
	"github.com/jason2071/pets/internal/server"
	"github.com/jason2071/pets/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("database connect: %v", err)
	}

	// Hybrid migration strategy:
	//   - dev:  AUTO_MIGRATE=true  -> GORM AutoMigrate shapes the schema from structs
	//   - prod: AUTO_MIGRATE=false -> schema is owned by SQL migrations (make migrate-up)
	// Never let both manage the same table at once.
	if cfg.AutoMigrate {
		// Order matters for foreign keys: parents (accounts, clinics) before dependents.
		if err := db.AutoMigrate(
			&domain.Account{},
			&domain.Clinic{},
			&domain.Doctor{},
			&domain.Pet{},
		); err != nil {
			log.Fatalf("auto migrate: %v", err)
		}
		log.Println("schema: GORM AutoMigrate applied")
	} else {
		log.Println("schema: AutoMigrate disabled, expecting SQL migrations")
	}

	petRepo := repository.NewPetRepository(db)
	petSvc := service.NewPetService(petRepo)
	petHandler := handler.NewPetHandler(petSvc)

	accRepo := repository.NewAccountRepository(db)
	accSvc := service.NewAccountService(accRepo)
	accHandler := handler.NewAccountHandler(accSvc)

	r := server.NewRouter(petHandler, accHandler)

	addr := ":" + cfg.Port
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server run: %v", err)
	}
}
