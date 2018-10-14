// Package main provides eve online t2 commodity product calculator
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var Rcomponents = map[string]int{}
var Rcompos = map[string]int{}
var Rinter = map[string]int{}
var Rsors = map[string]int{}

type Result map[string]int

var Data = map[string]Item{
	"Particle Accelerator Unit": {
		"T2 component",
		1,
		[]Ingr{
			{"Hypersynaptic Fibers", 1},
			{"Crystalline Carbonide", 28},
			{"Fullerides", 10},
		},
	},
	"Photon Microprocessor": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 6},
			{"Nanotransistors", 2},
			{"Crystalline Carbonide", 16},
			{"Photonic Metamaterials", 2},
		},
	},
	"Laser Focusing Crystals": {
		"T2 component",
		1,
		[]Ingr{
			{"Tungsten Carbide", 28},
			{"Hypersynaptic Fibers", 1},
			{"Fullerides", 10},
		},
	},
	"Nanoelectrical Microprocessor": {
		"T2 component",
		1,
		[]Ingr{
			{"Tungsten Carbide", 16},
			{"Phenolic Composites", 6},
			{"Terahertz Metamaterials", 2},
			{"Nanotransistors", 2},
		},
	},
	"Superconductor Rails": {
		"T2 component",
		1,
		[]Ingr{
			{"Hypersynaptic Fibers", 1},
			{"Fullerides", 10},
			{"Titanium Carbide", 28},
		},
	},
	"Sustained Shield Emitter": {
		"T2 component",
		1,
		[]Ingr{
			{"Ferrogel ", 1},
			{"Sylramic Fibers", 9},
			{"Titanium Carbide", 20},
		},
	},
	"Nanomechanical Microprocessor": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 6},
			{"Fernite Carbide", 16},
			{"Plasmonic Metamaterials", 2},
			{"Nanotransistors", 2},
		},
	},
	"Pulse Shield Emitter": {
		"T2 component",
		1,
		[]Ingr{
			{"Ferrogel", 1},
			{"Sylramic Fibers", 9},
			{"Crystalline Carbonide", 20},
		},
	},
	"Linear Shield Emitter": {
		"T2 component",
		1,
		[]Ingr{
			{"Tungsten Carbide", 20},
			{"Ferrogel", 1},
			{"Sylramic Fibers", 9},
		},
	},
	"Deflection Shield Emitter": {
		"T2 component",
		1,
		[]Ingr{
			{"Fernite Carbide", 20},
			{"Ferrogel", 1},
			{"Sylramic Fibers", 9},
		},
	},
	"Tesseract Capacitor Unit": {
		"T2 component",
		1,
		[]Ingr{
			{"Tungsten Carbide", 25},
			{"Nanotransistors", 1},
			{"Terahertz Metamaterials", 2},
			{"Fullerides", 10},
		},
	},
	"Quantum Microprocessor": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 6},
			{"Nanotransistors", 2},
			{"Nonlinear Metamaterials", 2},
			{"Titanium Carbide", 16},
		},
	},
	"Plasma Thruster": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 3},
			{"Fernite Carbide", 12},
			{"Ferrogel", 1},
		},
	},
	"Ladar Sensor Cluster": {
		"T2 component",
		1,
		[]Ingr{
			{"Fernite Carbide", 20},
			{"Hypersynaptic Fibers", 2},
			{"Nanotransistors", 1},
		},
	},
	"Plasma Pulse Generator": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 7},
			{"Nanotransistors", 2},
			{"Crystalline Carbonide", 20},
		},
	},
	"Oscillator Capacitor Unit": {
		"T2 component",
		1,
		[]Ingr{
			{"Nanotransistors", 1},
			{"Photonic Metamaterials", 2},
			{"Crystalline Carbonide", 25},
			{"Fullerides", 10},
		},
	},
	"Fernite Carbide Composite Armor Plate": {
		"T2 component",
		1,
		[]Ingr{
			{"Fernite Carbide", 40},
			{"Sylramic Fibers", 10},
		},
	},
	"Magnetometric Sensor Cluster": {
		"T2 component",
		1,
		[]Ingr{
			{"Nanotransistors", 1},
			{"Hypersynaptic Fibers", 2},
			{"Crystalline Carbonide", 20},
		},
	},
	"Ion Thruster": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 3},
			{"Ferrogel", 1},
			{"Crystalline Carbonide", 12},
		},
	},
	"Crystalline Carbonide Armor Plate": {
		"T2 component",
		1,
		[]Ingr{
			{"Sylramic Fibers", 10},
			{"Crystalline Carbonide", 40},
		},
	},
	"Fusion Reactor Unit": {
		"T2 component",
		1,
		[]Ingr{
			{"Fermionic Condensates", 2},
			{"Crystalline Carbonide", 9},
		},
	},
	"Thermonuclear Trigger Unit": {
		"T2 component",
		1,
		[]Ingr{
			{"Fernite Carbide", 28},
			{"Hypersynaptic Fibers", 1},
			{"Fullerides", 10},
		},
	},
	"Graviton Pulse Generator": {
		"T2 component",
		1,
		[]Ingr{
			{"Phenolic Composites", 7},
			{"Nanotransistors", 2},
			{"Titanium Carbide", 20},
		},
	},
	"Gravimetric Sensor Cluster": {
		"T2 component",
		1,
		[]Ingr{
			{"Nanotransistors", 1},
			{"Hypersynaptic Fibers", 2},
			{"Titanium Carbide", 20},
		},
	},
	"Hypersynaptic Fibers": {
		"Composite",
		750,
		[]Ingr{
			{"Solerium", 98},
			{"Dysporite", 98},
			{"Vanadium Hafnite", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Fullerides": {
		"Composite",
		3000,
		[]Ingr{
			{"Carbon Polymers", 98},
			{"Platinum Technite", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Crystalline Carbonide": {
		"Composite",
		10000,
		[]Ingr{
			{"Crystallite Alloy", 98},
			{"Carbon Polymers", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Tungsten Carbide": {
		"Composite",
		10000,
		[]Ingr{
			{"Rolled Tungsten Alloy", 98},
			{"Sulfuric Acid", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Titanium Carbide": {
		"Composite",
		10000,
		[]Ingr{
			{"Titanium Chromide", 98},
			{"Silicon Diborite", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Nanotransistors": {
		"Composite",
		1500,
		[]Ingr{
			{"Sulfuric Acid", 98},
			{"Platinum Technite", 98},
			{"Neo Mercurite", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Ferrogel": {
		"Composite",
		400,
		[]Ingr{
			{"Hexite", 98},
			{"Hyperflurite", 98},
			{"Ferrofluid", 98},
			{"Prometium", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Terahertz Metamaterials": {
		"Composite",
		300,
		[]Ingr{
			{"Rolled Tungsten Alloy", 98},
			{"Promethium Mercurite", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Fernite Carbide": {
		"Composite",
		10000,
		[]Ingr{
			{"Fernite Alloy", 98},
			{"Ceramic Powder", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Phenolic Composites": {
		"Composite",
		2200,
		[]Ingr{
			{"Silicon Diborite", 98},
			{"Caesarium Cadmide", 98},
			{"Vanadium Hafnite", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Nonlinear Metamaterials": {
		"Composite",
		300,
		[]Ingr{
			{"Titanium Chromide", 98},
			{"Ferrofluid", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Sylramic Fibers": {
		"Composite",
		600,
		[]Ingr{
			{"Ceramic Powder", 98},
			{"Hexite", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Photonic Metamaterials": {
		"Composite",
		300,
		[]Ingr{
			{"Crystallite Alloy", 98},
			{"Thulium Hafnite", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Fermionic Condensates": {
		"Composite",
		200,
		[]Ingr{
			{"Caesarium Cadmide", 98},
			{"Dysporite", 98},
			{"Fluxed Condensates", 98},
			{"Prometium", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Plasmonic Metamaterials": {
		"Composite",
		300,
		[]Ingr{
			{"Fernite Alloy", 98},
			{"Neo Mercurite", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Crystallite Alloy": {
		"Intermediate",
		200,
		[]Ingr{
			{"Cobalt", 98},
			{"Cadmium", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Carbon Polymers": {
		"Intermediate",
		200,
		[]Ingr{
			{"Hydrocarbons", 98},
			{"Silicates", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Platinum Technite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Platinum", 98},
			{"Technetium", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Solerium": {
		"Intermediate",
		200,
		[]Ingr{
			{"Chromium", 98},
			{"Caesium", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Dysporite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Mercury", 98},
			{"Dysprosium", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Vanadium Hafnite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Vanadium", 98},
			{"Hafnium", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Sulfuric Acid": {
		"Intermediate",
		200,
		[]Ingr{
			{"Atmospheric Gases", 98},
			{"Evaporite Deposits", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Rolled Tungsten Alloy": {
		"Intermediate",
		200,
		[]Ingr{
			{"Tungsten", 98},
			{"Platinum", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Titanium Chromide": {
		"Intermediate",
		200,
		[]Ingr{
			{"Titanium", 98},
			{"Chromium", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Silicon Diborite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Evaporite Deposits", 98},
			{"Silicates", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Neo Mercurite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Mercury", 98},
			{"Neodymium", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Hexite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Chromium", 98},
			{"Platinum", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Hyperflurite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Vanadium", 98},
			{"Promethium", 98},
			{"Nitrogen Fuel Block", 5},
		},
	},
	"Ferrofluid": {
		"Intermediate",
		200,
		[]Ingr{
			{"Hafnium", 98},
			{"Dysprosium", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Prometium": {
		"Intermediate",
		200,
		[]Ingr{
			{"Cadmium", 98},
			{"Promethium", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Promethium Mercurite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Mercury", 98},
			{"Promethium", 98},
			{"Helium Fuel Block", 5},
		},
	},
	"Caesarium Cadmide": {
		"Intermediate",
		200,
		[]Ingr{
			{"Cadmium", 98},
			{"Caesium", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
	"Ceramic Powder": {
		"Intermediate",
		200,
		[]Ingr{
			{"Evaporite Deposits", 98},
			{"Silicates", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Fernite Alloy": {
		"Intermediate",
		200,
		[]Ingr{
			{"Scandium", 98},
			{"Vanadium", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Thulium Hafnite": {
		"Intermediate",
		200,
		[]Ingr{
			{"Hafnium", 98},
			{"Thulium", 98},
			{"Hydrogen Fuel Block", 5},
		},
	},
	"Fluxed Condensates": {
		"Intermediate",
		200,
		[]Ingr{
			{"Neodymium", 98},
			{"Thulium", 98},
			{"Oxygen Fuel Block", 5},
		},
	},
}

type Item struct {
	Type string
	Res  int
	Ing  []Ingr
}

type Ingr struct {
	Name string
	Need int
}

func processOne(e Item, bc int) {
	if len(e.Ing) > 0 {
		runs := math.Ceil(float64(bc) / float64(e.Res))
		for i := 0; i < len(e.Ing); i++ {
			if e.Type == "T2 component" {
				if val, ok := Rcomponents[e.Ing[i].Name]; ok {
					Rcomponents[e.Ing[i].Name] = val + e.Ing[i].Need*int(runs)
				} else {
					Rcomponents[e.Ing[i].Name] = e.Ing[i].Need * int(runs)
				}
			} else if e.Type == "Composite" {
				if val, ok := Rcompos[e.Ing[i].Name]; ok {
					Rcompos[e.Ing[i].Name] = val + e.Ing[i].Need*int(runs)
				} else {
					Rcompos[e.Ing[i].Name] = e.Ing[i].Need * int(runs)
				}
			} else if e.Type == "Intermediate" {
				if val, ok := Rinter[e.Ing[i].Name]; ok {
					Rinter[e.Ing[i].Name] = val + e.Ing[i].Need*int(runs)
				} else {
					Rinter[e.Ing[i].Name] = e.Ing[i].Need * int(runs)
				}
			} else {
				if val, ok := Rsors[e.Ing[i].Name]; ok {
					Rsors[e.Ing[i].Name] = val + e.Ing[i].Need*int(runs)
				} else {
					Rsors[e.Ing[i].Name] = e.Ing[i].Need * int(runs)
				}
			}

			defer processOne(Data[e.Ing[i].Name], e.Ing[i].Need*int(runs))
		}
	}
}

func main() {
	var element string
	var buildcount int

	if len(os.Args) >= 2 {
		element = os.Args[1]
		c, _ := strconv.Atoi(os.Args[2])
		buildcount = c
	} else {
		fmt.Println("use: eve_t2_com_calc elementName elementNeedToProduct")
	}

	processOne(Data[element], buildcount)

	fmt.Println("Results:")
	fmt.Println(" ")
	fmt.Println("----------------  Composite components   ----------------")
	fmt.Println(" ")
	for k, v := range Rcomponents {
		fmt.Printf("%s:%d\n", k, v)
	}
	fmt.Println(" ")
	fmt.Println("---------------- Intermediate components ----------------")
	fmt.Println(" ")
	for k, v := range Rcompos {
		fmt.Printf("%s:%d\n", k, v)
	}
	fmt.Println(" ")
	fmt.Println("----------------        Resource         ----------------")
	fmt.Println(" ")
	for k, v := range Rinter {
		fmt.Printf("%s:%d\n", k, v)
	}
}
