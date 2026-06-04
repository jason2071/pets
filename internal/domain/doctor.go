package domain

import "time"

// Doctor is a veterinarian working at a clinic.
type Doctor struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ClinicID  uint      `gorm:"not null;index" json:"clinic_id"`
	Clinic    *Clinic   `gorm:"constraint:OnDelete:CASCADE" json:"clinic,omitempty"`
	FullName  string    `gorm:"not null" json:"full_name"`
	Specialty string    `gorm:"index" json:"specialty,omitempty"`
	LicenseNo string    `json:"license_no,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
