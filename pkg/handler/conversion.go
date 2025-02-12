package handler

import (
	"fmt"
	"strconv"
)

func ConvertUnits(valueStr, fromUnit, toUnit string) (float64, error) {

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %v", err)
	}

	convertedValue, err := performConversion(value, fromUnit, toUnit)
	if err != nil {
		return 0, err
	}

	return convertedValue, nil
}

func performConversion(value float64, fromUnit, toUnit string) (float64, error) {
	switch fromUnit {
	// Length conversions
	case "meters":
		switch toUnit {
		case "kilometers":
			return value / 1000, nil
		case "miles":
			return value * 0.000621371, nil
		case "millimeter":
			return value * 1000, nil
		case "centimeter":
			return value * 100, nil
		case "inch":
			return value * 39.3701, nil
		case "foot":
			return value * 3.28084, nil
		case "yard":
			return value * 1.09361, nil
		// Add more length conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
	case "kilometers":
		switch toUnit {
		case "meters":
			return value * 1000, nil
		case "miles":
			return value * 0.621371, nil
		case "millimeters":
			return value * 1_000_000, nil
		case "centimeters":
			return value * 100_000, nil
		case "inches":
			return value * 39_370.1, nil
		case "feet":
			return value * 3_280.84, nil
		case "yards":
			return value * 1_093.61, nil
		// Add more length conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
	case "miles":
		switch toUnit {
		case "meters":
			return value / 0.000621371, nil
		case "kilometers":
			return value / 0.621371, nil
        case "millimeters":
            return value * 1_609_344, nil
        case "centimeters":
            return value * 160_934.4, nil
        case "inches":
            return value * 63_360, nil
        case "feet":
            return value * 5_280, nil
        case "yards":
            return value * 1_760, nil
		// Add more length conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}

    case "millimeters":
		switch toUnit {
		case "meters":
			return value / 1000, nil
		case "kilometers":
			return value / 1_000_000, nil
		case "miles":
			return value / 1_609_344, nil
		case "centimeters":
			return value / 10, nil
		case "inches":
			return value / 25.4, nil
		case "feet":
			return value / 304.8, nil
		case "yards":
			return value / 914.4, nil
        default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
    case "centimeters":
		switch toUnit {
		case "meters":
			return value / 100, nil
		case "kilometers":
			return value / 100_000, nil
		case "miles":
			return value / 160_934.4, nil
		case "millimeters":
			return value * 10, nil
		case "inches":
			return value / 2.54, nil
		case "feet":
			return value / 30.48, nil
		case "yards":
			return value / 91.44, nil
        default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
    case "inches":
		switch toUnit {
		case "meters":
			return value / 39.3701, nil
		case "kilometers":
			return value / 39_370.1, nil
		case "miles":
			return value / 63_360, nil
		case "millimeters":
			return value * 25.4, nil
		case "centimeters":
			return value * 2.54, nil
		case "feet":
			return value / 12, nil
		case "yards":
			return value / 36, nil
        default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}

	case "feet":
		switch toUnit {
		case "meters":
			return value / 3.28084, nil
		case "kilometers":
			return value / 3_280.84, nil
		case "miles":
			return value / 5_280, nil
		case "millimeters":
			return value * 304.8, nil
		case "centimeters":
			return value * 30.48, nil
		case "inches":
			return value * 12, nil
		case "yards":
			return value / 3, nil
        default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}

	case "yards":
		switch toUnit {
		case "meters":
			return value / 1.09361, nil
		case "kilometers":
			return value / 1_093.61, nil
		case "miles":
			return value / 1_760, nil
		case "millimeters":
			return value * 914.4, nil
		case "centimeters":
			return value * 91.44, nil
		case "inches":
			return value * 36, nil
		case "feet":
			return value * 3, nil
        default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}

	// Weight conversions
case "grams":
    switch toUnit {
    case "kilograms":
        return value / 1000, nil
    case "pounds":
        return value * 0.00220462, nil
    case "milligrams":
        return value * 1000, nil
    case "ounces":
        return value * 0.03527396, nil
    default:
        return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
    }

case "kilograms":
    switch toUnit {
    case "grams":
        return value * 1000, nil
    case "pounds":
        return value * 2.20462, nil
    case "milligrams":
        return value * 1_000_000, nil
    case "ounces":
        return value * 35.27396, nil
    default:
        return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
    }

case "pounds":
    switch toUnit {
    case "grams":
        return value / 0.00220462, nil
    case "kilograms":
        return value / 2.20462, nil
    case "milligrams":
        return value * 453_592.37, nil
    case "ounces":
        return value * 16, nil
    default:
        return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
    }

case "milligrams":
    switch toUnit {
    case "grams":
        return value / 1000, nil
    case "kilograms":
        return value / 1_000_000, nil
    case "pounds":
        return value / 453_592.37, nil
    case "ounces":
        return value / 28_349.52, nil
    default:
        return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
    }

case "ounces":
    switch toUnit {
    case "grams":
        return value * 28.34952, nil
    case "kilograms":
        return value / 35.27396, nil
    case "pounds":
        return value / 16, nil
    case "milligrams":
        return value * 28_349.52, nil
    default:
        return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
    }

	// Temperature conversions
	case "celsius":
		switch toUnit {
		case "fahrenheit":
			return (value * 9 / 5) + 32, nil
		case "kelvin":
			return value + 273.15, nil
		// Add more temperature conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
	case "fahrenheit":
		switch toUnit {
		case "celsius":
			return (value - 32) * 5 / 9, nil
		case "kelvin":
			return ((value - 32) * 5 / 9) + 273.15, nil
		// Add more temperature conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}
	case "kelvin":
		switch toUnit {
		case "celsius":
			return value - 273.15, nil
		case "fahrenheit":
			return ((value - 273.15) * 9 / 5) + 32, nil
		// Add more temperature conversions as needed
		default:
			return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
		}

	// Add more unit categories as needed

	default:
		return 0, fmt.Errorf("conversion from %s to %s not supported", fromUnit, toUnit)
	}
}
