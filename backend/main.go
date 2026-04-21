package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Lens struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	FocalLength    string  `json:"focalLength"`
	MaxAperture    float64 `json:"maxAperture"`
	MinAperture    float64 `json:"minAperture"`
	ImageStab      bool    `json:"imageStab"`
	Weight         int     `json:"weight"`
	Price          float64 `json:"price"`
	Type           string  `json:"type"`
	ImageSide      string  `json:"imageSide"`
	ImageTop       string  `json:"imageTop"`
	ImageIsometric string  `json:"imageIsometric"`
	Length         float64 `json:"length"`         // Length in mm
	Diameter       float64 `json:"diameter"`       // Diameter in mm
	WeatherSealing string  `json:"weatherSealing"` // Weather sealing level: dust, drip, rain, sealed, submersible
}

type Car struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`           // Full name (computed or stored)
	Model          string  `json:"model"`          // Base model name (e.g., "Camry", "Taycan", "Air")
	Trim           string  `json:"trim"`           // Trim level (e.g., "Pure", "Turbo S", "1500")
	Variant        string  `json:"variant"`        // Variant (optional, e.g., body style variant)
	Manufacturer   string  `json:"manufacturer"`
	Year           int     `json:"year"`
	Type           string  `json:"type"`           // Sedan, SUV, Truck, etc.
	BodyStyle      string  `json:"bodyStyle"`      // Coupe, Hatchback, Convertible, etc.
	Engine         string  `json:"engine"`         // Engine description
	Horsepower     int     `json:"horsepower"`
	Torque         int     `json:"torque"`         // ft-lb or Nm
	Transmission   string  `json:"transmission"`   // Automatic, Manual, CVT
	Drivetrain     string  `json:"drivetrain"`     // FWD, RWD, AWD, 4WD
	MPGCity        float64 `json:"mpgCity"`
	MPGHighway     float64 `json:"mpgHighway"`
	MPGCombined    float64 `json:"mpgCombined"`
	FuelCapacity   float64 `json:"fuelCapacity"`   // gallons or liters
	BatteryCapacity float64 `json:"batteryCapacity"` // kWh for electric vehicles
	Seating        int     `json:"seating"`
	CargoCapacity  float64 `json:"cargoCapacity"`  // cubic feet or liters
	Length         float64 `json:"length"`         // inches or cm
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	Wheelbase      float64 `json:"wheelbase"`
	Weight         int     `json:"weight"`         // lbs or kg
	Price          float64 `json:"price"`
	ImageFront     string  `json:"imageFront"`
	ImageSide      string  `json:"imageSide"`
	ImageInterior  string  `json:"imageInterior"`
	ImageExterior  string  `json:"imageExterior"`
	TopSpeed       int     `json:"topSpeed"`       // mph or km/h
	ZeroToSixty    float64 `json:"zeroToSixty"`    // seconds
	SafetyRating   string  `json:"safetyRating"`   // NHTSA/IIHS rating
	FuelType       string  `json:"fuelType"`       // Gasoline, Electric, Hybrid, Diesel, Plug-in Hybrid
}

type Excavator struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Model           string  `json:"model"`
	Manufacturer    string  `json:"manufacturer"`
	Scale           string  `json:"scale"`           // 1:14, 1:16, 1:20, etc.
	Year            int     `json:"year"`
	Weight          float64 `json:"weight"`         // kg or lbs
	Length          float64 `json:"length"`         // cm or inches
	Width           float64 `json:"width"`
	Height          float64 `json:"height"`
	BatteryType     string  `json:"batteryType"`    // LiPo, NiMH, etc.
	BatteryCapacity float64 `json:"batteryCapacity"` // mAh
	OperatingTime   float64 `json:"operatingTime"`  // minutes
	ControlChannels int     `json:"controlChannels"` // number of channels
	DiggingDepth    float64 `json:"diggingDepth"`   // cm or inches
	LiftingCapacity float64 `json:"liftingCapacity"` // kg or lbs
	RotationAngle   int     `json:"rotationAngle"`  // degrees (usually 360)
	Material         string  `json:"material"`      // Metal, Plastic, etc.
	RemoteControl   string  `json:"remoteControl"`  // 2.4GHz, etc.
	Functions       string  `json:"functions"`      // List of functions
	Price           float64 `json:"price"`
	ImageFront      string  `json:"imageFront"`
	ImageSide       string  `json:"imageSide"`
	ImageTop        string  `json:"imageTop"`
	ImageDetail     string  `json:"imageDetail"`
}

var lenses = []Lens{
	// Ultra Wide Zoom Lenses
	{ID: "1", Name: "NIKKOR Z 14-30mm f/4 S", FocalLength: "14-30mm", MaxAperture: 4.0, MinAperture: 22.0, ImageStab: false, Weight: 485, Price: 1299.95, Type: "Ultra Wide Angle Zoom", ImageSide: "https://via.placeholder.com/300x85/667eea/ffffff", ImageTop: "https://via.placeholder.com/89x89/667eea/ffffff", ImageIsometric: "", Length: 85.0, Diameter: 89.0, WeatherSealing: "sealed"},
	{ID: "11", Name: "NIKKOR Z 14-24mm f/2.8 S", FocalLength: "14-24mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 650, Price: 2399.95, Type: "Ultra Wide Angle Zoom", ImageSide: "https://via.placeholder.com/300x125/764ba2/ffffff", ImageTop: "https://via.placeholder.com/98x98/764ba2/ffffff", ImageIsometric: "", Length: 125.0, Diameter: 98.0, WeatherSealing: "sealed"},

	// Standard Zoom Lenses
	{ID: "2", Name: "NIKKOR Z 24-70mm f/4 S", FocalLength: "24-70mm", MaxAperture: 4.0, MinAperture: 22.0, ImageStab: false, Weight: 500, Price: 999.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/300x88/667eea/ffffff", ImageTop: "https://via.placeholder.com/77x77/667eea/ffffff", ImageIsometric: "", Length: 88.5, Diameter: 77.5, WeatherSealing: "rain"},
	{ID: "3", Name: "NIKKOR Z 24-70mm f/2.8 S", FocalLength: "24-70mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 805, Price: 2399.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/300x126/764ba2/ffffff", ImageTop: "https://via.placeholder.com/89x89/764ba2/ffffff", ImageIsometric: "", Length: 126.0, Diameter: 89.0, WeatherSealing: "sealed"},
	{ID: "10", Name: "NIKKOR Z 24-120mm f/4 S", FocalLength: "24-120mm", MaxAperture: 4.0, MinAperture: 22.0, ImageStab: false, Weight: 630, Price: 1099.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/350x118/667eea/ffffff", ImageTop: "https://via.placeholder.com/84x84/667eea/ffffff", ImageIsometric: "", Length: 118.0, Diameter: 84.0, WeatherSealing: "sealed"},
	{ID: "12", Name: "NIKKOR Z 24-200mm f/4-6.3 VR", FocalLength: "24-200mm", MaxAperture: 4.0, MinAperture: 22.0, ImageStab: true, Weight: 570, Price: 899.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/350x114/667eea/ffffff", ImageTop: "https://via.placeholder.com/78x78/667eea/ffffff", ImageIsometric: "", Length: 114.0, Diameter: 76.5, WeatherSealing: "sealed"},
	{ID: "13", Name: "NIKKOR Z 28-75mm f/2.8", FocalLength: "28-75mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 565, Price: 1199.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/300x120/667eea/ffffff", ImageTop: "https://via.placeholder.com/75x75/667eea/ffffff", ImageIsometric: "", Length: 120.5, Diameter: 75.0, WeatherSealing: "sealed"},

	// Telephoto Zoom Lenses
	{ID: "4", Name: "NIKKOR Z 70-200mm f/2.8 VR S", FocalLength: "70-200mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: true, Weight: 1440, Price: 2699.95, Type: "Telephoto Zoom", ImageSide: "https://via.placeholder.com/400x220/667eea/ffffff", ImageTop: "https://via.placeholder.com/89x89/667eea/ffffff", ImageIsometric: "", Length: 220.0, Diameter: 89.0, WeatherSealing: "sealed"},
	{ID: "14", Name: "NIKKOR Z 70-180mm f/2.8", FocalLength: "70-180mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 795, Price: 1249.95, Type: "Telephoto Zoom", ImageSide: "https://via.placeholder.com/350x151/667eea/ffffff", ImageTop: "https://via.placeholder.com/83x83/667eea/ffffff", ImageIsometric: "", Length: 151.0, Diameter: 83.0, WeatherSealing: "sealed"},

	// Super Telephoto Zoom Lenses
	{ID: "5", Name: "NIKKOR Z 100-400mm f/4.5-5.6 VR S", FocalLength: "100-400mm", MaxAperture: 4.5, MinAperture: 32.0, ImageStab: true, Weight: 1435, Price: 2699.95, Type: "Super Telephoto Zoom", ImageSide: "https://via.placeholder.com/400x222/764ba2/ffffff", ImageTop: "https://via.placeholder.com/98x98/764ba2/ffffff", ImageIsometric: "", Length: 222.0, Diameter: 98.0, WeatherSealing: "sealed"},
	{ID: "15", Name: "NIKKOR Z 180-600mm f/5.6-6.3 VR", FocalLength: "180-600mm", MaxAperture: 5.6, MinAperture: 32.0, ImageStab: true, Weight: 2140, Price: 1699.95, Type: "Super Telephoto Zoom", ImageSide: "https://via.placeholder.com/450x315/764ba2/ffffff", ImageTop: "https://via.placeholder.com/110x110/764ba2/ffffff", ImageIsometric: "", Length: 315.5, Diameter: 110.0, WeatherSealing: "sealed"},

	// Prime Lenses - Ultra Wide
	{ID: "9", Name: "NIKKOR Z 20mm f/1.8 S", FocalLength: "20mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 505, Price: 996.95, Type: "Ultra Wide Prime", ImageSide: "https://via.placeholder.com/300x108/667eea/ffffff", ImageTop: "https://via.placeholder.com/84x84/667eea/ffffff", ImageIsometric: "", Length: 108.5, Diameter: 84.5, WeatherSealing: "sealed"},
	{ID: "16", Name: "NIKKOR Z 24mm f/1.8 S", FocalLength: "24mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 450, Price: 996.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x103/667eea/ffffff", ImageTop: "https://via.placeholder.com/72x72/667eea/ffffff", ImageIsometric: "", Length: 103.0, Diameter: 72.0, WeatherSealing: "sealed"},
	{ID: "17", Name: "NIKKOR Z 26mm f/2.8", FocalLength: "26mm", MaxAperture: 2.8, MinAperture: 16.0, ImageStab: false, Weight: 125, Price: 496.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x23/667eea/ffffff", ImageTop: "https://via.placeholder.com/70x70/667eea/ffffff", ImageIsometric: "", Length: 23.5, Diameter: 70.0, WeatherSealing: "dust"},
	{ID: "18", Name: "NIKKOR Z 28mm f/2.8 SE", FocalLength: "28mm", MaxAperture: 2.8, MinAperture: 16.0, ImageStab: false, Weight: 160, Price: 296.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x43/667eea/ffffff", ImageTop: "https://via.placeholder.com/72x72/667eea/ffffff", ImageIsometric: "", Length: 43.0, Diameter: 72.0, WeatherSealing: "dust"},

	// Prime Lenses - Standard
	{ID: "8", Name: "NIKKOR Z 35mm f/1.8 S", FocalLength: "35mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 370, Price: 846.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x86/667eea/ffffff", ImageTop: "https://via.placeholder.com/73x73/667eea/ffffff", ImageIsometric: "", Length: 86.0, Diameter: 73.0, WeatherSealing: "sealed"},
	{ID: "19", Name: "NIKKOR Z 40mm f/2", FocalLength: "40mm", MaxAperture: 2.0, MinAperture: 16.0, ImageStab: false, Weight: 170, Price: 296.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x46/667eea/ffffff", ImageTop: "https://via.placeholder.com/70x70/667eea/ffffff", ImageIsometric: "", Length: 45.5, Diameter: 70.0, WeatherSealing: "drip"},
	{ID: "6", Name: "NIKKOR Z 50mm f/1.8 S", FocalLength: "50mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 415, Price: 496.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x86/667eea/ffffff", ImageTop: "https://via.placeholder.com/76x76/667eea/ffffff", ImageIsometric: "", Length: 86.5, Diameter: 76.0, WeatherSealing: "drip"},
	{ID: "20", Name: "NIKKOR Z 50mm f/1.2 S", FocalLength: "50mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 1090, Price: 2099.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x150/764ba2/ffffff", ImageTop: "https://via.placeholder.com/108x108/764ba2/ffffff", ImageIsometric: "", Length: 150.0, Diameter: 108.0, WeatherSealing: "sealed"},
	{ID: "21", Name: "NIKKOR Z 58mm f/0.95 S Noct", FocalLength: "58mm", MaxAperture: 0.95, MinAperture: 16.0, ImageStab: false, Weight: 2000, Price: 7999.95, Type: "Prime", ImageSide: "https://via.placeholder.com/350x153/764ba2/ffffff", ImageTop: "https://via.placeholder.com/103x103/764ba2/ffffff", ImageIsometric: "", Length: 153.0, Diameter: 103.0, WeatherSealing: "sealed"},

	// Prime Lenses - Portrait/Telephoto
	{ID: "7", Name: "NIKKOR Z 85mm f/1.8 S", FocalLength: "85mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 470, Price: 796.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x99/667eea/ffffff", ImageTop: "https://via.placeholder.com/75x75/667eea/ffffff", ImageIsometric: "", Length: 99.0, Diameter: 75.0, WeatherSealing: "rain"},
	{ID: "22", Name: "NIKKOR Z 85mm f/1.2 S", FocalLength: "85mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 1160, Price: 2799.95, Type: "Prime", ImageSide: "https://via.placeholder.com/350x142/764ba2/ffffff", ImageTop: "https://via.placeholder.com/102x102/764ba2/ffffff", ImageIsometric: "", Length: 142.0, Diameter: 102.0, WeatherSealing: "sealed"},
	{ID: "23", Name: "NIKKOR Z 105mm f/2.8 VR S Macro", FocalLength: "105mm", MaxAperture: 2.8, MinAperture: 32.0, ImageStab: true, Weight: 630, Price: 999.95, Type: "Macro Prime", ImageSide: "https://via.placeholder.com/300x140/667eea/ffffff", ImageTop: "https://via.placeholder.com/75x75/667eea/ffffff", ImageIsometric: "", Length: 140.0, Diameter: 75.0, WeatherSealing: "sealed"},
	{ID: "24", Name: "NIKKOR Z 135mm f/1.8 S Plena", FocalLength: "135mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 995, Price: 2499.95, Type: "Telephoto Prime", ImageSide: "https://via.placeholder.com/350x139/764ba2/ffffff", ImageTop: "https://via.placeholder.com/89x89/764ba2/ffffff", ImageIsometric: "", Length: 139.5, Diameter: 89.0, WeatherSealing: "sealed"},

	// Super Telephoto Prime Lenses
	{ID: "25", Name: "NIKKOR Z 400mm f/2.8 TC VR S", FocalLength: "400mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: true, Weight: 2950, Price: 13999.95, Type: "Super Telephoto Prime", ImageSide: "https://via.placeholder.com/450x380/764ba2/ffffff", ImageTop: "https://via.placeholder.com/160x160/764ba2/ffffff", ImageIsometric: "", Length: 380.0, Diameter: 160.0, WeatherSealing: "sealed"},
	{ID: "26", Name: "NIKKOR Z 400mm f/4.5 VR S", FocalLength: "400mm", MaxAperture: 4.5, MinAperture: 32.0, ImageStab: true, Weight: 1245, Price: 3299.95, Type: "Super Telephoto Prime", ImageSide: "https://via.placeholder.com/400x234/764ba2/ffffff", ImageTop: "https://via.placeholder.com/104x104/764ba2/ffffff", ImageIsometric: "", Length: 234.5, Diameter: 104.0, WeatherSealing: "sealed"},
	{ID: "27", Name: "NIKKOR Z 600mm f/4 TC VR S", FocalLength: "600mm", MaxAperture: 4.0, MinAperture: 22.0, ImageStab: true, Weight: 3260, Price: 15499.95, Type: "Super Telephoto Prime", ImageSide: "https://via.placeholder.com/500x437/764ba2/ffffff", ImageTop: "https://via.placeholder.com/168x168/764ba2/ffffff", ImageIsometric: "", Length: 437.0, Diameter: 168.0, WeatherSealing: "sealed"},
	{ID: "28", Name: "NIKKOR Z 600mm f/6.3 VR S PF", FocalLength: "600mm", MaxAperture: 6.3, MinAperture: 32.0, ImageStab: true, Weight: 1470, Price: 4799.95, Type: "Super Telephoto Prime", ImageSide: "https://via.placeholder.com/400x278/764ba2/ffffff", ImageTop: "https://via.placeholder.com/106x106/764ba2/ffffff", ImageIsometric: "", Length: 278.0, Diameter: 106.0, WeatherSealing: "sealed"},
	{ID: "29", Name: "NIKKOR Z 800mm f/6.3 VR S PF", FocalLength: "800mm", MaxAperture: 6.3, MinAperture: 32.0, ImageStab: true, Weight: 2385, Price: 6499.95, Type: "Super Telephoto Prime", ImageSide: "https://via.placeholder.com/450x385/764ba2/ffffff", ImageTop: "https://via.placeholder.com/140x140/764ba2/ffffff", ImageIsometric: "", Length: 385.0, Diameter: 140.0, WeatherSealing: "sealed"},

	// Third-Party Z Mount Lenses - Tamron
	{ID: "30", Name: "Tamron 28-75mm f/2.8 Di III VXD G2", FocalLength: "28-75mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 540, Price: 899.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/300x117/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/73x73/4a90e2/ffffff", ImageIsometric: "", Length: 117.0, Diameter: 73.0, WeatherSealing: "sealed"},
	{ID: "31", Name: "Tamron 70-180mm f/2.8 Di III VC VXD G2", FocalLength: "70-180mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: true, Weight: 855, Price: 1199.95, Type: "Telephoto Zoom", ImageSide: "https://via.placeholder.com/350x156/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/81x81/4a90e2/ffffff", ImageIsometric: "", Length: 156.0, Diameter: 81.0, WeatherSealing: "sealed"},
	{ID: "32", Name: "Tamron 35-150mm f/2-2.8 Di III VXD", FocalLength: "35-150mm", MaxAperture: 2.0, MinAperture: 22.0, ImageStab: false, Weight: 1165, Price: 1899.95, Type: "Standard Zoom", ImageSide: "https://via.placeholder.com/350x158/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/89x89/4a90e2/ffffff", ImageIsometric: "", Length: 158.0, Diameter: 89.0, WeatherSealing: "sealed"},
	{ID: "33", Name: "Tamron 50-400mm f/4.5-6.3 Di III VC VXD", FocalLength: "50-400mm", MaxAperture: 4.5, MinAperture: 32.0, ImageStab: true, Weight: 1155, Price: 1299.95, Type: "Super Telephoto Zoom", ImageSide: "https://via.placeholder.com/400x183/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/87x87/4a90e2/ffffff", ImageIsometric: "", Length: 183.0, Diameter: 87.0, WeatherSealing: "sealed"},
	{ID: "34", Name: "Tamron 150-500mm f/5-6.7 Di III VC VXD", FocalLength: "150-500mm", MaxAperture: 5.0, MinAperture: 32.0, ImageStab: true, Weight: 1725, Price: 1399.95, Type: "Super Telephoto Zoom", ImageSide: "https://via.placeholder.com/400x283/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/93x93/4a90e2/ffffff", ImageIsometric: "", Length: 283.0, Diameter: 93.0, WeatherSealing: "sealed"},
	{ID: "35", Name: "Tamron 90mm f/2.8 Di III Macro VXD", FocalLength: "90mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 610, Price: 649.95, Type: "Macro Prime", ImageSide: "https://via.placeholder.com/300x125/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/73x73/4a90e2/ffffff", ImageIsometric: "", Length: 125.0, Diameter: 73.0, WeatherSealing: "sealed"},
	{ID: "36", Name: "Tamron 16-30mm f/2.8 Di III VXD G2", FocalLength: "16-30mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 620, Price: 999.95, Type: "Ultra Wide Angle Zoom", ImageSide: "https://via.placeholder.com/300x107/4a90e2/ffffff", ImageTop: "https://via.placeholder.com/82x82/4a90e2/ffffff", ImageIsometric: "", Length: 107.0, Diameter: 82.0, WeatherSealing: "sealed"},

	// Third-Party Z Mount Lenses - Sigma
	{ID: "37", Name: "Sigma 16mm f/1.4 DC DN Contemporary", FocalLength: "16mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 405, Price: 449.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x92/e8a87c/ffffff", ImageTop: "https://via.placeholder.com/73x73/e8a87c/ffffff", ImageIsometric: "", Length: 92.0, Diameter: 73.0, WeatherSealing: "dust"},
	{ID: "38", Name: "Sigma 30mm f/1.4 DC DN Contemporary", FocalLength: "30mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 265, Price: 339.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x74/e8a87c/ffffff", ImageTop: "https://via.placeholder.com/70x70/e8a87c/ffffff", ImageIsometric: "", Length: 74.0, Diameter: 70.0, WeatherSealing: "dust"},
	{ID: "39", Name: "Sigma 56mm f/1.4 DC DN Contemporary", FocalLength: "56mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 280, Price: 479.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x60/e8a87c/ffffff", ImageTop: "https://via.placeholder.com/67x67/e8a87c/ffffff", ImageIsometric: "", Length: 60.0, Diameter: 67.0, WeatherSealing: "dust"},

	// Third-Party Z Mount Lenses - Viltrox
	{ID: "40", Name: "Viltrox 16mm f/1.8 Z", FocalLength: "16mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 495, Price: 349.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x96/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/73x73/8e44ad/ffffff", ImageIsometric: "", Length: 96.0, Diameter: 73.0, WeatherSealing: "dust"},
	{ID: "41", Name: "Viltrox 24mm f/1.8 Z", FocalLength: "24mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 365, Price: 279.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x89/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/68x68/8e44ad/ffffff", ImageIsometric: "", Length: 89.0, Diameter: 68.0, WeatherSealing: "dust"},
	{ID: "42", Name: "Viltrox 35mm f/1.8 Z", FocalLength: "35mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 295, Price: 249.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x72/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/67x67/8e44ad/ffffff", ImageIsometric: "", Length: 72.0, Diameter: 67.0, WeatherSealing: "dust"},
	{ID: "43", Name: "Viltrox 50mm f/1.8 Z", FocalLength: "50mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 285, Price: 199.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x60/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/62x62/8e44ad/ffffff", ImageIsometric: "", Length: 60.0, Diameter: 62.0, WeatherSealing: "dust"},
	{ID: "44", Name: "Viltrox 85mm f/1.8 Z", FocalLength: "85mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 485, Price: 349.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x86/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/73x73/8e44ad/ffffff", ImageIsometric: "", Length: 86.0, Diameter: 73.0, WeatherSealing: "dust"},
	{ID: "45", Name: "Viltrox 135mm f/1.8 Z LAB", FocalLength: "135mm", MaxAperture: 1.8, MinAperture: 22.0, ImageStab: false, Weight: 935, Price: 849.95, Type: "Telephoto Prime", ImageSide: "https://via.placeholder.com/350x120/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/82x82/8e44ad/ffffff", ImageIsometric: "", Length: 120.0, Diameter: 82.0, WeatherSealing: "sealed"},
	{ID: "46", Name: "Viltrox 13mm f/1.4 Z", FocalLength: "13mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 480, Price: 349.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x91/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/72x72/8e44ad/ffffff", ImageIsometric: "", Length: 91.0, Diameter: 72.0, WeatherSealing: "dust"},
	{ID: "47", Name: "Viltrox 27mm f/1.2 Z Pro", FocalLength: "27mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 565, Price: 449.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x98/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/77x77/8e44ad/ffffff", ImageIsometric: "", Length: 98.0, Diameter: 77.0, WeatherSealing: "dust"},
	{ID: "48", Name: "Viltrox 56mm f/1.4 Z", FocalLength: "56mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 290, Price: 329.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x66/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/68x68/8e44ad/ffffff", ImageIsometric: "", Length: 66.0, Diameter: 68.0, WeatherSealing: "dust"},
	{ID: "49", Name: "Viltrox 75mm f/1.2 Z Pro", FocalLength: "75mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 670, Price: 599.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x95/8e44ad/ffffff", ImageTop: "https://via.placeholder.com/78x78/8e44ad/ffffff", ImageIsometric: "", Length: 95.0, Diameter: 78.0, WeatherSealing: "dust"},

	// Third-Party Z Mount Lenses - Sirui
	{ID: "50", Name: "Sirui 35mm f/1.4 Aurora", FocalLength: "35mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 530, Price: 599.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x87/f39c12/ffffff", ImageTop: "https://via.placeholder.com/72x72/f39c12/ffffff", ImageIsometric: "", Length: 87.0, Diameter: 72.0, WeatherSealing: "sealed"},
	{ID: "51", Name: "Sirui 85mm f/1.4 Aurora", FocalLength: "85mm", MaxAperture: 1.4, MinAperture: 16.0, ImageStab: false, Weight: 870, Price: 799.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x108/f39c12/ffffff", ImageTop: "https://via.placeholder.com/78x78/f39c12/ffffff", ImageIsometric: "", Length: 108.0, Diameter: 78.0, WeatherSealing: "sealed"},
	{ID: "52", Name: "Sirui 16mm f/1.2 Sniper", FocalLength: "16mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 620, Price: 649.95, Type: "Wide Prime", ImageSide: "https://via.placeholder.com/300x93/f39c12/ffffff", ImageTop: "https://via.placeholder.com/75x75/f39c12/ffffff", ImageIsometric: "", Length: 93.0, Diameter: 75.0, WeatherSealing: "dust"},
	{ID: "53", Name: "Sirui 56mm f/1.2 Sniper", FocalLength: "56mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 440, Price: 549.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x82/f39c12/ffffff", ImageTop: "https://via.placeholder.com/73x73/f39c12/ffffff", ImageIsometric: "", Length: 82.0, Diameter: 73.0, WeatherSealing: "dust"},
	{ID: "54", Name: "Sirui 75mm f/1.2 Sniper", FocalLength: "75mm", MaxAperture: 1.2, MinAperture: 16.0, ImageStab: false, Weight: 660, Price: 649.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x97/f39c12/ffffff", ImageTop: "https://via.placeholder.com/77x77/f39c12/ffffff", ImageIsometric: "", Length: 97.0, Diameter: 77.0, WeatherSealing: "dust"},

	// Third-Party Z Mount Lenses - Meike
	{ID: "55", Name: "Meike 35mm f/1.8 Pro AF", FocalLength: "35mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 325, Price: 279.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x78/27ae60/ffffff", ImageTop: "https://via.placeholder.com/70x70/27ae60/ffffff", ImageIsometric: "", Length: 78.0, Diameter: 70.0, WeatherSealing: "dust"},
	{ID: "56", Name: "Meike 50mm f/1.8 AF", FocalLength: "50mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 290, Price: 249.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x65/27ae60/ffffff", ImageTop: "https://via.placeholder.com/67x67/27ae60/ffffff", ImageIsometric: "", Length: 65.0, Diameter: 67.0, WeatherSealing: "dust"},
	{ID: "57", Name: "Meike 85mm f/1.8 AF STM", FocalLength: "85mm", MaxAperture: 1.8, MinAperture: 16.0, ImageStab: false, Weight: 425, Price: 299.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x84/27ae60/ffffff", ImageTop: "https://via.placeholder.com/72x72/27ae60/ffffff", ImageIsometric: "", Length: 84.0, Diameter: 72.0, WeatherSealing: "dust"},

	// Third-Party Z Mount Lenses - Laowa
	{ID: "58", Name: "Laowa 10mm f/2.8 Zero-D", FocalLength: "10mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 475, Price: 599.95, Type: "Ultra Wide Prime", ImageSide: "https://via.placeholder.com/300x75/34495e/ffffff", ImageTop: "https://via.placeholder.com/71x71/34495e/ffffff", ImageIsometric: "", Length: 75.0, Diameter: 71.0, WeatherSealing: "sealed"},
	{ID: "59", Name: "Laowa 12mm f/2.8 Zero-D", FocalLength: "12mm", MaxAperture: 2.8, MinAperture: 22.0, ImageStab: false, Weight: 545, Price: 549.95, Type: "Ultra Wide Prime", ImageSide: "https://via.placeholder.com/300x72/34495e/ffffff", ImageTop: "https://via.placeholder.com/70x70/34495e/ffffff", ImageIsometric: "", Length: 72.0, Diameter: 70.0, WeatherSealing: "sealed"},
	{ID: "60", Name: "Laowa 200mm f/2", FocalLength: "200mm", MaxAperture: 2.0, MinAperture: 22.0, ImageStab: false, Weight: 1850, Price: 3499.95, Type: "Telephoto Prime", ImageSide: "https://via.placeholder.com/350x211/34495e/ffffff", ImageTop: "https://via.placeholder.com/105x105/34495e/ffffff", ImageIsometric: "", Length: 211.0, Diameter: 105.0, WeatherSealing: "sealed"},

	// Third-Party Z Mount Lenses - TTArtisan
	{ID: "61", Name: "TTArtisan 32mm f/2.8 Z", FocalLength: "32mm", MaxAperture: 2.8, MinAperture: 16.0, ImageStab: false, Weight: 125, Price: 149.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x25/c0392b/ffffff", ImageTop: "https://via.placeholder.com/52x52/c0392b/ffffff", ImageIsometric: "", Length: 25.0, Diameter: 52.0, WeatherSealing: "dust"},
	{ID: "62", Name: "TTArtisan 75mm f/2 Z", FocalLength: "75mm", MaxAperture: 2.0, MinAperture: 16.0, ImageStab: false, Weight: 310, Price: 199.95, Type: "Prime", ImageSide: "https://via.placeholder.com/300x68/c0392b/ffffff", ImageTop: "https://via.placeholder.com/62x62/c0392b/ffffff", ImageIsometric: "", Length: 68.0, Diameter: 62.0, WeatherSealing: "dust"},
}

var cars = []Car{
	// Sedans
	{ID: "c1", Name: "Camry", Model: "Camry", Trim: "", Variant: "", Manufacturer: "Toyota", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "2.5L I4", Horsepower: 203, Torque: 184, Transmission: "Automatic", Drivetrain: "FWD", MPGCity: 28, MPGHighway: 39, MPGCombined: 32, FuelCapacity: 15.8, Seating: 5, CargoCapacity: 15.1, Length: 192.1, Width: 72.4, Height: 56.9, Wheelbase: 111.2, Weight: 3307, Price: 26520, ImageFront: "https://via.placeholder.com/400x200/e74c3c/ffffff", ImageSide: "https://via.placeholder.com/600x200/e74c3c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", TopSpeed: 120, ZeroToSixty: 7.8, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c2", Name: "Model 3", Model: "Model 3", Trim: "", Variant: "", Manufacturer: "Tesla", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Electric Motor", Horsepower: 283, Torque: 302, Transmission: "Single-Speed", Drivetrain: "RWD", MPGCity: 134, MPGHighway: 126, MPGCombined: 130, FuelCapacity: 0, BatteryCapacity: 75, Seating: 5, CargoCapacity: 15.0, Length: 184.8, Width: 72.8, Height: 56.8, Wheelbase: 113.2, Weight: 3627, Price: 38990, ImageFront: "https://via.placeholder.com/400x200/3498db/ffffff", ImageSide: "https://via.placeholder.com/600x200/3498db/ffffff", ImageInterior: "https://via.placeholder.com/400x250/3498db/ffffff", ImageExterior: "https://via.placeholder.com/400x250/3498db/ffffff", TopSpeed: 140, ZeroToSixty: 5.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c3", Name: "Accord", Model: "Accord", Trim: "", Variant: "", Manufacturer: "Honda", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "1.5L Turbo I4", Horsepower: 192, Torque: 192, Transmission: "CVT", Drivetrain: "FWD", MPGCity: 29, MPGHighway: 37, MPGCombined: 32, FuelCapacity: 14.8, Seating: 5, CargoCapacity: 16.7, Length: 195.7, Width: 73.3, Height: 57.1, Wheelbase: 111.4, Weight: 3153, Price: 27295, ImageFront: "https://via.placeholder.com/400x200/2ecc71/ffffff", ImageSide: "https://via.placeholder.com/600x200/2ecc71/ffffff", ImageInterior: "https://via.placeholder.com/400x250/2ecc71/ffffff", ImageExterior: "https://via.placeholder.com/400x250/2ecc71/ffffff", TopSpeed: 118, ZeroToSixty: 7.5, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c4", Name: "3 Series", Model: "3 Series", Trim: "", Variant: "", Manufacturer: "BMW", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "2.0L Turbo I4", Horsepower: 255, Torque: 295, Transmission: "Automatic", Drivetrain: "RWD", MPGCity: 26, MPGHighway: 36, MPGCombined: 30, FuelCapacity: 15.8, Seating: 5, CargoCapacity: 17.0, Length: 185.7, Width: 71.9, Height: 56.8, Wheelbase: 112.2, Weight: 3563, Price: 43245, ImageFront: "https://via.placeholder.com/400x200/9b59b6/ffffff", ImageSide: "https://via.placeholder.com/600x200/9b59b6/ffffff", ImageInterior: "https://via.placeholder.com/400x250/9b59b6/ffffff", ImageExterior: "https://via.placeholder.com/400x250/9b59b6/ffffff", TopSpeed: 130, ZeroToSixty: 5.8, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c5", Name: "C-Class", Model: "C-Class", Trim: "", Variant: "", Manufacturer: "Mercedes-Benz", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "2.0L Turbo I4", Horsepower: 255, Torque: 295, Transmission: "Automatic", Drivetrain: "RWD", MPGCity: 25, MPGHighway: 36, MPGCombined: 29, FuelCapacity: 15.9, Seating: 5, CargoCapacity: 12.6, Length: 187.0, Width: 72.2, Height: 56.7, Wheelbase: 111.5, Weight: 3642, Price: 44850, ImageFront: "https://via.placeholder.com/400x200/f39c12/ffffff", ImageSide: "https://via.placeholder.com/600x200/f39c12/ffffff", ImageInterior: "https://via.placeholder.com/400x250/f39c12/ffffff", ImageExterior: "https://via.placeholder.com/400x250/f39c12/ffffff", TopSpeed: 130, ZeroToSixty: 6.0, SafetyRating: "5 Stars", FuelType: "Gasoline"},

	// SUVs
	{ID: "c6", Name: "RAV4", Model: "RAV4", Trim: "", Variant: "", Manufacturer: "Toyota", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "2.5L I4 Hybrid", Horsepower: 219, Torque: 163, Transmission: "CVT", Drivetrain: "AWD", MPGCity: 41, MPGHighway: 38, MPGCombined: 40, FuelCapacity: 14.5, Seating: 5, CargoCapacity: 37.5, Length: 181.5, Width: 73.4, Height: 67.0, Wheelbase: 105.9, Weight: 3660, Price: 31075, ImageFront: "https://via.placeholder.com/400x220/e67e22/ffffff", ImageSide: "https://via.placeholder.com/600x220/e67e22/ffffff", ImageInterior: "https://via.placeholder.com/400x250/e67e22/ffffff", ImageExterior: "https://via.placeholder.com/400x250/e67e22/ffffff", TopSpeed: 115, ZeroToSixty: 7.8, SafetyRating: "5 Stars", FuelType: "Hybrid"},
	{ID: "c7", Name: "Model Y", Model: "Model Y", Trim: "", Variant: "", Manufacturer: "Tesla", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Electric Motor", Horsepower: 384, Torque: 376, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 127, MPGHighway: 117, MPGCombined: 122, FuelCapacity: 0, BatteryCapacity: 75, Seating: 5, CargoCapacity: 76.2, Length: 187.0, Width: 75.6, Height: 64.0, Wheelbase: 113.8, Weight: 4416, Price: 47740, ImageFront: "https://via.placeholder.com/400x220/2980b9/ffffff", ImageSide: "https://via.placeholder.com/600x220/2980b9/ffffff", ImageInterior: "https://via.placeholder.com/400x250/2980b9/ffffff", ImageExterior: "https://via.placeholder.com/400x250/2980b9/ffffff", TopSpeed: 135, ZeroToSixty: 5.0, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c8", Name: "CR-V", Model: "CR-V", Trim: "", Variant: "", Manufacturer: "Honda", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "1.5L Turbo I4", Horsepower: 190, Torque: 179, Transmission: "CVT", Drivetrain: "AWD", MPGCity: 28, MPGHighway: 34, MPGCombined: 30, FuelCapacity: 14.0, Seating: 5, CargoCapacity: 39.2, Length: 184.8, Width: 73.0, Height: 66.5, Wheelbase: 106.3, Weight: 3373, Price: 31350, ImageFront: "https://via.placeholder.com/400x220/27ae60/ffffff", ImageSide: "https://via.placeholder.com/600x220/27ae60/ffffff", ImageInterior: "https://via.placeholder.com/400x250/27ae60/ffffff", ImageExterior: "https://via.placeholder.com/400x250/27ae60/ffffff", TopSpeed: 115, ZeroToSixty: 7.5, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c9", Name: "X3", Model: "X3", Trim: "", Variant: "", Manufacturer: "BMW", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "2.0L Turbo I4", Horsepower: 248, Torque: 258, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 23, MPGHighway: 29, MPGCombined: 25, FuelCapacity: 17.2, Seating: 5, CargoCapacity: 28.7, Length: 185.9, Width: 74.4, Height: 66.0, Wheelbase: 112.8, Weight: 4145, Price: 46695, ImageFront: "https://via.placeholder.com/400x220/8e44ad/ffffff", ImageSide: "https://via.placeholder.com/600x220/8e44ad/ffffff", ImageInterior: "https://via.placeholder.com/400x250/8e44ad/ffffff", ImageExterior: "https://via.placeholder.com/400x250/8e44ad/ffffff", TopSpeed: 130, ZeroToSixty: 6.0, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c10", Name: "GLC-Class", Model: "GLC-Class", Trim: "", Variant: "", Manufacturer: "Mercedes-Benz", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "2.0L Turbo I4", Horsepower: 255, Torque: 273, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 22, MPGHighway: 29, MPGCombined: 25, FuelCapacity: 17.4, Seating: 5, CargoCapacity: 19.4, Length: 183.3, Width: 74.4, Height: 64.7, Wheelbase: 113.1, Weight: 4092, Price: 48100, ImageFront: "https://via.placeholder.com/400x220/e67e22/ffffff", ImageSide: "https://via.placeholder.com/600x220/e67e22/ffffff", ImageInterior: "https://via.placeholder.com/400x250/e67e22/ffffff", ImageExterior: "https://via.placeholder.com/400x250/e67e22/ffffff", TopSpeed: 130, ZeroToSixty: 6.2, SafetyRating: "5 Stars", FuelType: "Gasoline"},

	// Trucks
	{ID: "c11", Name: "F-150", Model: "F-150", Trim: "", Variant: "", Manufacturer: "Ford", Year: 2024, Type: "Truck", BodyStyle: "Pickup", Engine: "3.5L V6 EcoBoost", Horsepower: 400, Torque: 500, Transmission: "Automatic", Drivetrain: "4WD", MPGCity: 18, MPGHighway: 24, MPGCombined: 20, FuelCapacity: 36.0, Seating: 5, CargoCapacity: 77.4, Length: 231.9, Width: 79.9, Height: 77.2, Wheelbase: 145.4, Weight: 4750, Price: 36570, ImageFront: "https://via.placeholder.com/500x250/3498db/ffffff", ImageSide: "https://via.placeholder.com/700x250/3498db/ffffff", ImageInterior: "https://via.placeholder.com/400x250/3498db/ffffff", ImageExterior: "https://via.placeholder.com/400x250/3498db/ffffff", TopSpeed: 112, ZeroToSixty: 5.8, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c12", Name: "Silverado 1500", Model: "Silverado", Trim: "1500", Variant: "", Manufacturer: "Chevrolet", Year: 2024, Type: "Truck", BodyStyle: "Pickup", Engine: "5.3L V8", Horsepower: 355, Torque: 383, Transmission: "Automatic", Drivetrain: "4WD", MPGCity: 16, MPGHighway: 20, MPGCombined: 17, FuelCapacity: 24.0, Seating: 5, CargoCapacity: 71.7, Length: 231.7, Width: 81.2, Height: 75.5, Wheelbase: 147.4, Weight: 4985, Price: 37200, ImageFront: "https://via.placeholder.com/500x250/e74c3c/ffffff", ImageSide: "https://via.placeholder.com/700x250/e74c3c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", TopSpeed: 110, ZeroToSixty: 6.5, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c13", Name: "Ram 1500", Model: "Ram", Trim: "1500", Variant: "", Manufacturer: "Ram", Year: 2024, Type: "Truck", BodyStyle: "Pickup", Engine: "3.6L V6", Horsepower: 305, Torque: 269, Transmission: "Automatic", Drivetrain: "4WD", MPGCity: 20, MPGHighway: 25, MPGCombined: 22, FuelCapacity: 26.0, Seating: 5, CargoCapacity: 76.3, Length: 231.7, Width: 82.1, Height: 77.5, Wheelbase: 144.5, Weight: 4790, Price: 39105, ImageFront: "https://via.placeholder.com/500x250/16a085/ffffff", ImageSide: "https://via.placeholder.com/700x250/16a085/ffffff", ImageInterior: "https://via.placeholder.com/400x250/16a085/ffffff", ImageExterior: "https://via.placeholder.com/400x250/16a085/ffffff", TopSpeed: 110, ZeroToSixty: 7.2, SafetyRating: "5 Stars", FuelType: "Gasoline"},

	// Sports Cars
	{ID: "c14", Name: "Mustang", Model: "Mustang", Trim: "", Variant: "", Manufacturer: "Ford", Year: 2024, Type: "Coupe", BodyStyle: "Coupe", Engine: "5.0L V8", Horsepower: 450, Torque: 410, Transmission: "Manual", Drivetrain: "RWD", MPGCity: 16, MPGHighway: 25, MPGCombined: 19, FuelCapacity: 15.5, Seating: 4, CargoCapacity: 13.5, Length: 188.5, Width: 75.4, Height: 54.3, Wheelbase: 107.1, Weight: 3706, Price: 32105, ImageFront: "https://via.placeholder.com/400x180/3498db/ffffff", ImageSide: "https://via.placeholder.com/600x180/3498db/ffffff", ImageInterior: "https://via.placeholder.com/400x250/3498db/ffffff", ImageExterior: "https://via.placeholder.com/400x250/3498db/ffffff", TopSpeed: 155, ZeroToSixty: 4.3, SafetyRating: "5 Stars", FuelType: "Gasoline"},
	{ID: "c15", Name: "Challenger", Model: "Challenger", Trim: "", Variant: "", Manufacturer: "Dodge", Year: 2024, Type: "Coupe", BodyStyle: "Coupe", Engine: "5.7L V8", Horsepower: 375, Torque: 410, Transmission: "Automatic", Drivetrain: "RWD", MPGCity: 16, MPGHighway: 25, MPGCombined: 19, FuelCapacity: 18.5, Seating: 5, CargoCapacity: 16.2, Length: 197.9, Width: 75.7, Height: 57.5, Wheelbase: 116.2, Weight: 3950, Price: 34395, ImageFront: "https://via.placeholder.com/400x180/e74c3c/ffffff", ImageSide: "https://via.placeholder.com/600x180/e74c3c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/e74c3c/ffffff", TopSpeed: 155, ZeroToSixty: 5.0, SafetyRating: "4 Stars", FuelType: "Gasoline"},

	// Lucid Air Variants
	{ID: "c16", Name: "Air Pure", Model: "Air", Trim: "Pure", Variant: "", Manufacturer: "Lucid", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Single Motor", Horsepower: 430, Torque: 406, Transmission: "Single-Speed", Drivetrain: "RWD", MPGCity: 131, MPGHighway: 124, MPGCombined: 127, FuelCapacity: 0, BatteryCapacity: 88, Seating: 5, CargoCapacity: 28.1, Length: 195.9, Width: 76.2, Height: 55.5, Wheelbase: 116.5, Weight: 4654, Price: 77400, ImageFront: "https://via.placeholder.com/400x200/1abc9c/ffffff", ImageSide: "https://via.placeholder.com/600x200/1abc9c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", TopSpeed: 124, ZeroToSixty: 4.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c17", Name: "Air Touring", Model: "Air", Trim: "Touring", Variant: "", Manufacturer: "Lucid", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 620, Torque: 885, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 127, MPGHighway: 121, MPGCombined: 124, FuelCapacity: 0, BatteryCapacity: 92, Seating: 5, CargoCapacity: 28.1, Length: 195.9, Width: 76.2, Height: 55.5, Wheelbase: 116.5, Weight: 4839, Price: 95000, ImageFront: "https://via.placeholder.com/400x200/1abc9c/ffffff", ImageSide: "https://via.placeholder.com/600x200/1abc9c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", TopSpeed: 130, ZeroToSixty: 3.4, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c18", Name: "Air Grand Touring", Model: "Air", Trim: "Grand Touring", Variant: "", Manufacturer: "Lucid", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 819, Torque: 885, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 121, MPGHighway: 117, MPGCombined: 119, FuelCapacity: 0, BatteryCapacity: 112, Seating: 5, CargoCapacity: 28.1, Length: 195.9, Width: 76.2, Height: 55.5, Wheelbase: 116.5, Weight: 5162, Price: 125600, ImageFront: "https://via.placeholder.com/400x200/1abc9c/ffffff", ImageSide: "https://via.placeholder.com/600x200/1abc9c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", TopSpeed: 168, ZeroToSixty: 3.0, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c19", Name: "Air Sapphire", Model: "Air", Trim: "Sapphire", Variant: "", Manufacturer: "Lucid", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Tri Motor", Horsepower: 1234, Torque: 1430, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 105, MPGHighway: 102, MPGCombined: 103, FuelCapacity: 0, BatteryCapacity: 118, Seating: 5, CargoCapacity: 28.1, Length: 195.9, Width: 76.2, Height: 55.5, Wheelbase: 116.5, Weight: 5370, Price: 249000, ImageFront: "https://via.placeholder.com/400x200/1abc9c/ffffff", ImageSide: "https://via.placeholder.com/600x200/1abc9c/ffffff", ImageInterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", ImageExterior: "https://via.placeholder.com/400x250/1abc9c/ffffff", TopSpeed: 205, ZeroToSixty: 1.89, SafetyRating: "5 Stars", FuelType: "Electric"},

	// Porsche Taycan Variants
	{ID: "c20", Name: "Taycan", Model: "Taycan", Trim: "", Variant: "", Manufacturer: "Porsche", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Single Motor", Horsepower: 402, Torque: 254, Transmission: "Automatic", Drivetrain: "RWD", MPGCity: 112, MPGHighway: 96, MPGCombined: 104, FuelCapacity: 0, BatteryCapacity: 79.2, Seating: 4, CargoCapacity: 17.0, Length: 195.4, Width: 77.4, Height: 54.3, Wheelbase: 114.2, Weight: 4566, Price: 90500, ImageFront: "https://via.placeholder.com/400x200/c0392b/ffffff", ImageSide: "https://via.placeholder.com/600x200/c0392b/ffffff", ImageInterior: "https://via.placeholder.com/400x250/c0392b/ffffff", ImageExterior: "https://via.placeholder.com/400x250/c0392b/ffffff", TopSpeed: 143, ZeroToSixty: 5.1, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c21", Name: "Taycan 4S", Model: "Taycan", Trim: "4S", Variant: "", Manufacturer: "Porsche", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 522, Torque: 472, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 108, MPGHighway: 93, MPGCombined: 101, FuelCapacity: 0, BatteryCapacity: 93.4, Seating: 4, CargoCapacity: 17.0, Length: 195.4, Width: 77.4, Height: 54.3, Wheelbase: 114.2, Weight: 4832, Price: 108700, ImageFront: "https://via.placeholder.com/400x200/c0392b/ffffff", ImageSide: "https://via.placeholder.com/600x200/c0392b/ffffff", ImageInterior: "https://via.placeholder.com/400x250/c0392b/ffffff", ImageExterior: "https://via.placeholder.com/400x250/c0392b/ffffff", TopSpeed: 155, ZeroToSixty: 3.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c22", Name: "Taycan GTS", Model: "Taycan", Trim: "GTS", Variant: "", Manufacturer: "Porsche", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 590, Torque: 626, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 105, MPGHighway: 90, MPGCombined: 98, FuelCapacity: 0, BatteryCapacity: 93.4, Seating: 4, CargoCapacity: 17.0, Length: 195.4, Width: 77.4, Height: 54.3, Wheelbase: 114.2, Weight: 4884, Price: 132700, ImageFront: "https://via.placeholder.com/400x200/c0392b/ffffff", ImageSide: "https://via.placeholder.com/600x200/c0392b/ffffff", ImageInterior: "https://via.placeholder.com/400x250/c0392b/ffffff", ImageExterior: "https://via.placeholder.com/400x250/c0392b/ffffff", TopSpeed: 155, ZeroToSixty: 3.5, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c23", Name: "Taycan Turbo", Model: "Taycan", Trim: "Turbo", Variant: "", Manufacturer: "Porsche", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 670, Torque: 626, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 103, MPGHighway: 88, MPGCombined: 96, FuelCapacity: 0, BatteryCapacity: 93.4, Seating: 4, CargoCapacity: 17.0, Length: 195.4, Width: 77.4, Height: 54.3, Wheelbase: 114.2, Weight: 5071, Price: 167500, ImageFront: "https://via.placeholder.com/400x200/c0392b/ffffff", ImageSide: "https://via.placeholder.com/600x200/c0392b/ffffff", ImageInterior: "https://via.placeholder.com/400x250/c0392b/ffffff", ImageExterior: "https://via.placeholder.com/400x250/c0392b/ffffff", TopSpeed: 161, ZeroToSixty: 3.0, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c24", Name: "Taycan Turbo S", Model: "Taycan", Trim: "Turbo S", Variant: "", Manufacturer: "Porsche", Year: 2024, Type: "Sedan", BodyStyle: "Sedan", Engine: "Dual Motor", Horsepower: 750, Torque: 774, Transmission: "Automatic", Drivetrain: "AWD", MPGCity: 101, MPGHighway: 86, MPGCombined: 94, FuelCapacity: 0, BatteryCapacity: 93.4, Seating: 4, CargoCapacity: 17.0, Length: 195.4, Width: 77.4, Height: 54.3, Wheelbase: 114.2, Weight: 5125, Price: 187000, ImageFront: "https://via.placeholder.com/400x200/c0392b/ffffff", ImageSide: "https://via.placeholder.com/600x200/c0392b/ffffff", ImageInterior: "https://via.placeholder.com/400x250/c0392b/ffffff", ImageExterior: "https://via.placeholder.com/400x250/c0392b/ffffff", TopSpeed: 161, ZeroToSixty: 2.6, SafetyRating: "5 Stars", FuelType: "Electric"},

	// Ford Mustang Mach-E Variants
	{ID: "c25", Name: "Mach-E Select", Model: "Mustang Mach-E", Trim: "Select", Variant: "RWD", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Single Motor", Horsepower: 266, Torque: 317, Transmission: "Single-Speed", Drivetrain: "RWD", MPGCity: 115, MPGHighway: 100, MPGCombined: 108, FuelCapacity: 0, BatteryCapacity: 70, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 4398, Price: 42995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 111, ZeroToSixty: 6.1, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c26", Name: "Mach-E Select", Model: "Mustang Mach-E", Trim: "Select", Variant: "AWD", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Dual Motor", Horsepower: 346, Torque: 428, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 110, MPGHighway: 96, MPGCombined: 103, FuelCapacity: 0, BatteryCapacity: 70, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 4610, Price: 45995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 111, ZeroToSixty: 5.2, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c27", Name: "Mach-E Premium", Model: "Mustang Mach-E", Trim: "Premium", Variant: "RWD", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Single Motor", Horsepower: 290, Torque: 317, Transmission: "Single-Speed", Drivetrain: "RWD", MPGCity: 113, MPGHighway: 99, MPGCombined: 106, FuelCapacity: 0, BatteryCapacity: 91, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 4563, Price: 51995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 111, ZeroToSixty: 5.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c28", Name: "Mach-E Premium", Model: "Mustang Mach-E", Trim: "Premium", Variant: "AWD", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Dual Motor", Horsepower: 346, Torque: 428, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 108, MPGHighway: 94, MPGCombined: 101, FuelCapacity: 0, BatteryCapacity: 91, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 4775, Price: 54995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 111, ZeroToSixty: 4.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c29", Name: "Mach-E GT", Model: "Mustang Mach-E", Trim: "GT", Variant: "", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Dual Motor", Horsepower: 480, Torque: 600, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 103, MPGHighway: 90, MPGCombined: 97, FuelCapacity: 0, BatteryCapacity: 91, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 4920, Price: 59995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 124, ZeroToSixty: 3.8, SafetyRating: "5 Stars", FuelType: "Electric"},
	{ID: "c30", Name: "Mach-E Rally", Model: "Mustang Mach-E", Trim: "Rally", Variant: "", Manufacturer: "Ford", Year: 2024, Type: "SUV", BodyStyle: "SUV", Engine: "Dual Motor", Horsepower: 480, Torque: 700, Transmission: "Single-Speed", Drivetrain: "AWD", MPGCity: 100, MPGHighway: 88, MPGCombined: 94, FuelCapacity: 0, BatteryCapacity: 91, Seating: 5, CargoCapacity: 29.7, Length: 186.0, Width: 74.1, Height: 63.1, Wheelbase: 117.5, Weight: 5020, Price: 62995, ImageFront: "https://via.placeholder.com/400x220/0052D4/ffffff", ImageSide: "https://via.placeholder.com/600x220/0052D4/ffffff", ImageInterior: "https://via.placeholder.com/400x250/0052D4/ffffff", ImageExterior: "https://via.placeholder.com/400x250/0052D4/ffffff", TopSpeed: 124, ZeroToSixty: 3.5, SafetyRating: "5 Stars", FuelType: "Electric"},
}

var excavators = []Excavator{
	// Huina Excavators - 1:14 Scale
	{ID: "e1", Name: "Huina 1580", Model: "1580", Manufacturer: "Huina", Scale: "1:14", Year: 2024, Weight: 3.2, Length: 45.0, Width: 12.0, Height: 25.0, BatteryType: "LiPo", BatteryCapacity: 1500, OperatingTime: 30, ControlChannels: 8, DiggingDepth: 8.0, LiftingCapacity: 0.5, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt", Price: 89.99, ImageFront: "https://via.placeholder.com/400x200/ff6b6b/ffffff", ImageSide: "https://via.placeholder.com/600x200/ff6b6b/ffffff", ImageTop: "https://via.placeholder.com/400x250/ff6b6b/ffffff", ImageDetail: "https://via.placeholder.com/400x250/ff6b6b/ffffff"},
	{ID: "e2", Name: "Huina 1593", Model: "1593", Manufacturer: "Huina", Scale: "1:14", Year: 2024, Weight: 3.5, Length: 48.0, Width: 13.0, Height: 26.0, BatteryType: "LiPo", BatteryCapacity: 1800, OperatingTime: 35, ControlChannels: 8, DiggingDepth: 9.0, LiftingCapacity: 0.6, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement", Price: 129.99, ImageFront: "https://via.placeholder.com/400x200/ff6b6b/ffffff", ImageSide: "https://via.placeholder.com/600x200/ff6b6b/ffffff", ImageTop: "https://via.placeholder.com/400x250/ff6b6b/ffffff", ImageDetail: "https://via.placeholder.com/400x250/ff6b6b/ffffff"},
	{ID: "e3", Name: "Huina 1599", Model: "1599", Manufacturer: "Huina", Scale: "1:14", Year: 2024, Weight: 4.0, Length: 50.0, Width: 14.0, Height: 28.0, BatteryType: "LiPo", BatteryCapacity: 2000, OperatingTime: 40, ControlChannels: 10, DiggingDepth: 10.0, LiftingCapacity: 0.7, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement, Sound Effects", Price: 179.99, ImageFront: "https://via.placeholder.com/400x200/ff6b6b/ffffff", ImageSide: "https://via.placeholder.com/600x200/ff6b6b/ffffff", ImageTop: "https://via.placeholder.com/400x250/ff6b6b/ffffff", ImageDetail: "https://via.placeholder.com/400x250/ff6b6b/ffffff"},
	{ID: "e4", Name: "Huina 1603", Model: "1603", Manufacturer: "Huina", Scale: "1:14", Year: 2024, Weight: 4.5, Length: 52.0, Width: 15.0, Height: 30.0, BatteryType: "LiPo", BatteryCapacity: 2200, OperatingTime: 45, ControlChannels: 12, DiggingDepth: 11.0, LiftingCapacity: 0.8, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement, Sound Effects, LED Lights", Price: 249.99, ImageFront: "https://via.placeholder.com/400x200/ff6b6b/ffffff", ImageSide: "https://via.placeholder.com/600x200/ff6b6b/ffffff", ImageTop: "https://via.placeholder.com/400x250/ff6b6b/ffffff", ImageDetail: "https://via.placeholder.com/400x250/ff6b6b/ffffff"},
	
	// Kabolite Excavators - 1:14 Scale
	{ID: "e5", Name: "Kabolite KB-1580", Model: "KB-1580", Manufacturer: "Kabolite", Scale: "1:14", Year: 2024, Weight: 3.3, Length: 46.0, Width: 12.5, Height: 25.5, BatteryType: "LiPo", BatteryCapacity: 1600, OperatingTime: 32, ControlChannels: 8, DiggingDepth: 8.5, LiftingCapacity: 0.55, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt", Price: 99.99, ImageFront: "https://via.placeholder.com/400x200/4ecdc4/ffffff", ImageSide: "https://via.placeholder.com/600x200/4ecdc4/ffffff", ImageTop: "https://via.placeholder.com/400x250/4ecdc4/ffffff", ImageDetail: "https://via.placeholder.com/400x250/4ecdc4/ffffff"},
	{ID: "e6", Name: "Kabolite KB-1593", Model: "KB-1593", Manufacturer: "Kabolite", Scale: "1:14", Year: 2024, Weight: 3.8, Length: 49.0, Width: 13.5, Height: 27.0, BatteryType: "LiPo", BatteryCapacity: 1900, OperatingTime: 38, ControlChannels: 9, DiggingDepth: 9.5, LiftingCapacity: 0.65, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement", Price: 149.99, ImageFront: "https://via.placeholder.com/400x200/4ecdc4/ffffff", ImageSide: "https://via.placeholder.com/600x200/4ecdc4/ffffff", ImageTop: "https://via.placeholder.com/400x250/4ecdc4/ffffff", ImageDetail: "https://via.placeholder.com/400x250/4ecdc4/ffffff"},
	{ID: "e7", Name: "Kabolite KB-1600", Model: "KB-1600", Manufacturer: "Kabolite", Scale: "1:14", Year: 2024, Weight: 4.2, Length: 51.0, Width: 14.5, Height: 29.0, BatteryType: "LiPo", BatteryCapacity: 2100, OperatingTime: 42, ControlChannels: 10, DiggingDepth: 10.5, LiftingCapacity: 0.75, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement, Sound Effects", Price: 199.99, ImageFront: "https://via.placeholder.com/400x200/4ecdc4/ffffff", ImageSide: "https://via.placeholder.com/600x200/4ecdc4/ffffff", ImageTop: "https://via.placeholder.com/400x250/4ecdc4/ffffff", ImageDetail: "https://via.placeholder.com/400x250/4ecdc4/ffffff"},
	
	// Lesu Excavators - 1:14 Scale
	{ID: "e8", Name: "Lesu LS-1580", Model: "LS-1580", Manufacturer: "Lesu", Scale: "1:14", Year: 2024, Weight: 3.4, Length: 47.0, Width: 12.8, Height: 26.0, BatteryType: "LiPo", BatteryCapacity: 1700, OperatingTime: 33, ControlChannels: 8, DiggingDepth: 8.8, LiftingCapacity: 0.58, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt", Price: 109.99, ImageFront: "https://via.placeholder.com/400x200/95e1d3/ffffff", ImageSide: "https://via.placeholder.com/600x200/95e1d3/ffffff", ImageTop: "https://via.placeholder.com/400x250/95e1d3/ffffff", ImageDetail: "https://via.placeholder.com/400x250/95e1d3/ffffff"},
	{ID: "e9", Name: "Lesu LS-1593", Model: "LS-1593", Manufacturer: "Lesu", Scale: "1:14", Year: 2024, Weight: 3.9, Length: 50.0, Width: 14.0, Height: 28.0, BatteryType: "LiPo", BatteryCapacity: 1950, OperatingTime: 39, ControlChannels: 9, DiggingDepth: 9.8, LiftingCapacity: 0.68, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement", Price: 159.99, ImageFront: "https://via.placeholder.com/400x200/95e1d3/ffffff", ImageSide: "https://via.placeholder.com/600x200/95e1d3/ffffff", ImageTop: "https://via.placeholder.com/400x250/95e1d3/ffffff", ImageDetail: "https://via.placeholder.com/400x250/95e1d3/ffffff"},
	{ID: "e10", Name: "Lesu LS-1605", Model: "LS-1605", Manufacturer: "Lesu", Scale: "1:14", Year: 2024, Weight: 4.3, Length: 53.0, Width: 15.2, Height: 30.5, BatteryType: "LiPo", BatteryCapacity: 2300, OperatingTime: 46, ControlChannels: 12, DiggingDepth: 11.5, LiftingCapacity: 0.85, RotationAngle: 360, Material: "Metal", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt, Track Movement, Sound Effects, LED Lights", Price: 269.99, ImageFront: "https://via.placeholder.com/400x200/95e1d3/ffffff", ImageSide: "https://via.placeholder.com/600x200/95e1d3/ffffff", ImageTop: "https://via.placeholder.com/400x250/95e1d3/ffffff", ImageDetail: "https://via.placeholder.com/400x250/95e1d3/ffffff"},
	
	// Other Manufacturers - Various Scales
	{ID: "e11", Name: "Bruder CAT 420F", Model: "420F", Manufacturer: "Bruder", Scale: "1:16", Year: 2024, Weight: 2.8, Length: 42.0, Width: 11.0, Height: 23.0, BatteryType: "NiMH", BatteryCapacity: 1200, OperatingTime: 25, ControlChannels: 6, DiggingDepth: 7.0, LiftingCapacity: 0.4, RotationAngle: 360, Material: "Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation", Price: 69.99, ImageFront: "https://via.placeholder.com/400x200/f38181/ffffff", ImageSide: "https://via.placeholder.com/600x200/f38181/ffffff", ImageTop: "https://via.placeholder.com/400x250/f38181/ffffff", ImageDetail: "https://via.placeholder.com/400x250/f38181/ffffff"},
	{ID: "e12", Name: "Siku CAT 320", Model: "320", Manufacturer: "Siku", Scale: "1:20", Year: 2024, Weight: 1.5, Length: 28.0, Width: 8.0, Height: 15.0, BatteryType: "AA", BatteryCapacity: 0, OperatingTime: 0, ControlChannels: 0, DiggingDepth: 4.0, LiftingCapacity: 0.2, RotationAngle: 360, Material: "Plastic", RemoteControl: "None", Functions: "Manual Operation", Price: 29.99, ImageFront: "https://via.placeholder.com/400x200/a8e6cf/ffffff", ImageSide: "https://via.placeholder.com/600x200/a8e6cf/ffffff", ImageTop: "https://via.placeholder.com/400x250/a8e6cf/ffffff", ImageDetail: "https://via.placeholder.com/400x250/a8e6cf/ffffff"},
	{ID: "e13", Name: "Huina 1580 Pro", Model: "1580 Pro", Manufacturer: "Huina", Scale: "1:16", Year: 2024, Weight: 2.5, Length: 38.0, Width: 10.0, Height: 20.0, BatteryType: "LiPo", BatteryCapacity: 1200, OperatingTime: 28, ControlChannels: 7, DiggingDepth: 6.5, LiftingCapacity: 0.35, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt", Price: 79.99, ImageFront: "https://via.placeholder.com/400x200/ff6b6b/ffffff", ImageSide: "https://via.placeholder.com/600x200/ff6b6b/ffffff", ImageTop: "https://via.placeholder.com/400x250/ff6b6b/ffffff", ImageDetail: "https://via.placeholder.com/400x250/ff6b6b/ffffff"},
	{ID: "e14", Name: "Kabolite KB-1580 Mini", Model: "KB-1580 Mini", Manufacturer: "Kabolite", Scale: "1:20", Year: 2024, Weight: 1.8, Length: 32.0, Width: 9.0, Height: 18.0, BatteryType: "LiPo", BatteryCapacity: 800, OperatingTime: 20, ControlChannels: 6, DiggingDepth: 5.0, LiftingCapacity: 0.25, RotationAngle: 360, Material: "Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation", Price: 59.99, ImageFront: "https://via.placeholder.com/400x200/4ecdc4/ffffff", ImageSide: "https://via.placeholder.com/600x200/4ecdc4/ffffff", ImageTop: "https://via.placeholder.com/400x250/4ecdc4/ffffff", ImageDetail: "https://via.placeholder.com/400x250/4ecdc4/ffffff"},
	{ID: "e15", Name: "Lesu LS-1580 Compact", Model: "LS-1580 Compact", Manufacturer: "Lesu", Scale: "1:18", Year: 2024, Weight: 2.2, Length: 35.0, Width: 9.5, Height: 19.0, BatteryType: "LiPo", BatteryCapacity: 1000, OperatingTime: 24, ControlChannels: 7, DiggingDepth: 5.5, LiftingCapacity: 0.3, RotationAngle: 360, Material: "Metal/Plastic", RemoteControl: "2.4GHz", Functions: "Digging, Lifting, Rotation, Bucket Tilt", Price: 89.99, ImageFront: "https://via.placeholder.com/400x200/95e1d3/ffffff", ImageSide: "https://via.placeholder.com/600x200/95e1d3/ffffff", ImageTop: "https://via.placeholder.com/400x250/95e1d3/ffffff", ImageDetail: "https://via.placeholder.com/400x250/95e1d3/ffffff"},
}

func getLenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lenses)
}

func getLens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, lens := range lenses {
		if lens.ID == params["id"] {
			json.NewEncoder(w).Encode(lens)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Lens not found"})
}

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, car := range cars {
		if car.ID == params["id"] {
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Car not found"})
}

func getExcavators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(excavators)
}

func getExcavator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, excavator := range excavators {
		if excavator.ID == params["id"] {
			json.NewEncoder(w).Encode(excavator)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Excavator not found"})
}

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/lenses", getLenses).Methods("GET")
	r.HandleFunc("/api/lenses/{id}", getLens).Methods("GET")
	r.HandleFunc("/api/cars", getCars).Methods("GET")
	r.HandleFunc("/api/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/api/excavators", getExcavators).Methods("GET")
	r.HandleFunc("/api/excavators/{id}", getExcavator).Methods("GET")

	// CORS middleware
	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins == "" {
		corsOrigins = "http://localhost"
	}
	allowedOrigins := strings.Split(corsOrigins, ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
