// Listing 11.16  The setField helper function

func setField(name, value string, t reflect.Type, v reflect.Value) {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if name == fieldName(field) {
			var dest reflect.Value
			switch field.Type.Kind() {
			default:
				fmt.Printf("Kind %s not supported.\n",
					field.Type.Kind())
				continue
			case reflect.Int:
				ival, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to int: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(ival)
			case reflect.Float64:
				fval, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to float64: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(fval)
			case reflect.String:
				dest = reflect.ValueOf(value)
			case reflect.Bool:
				bval, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to bool: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(bval)
			}
			v.Field(i).Set(dest)
		}
	}
}
