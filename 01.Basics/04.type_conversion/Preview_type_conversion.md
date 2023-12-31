Numeric Conversions 
    The most common numeric conversions
            -> Atoi (string to int)
            -> Itoa (int to string).
    
    example:
            i, err := strconv.Atoi("-42")
            s := strconv.Itoa(-42)

ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:

        b, err := strconv.ParseBool("true")
        f, err := strconv.ParseFloat("3.1415", 64)
        i, err := strconv.ParseInt("-42", 10, 64)
        u, err := strconv.ParseUint("42", 10, 64)


FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:

        s := strconv.FormatBool(true)
        s := strconv.FormatFloat(3.1415, 'E', -1, 64)
        s := strconv.FormatInt(-42, 16)
        s := strconv.FormatUint(42, 16)

reference -> https://pkg.go.dev/strconv#Itoa

Type Conversion:
        Type conversion is used to explicitly change the type of a value.
        It is used when you want to convert a value from one type to another explicitly.

Type Assertion:
        Type assertion is used to extract the underlying concrete value of an interface.
        Type assertion returns two values: the concrete value and a boolean indicating whether the assertion succeeded.