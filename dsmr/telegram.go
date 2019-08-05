// Package dsmr handles dsmr specific things
package dsmr

// Telegram can store contents of a p1 telegram
type Telegram struct {
	Identification          string  `json:"identification"`
	P1Version               string  `json:"p1_version"`
	Timestamp               string  `json:"timestamp"`
	EquipmentID             string  `json:"equipment_id"`
	EnergyDeliveredTariff1  float64 `json:"energy_delivered_tariff1"`
	EnergyDeliveredTariff2  float64 `json:"energy_delivered_tariff2"`
	EnergyReturnedTariff1   int     `json:"energy_returned_tariff1"`
	EnergyReturnedTariff2   int     `json:"energy_returned_tariff2"`
	ElectricityTariff       string  `json:"electricity_tariff"`
	PowerDelivered          float64 `json:"power_delivered"`
	PowerReturned           int     `json:"power_returned"`
	ElectricityFailures     int     `json:"electricity_failures"`
	ElectricityLongFailures int     `json:"electricity_long_failures"`
	ElectricityFailureLog   string  `json:"electricity_failure_log"`
	ElectricitySagsL1       int     `json:"electricity_sags_l1"`
	ElectricitySwellsL1     int     `json:"electricity_swells_l1"`
	MessageLong             string  `json:"message_long"`
	VoltageL1               float64 `json:"voltage_l1"`
	CurrentL1               int     `json:"current_l1"`
	PowerDeliveredL1        float64 `json:"power_delivered_l1"`
	PowerReturnedL1         int     `json:"power_returned_l1"`
	GasDeviceType           int     `json:"gas_device_type"`
	GasEquipmentID          string  `json:"gas_equipment_id"`
	GasDelivered            float64 `json:"gas_delivered"`
}
