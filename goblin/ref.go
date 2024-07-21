package goblin

// // Helper function to discover possible concrete types
// // func discoverConcreteTypes(interfaceType reflect.Type) {
// // 	// This is a simplification. In a real-world scenario, you'd need to
// // 	// implement logic to discover all types in your program that implement
// // 	// the interface. This could involve parsing source code or maintaining
// // 	// a registry of types.

// // 	// For demonstration, we'll check a few predefined types
// // 	possibleTypes := []interface{}{
// // 		struct{ Field string }{},
// // 		struct{ Value int }{},
// // 		struct{ Data []byte }{},
// // 	}

// // 	for _, t := range possibleTypes {
// // 		concreteType := reflect.TypeOf(t)
// // 		if concreteType.Implements(interfaceType) {
// // 			fmt.Printf("    - %v\n", concreteType)
// // 		}
// // 	}
// // }

// // Example interface and structs
// type ExampleInterface interface {
// 	Method()
// }

// type ConcreteTypeA struct {
// 	Field string
// }

// func (c ConcreteTypeA) Method() {}

// type ConcreteTypeB struct {
// 	Value int
// }

// func (c ConcreteTypeB) Method() {}

// // Example function that returns an interface
// func returnInterface() ExampleInterface {
// 	return ConcreteTypeB{Value: 1}
// }

// func analyzeFunction(fn interface{}) {
// 	// Get the reflect.Value of the function
// 	funcValue := reflect.ValueOf(fn)

// 	// Get the reflect.Type of the function
// 	funcType := funcValue.Type()

// 	// Print function name
// 	fmt.Printf("Function Name: %s\n", funcType.Name())

// 	// Print number of input parameters
// 	fmt.Printf("Number of input parameters: %d\n", funcType.NumIn())

// 	// Print details of each input parameter
// 	for i := 0; i < funcType.NumIn(); i++ {
// 		paramType := funcType.In(i)
// 		fmt.Printf("Input parameter %d: Type = %v\n", i, paramType)
// 	}

// 	// Print number of return values
// 	fmt.Printf("Number of return values: %d\n", funcType.NumOut())

// 	// Print details of each return value
// 	// for i := 0; i < funcType.NumOut(); i++ {
// 	// 	returnType := funcType.Out(i)
// 	// 	fmt.Printf("Return value %d: Type = %v\n", i, returnType)
// 	// }

// 	for i := 0; i < funcType.NumOut(); i++ {
// 		returnType := funcType.Out(i)
// 		fmt.Printf("Return value %d:\n", i)
// 		fmt.Printf("  Declared type: %v\n", returnType)

// 		if returnType.Kind() == reflect.Interface {
// 			fmt.Println("  This is an interface. Possible concrete types:")
// 			discoverConcreteTypes(returnType)
// 		}
// 	}
// }

// // // Analyzer function that attempts to discover concrete types
// // func analyzeFunction(fn interface{}) {
// // 	funcValue := reflect.ValueOf(fn)
// // 	funcType := funcValue.Type()

// // 	fmt.Printf("Function: %s\n", runtime.FuncForPC(funcValue.Pointer()).Name())
// // 	fmt.Printf("Number of return values: %d\n", funcType.NumOut())

// // 	for i := 0; i < funcType.NumOut(); i++ {
// // 		returnType := funcType.Out(i)
// // 		fmt.Printf("Return value %d:\n", i)
// // 		fmt.Printf("  Declared type: %v\n", returnType)

// //			if returnType.Kind() == reflect.Interface {
// //				fmt.Println("  This is an interface. Possible concrete types:")
// //				discoverConcreteTypes(returnType)
// //			}
// //		}
// //	}
// func getResponses(fn interface{}) {
// 	funcValue := reflect.ValueOf(fn)
// 	funcType := funcValue.Type()

// 	fmt.Printf("Function: %s\n", runtime.FuncForPC(funcValue.Pointer()).Name())
// 	fmt.Printf("Number of return values: %d\n", funcType.NumOut())

// 	// Analyze and create mock input parameters
// 	var inArgs []reflect.Value
// 	for i := 0; i < funcType.NumIn(); i++ {
// 		argType := funcType.In(i)
// 		fmt.Printf("Input parameter %d:\n", i)
// 		fmt.Printf("  Type: %v\n", argType)
// 		inArgs = append(inArgs, reflect.Zero(argType))
// 	}

// 	// Call the function with mock arguments
// 	returnValues := funcValue.Call(inArgs)

// 	for i, returnValue := range returnValues {
// 		fmt.Printf("Return value %d:\n", i)
// 		fmt.Printf("  Declared type: %v\n", funcType.Out(i))

// 		if returnValue.Kind() == reflect.Interface && !returnValue.IsNil() {
// 			concreteValue := reflect.ValueOf(returnValue.Interface())
// 			fmt.Printf("  Actual concrete type: %v\n", concreteValue.Type())
// 			fmt.Printf("  Concrete value: %+v\n", concreteValue.Interface())

// 		} else {
// 			fmt.Printf("  Type: %v\n", returnValue.Type())
// 			fmt.Printf("  Value: %v\n", returnValue.Interface())
// 		}
// 	}
// }

// // Analyzer function that determines the actual concrete type
// // func analyzeFunction(fn interface{}) {
// // 	funcValue := reflect.ValueOf(fn)
// // 	funcType := funcValue.Type()

// // 	fmt.Printf("Function: %s\n", runtime.FuncForPC(funcValue.Pointer()).Name())
// // 	fmt.Printf("Number of return values: %d\n", funcType.NumOut())

// // 	// Analyze and create mock input parameters
// // 	var inArgs []reflect.Value
// // 	for i := 0; i < funcType.NumIn(); i++ {
// // 		argType := funcType.In(i)
// // 		fmt.Printf("Input parameter %d:\n", i)
// // 		fmt.Printf("  Type: %v\n", argType)
// // 		inArgs = append(inArgs, reflect.Zero(argType))
// // 	}

// // 	// Call the function with mock arguments
// // 	returnValues := funcValue.Call(inArgs)

// // 	for i, returnValue := range returnValues {
// // 		fmt.Printf("Return value %d:\n", i)
// // 		fmt.Printf("  Declared type: %v\n", funcType.Out(i))

// // 		if returnValue.Kind() == reflect.Interface && !returnValue.IsNil() {
// // 			concreteValue := reflect.ValueOf(returnValue.Interface())
// // 			fmt.Printf("  Actual concrete type: %v\n", concreteValue.Type())
// // 			fmt.Printf("  Concrete value: %+v\n", concreteValue.Interface())
// // 		} else {
// // 			fmt.Printf("  Type: %v\n", returnValue.Type())
// // 			fmt.Printf("  Value: %v\n", returnValue.Interface())
// // 		}
// // 	}
// // }

// // Helper function to discover possible concrete types
// func discoverConcreteTypes(interfaceType reflect.Type) {
// 	// Define some example types that might implement the interface
// 	possibleTypes := []interface{}{
// 		ConcreteTypeA{},
// 		ConcreteTypeB{},
// 		struct{ Data []byte }{},
// 	}

// 	for _, t := range possibleTypes {
// 		concreteType := reflect.TypeOf(t)
// 		if concreteType.Implements(interfaceType) {
// 			fmt.Printf("    - %v\n", concreteType)
// 		}
// 	}
// }

// // type response struct {
// // 	Description *string             `json:"description,omitempty" yaml:"description,omitempty"`
// // 	Headers     map[string]             `json:"headers,omitempty" yaml:"headers,omitempty"`
// // 	Model       map[int]interface{} `json:"model,omitempty" yaml:"model,omitempty"`
// // 	Links       Links               `json:"links,omitempty" yaml:"links,omitempty"`
// // }

// // type Config struct {
// // 	path        string
// // 	method      Method
// // 	summary     string
// // 	description string
// // 	responses   map[int]interface{}
// // }

// // func idea(config Config, handler Handler) {}

// // Define your routes here
// // var routes = []Route{
// // 	{
// // 		Method: http.MethodGet,
// // 		Path:   "/hello",
// // 		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
// // 			w.Write([]byte("Hello, World!"))
// // 		},
// // 		Operation: &openapi3.Operation{
// // 			Summary:     "Say Hello",
// // 			Description: "Returns a hello message",
// // 			Responses: openapi3.NewResponses(
// // 				openapi3.WithStatus(200, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("OK").
// // 						WithContent(
// // 							openapi3.NewContentWithSchema(
// // 								openapi3.NewStringSchema(),
// // 								[]string{"text/plain"},
// // 							),
// // 						),
// // 				}),
// // 			),
// // 		},
// // 	},
// // 	{
// // 		Method: http.MethodGet,
// // 		Path:   "/users/{id}",
// // 		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
// // 			parts := strings.Split(r.URL.Path, "/")
// // 			if len(parts) != 3 {
// // 				http.Error(w, "Invalid URL", http.StatusBadRequest)
// // 				return
// // 			}
// // 			userID := parts[2]
// // 			if userID == "1" {
// // 				user := User{ID: 1, Name: "John Doe", Age: 30}
// // 				w.Header().Set("Content-Type", "application/json")
// // 				json.NewEncoder(w).Encode(user)
// // 			} else {
// // 				http.Error(w, "User not found", http.StatusNotFound)
// // 			}
// // 		},
// // 		Operation: &openapi3.Operation{
// // 			Summary:     "Get User",
// // 			Description: "Returns a user by ID",
// // 			Parameters: openapi3.Parameters{
// // 				&openapi3.ParameterRef{
// // 					Value: openapi3.NewPathParameter("id").
// // 						WithSchema(openapi3.NewIntegerSchema()),
// // 				},
// // 			},
// // 			Responses: openapi3.NewResponses(
// // 				openapi3.WithStatus(200, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("OK").
// // 						WithContent(
// // 							openapi3.NewContentWithSchemaRef(
// // 								openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
// // 								[]string{"application/json"},
// // 							),
// // 						),
// // 				}),
// // 				openapi3.WithStatus(404, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("Not Found"),
// // 				}),
// // 			),
// // 		},
// // 	},
// // 	{
// // 		Method: http.MethodPost,
// // 		Path:   "/users",
// // 		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
// // 			var user User
// // 			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// // 				http.Error(w, err.Error(), http.StatusBadRequest)
// // 				return
// // 			}
// // 			user.ID = 1 // Example assignment of user ID
// // 			w.Header().Set("Content-Type", "application/json")
// // 			w.WriteHeader(http.StatusCreated)
// // 			json.NewEncoder(w).Encode(user)
// // 		},
// // 		Operation: &openapi3.Operation{
// // 			Summary:     "Create User",
// // 			Description: "Creates a new user",
// // 			RequestBody: &openapi3.RequestBodyRef{
// // 				Value: &openapi3.RequestBody{
// // 					Required: true,
// // 					Content: openapi3.NewContentWithSchemaRef(
// // 						openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
// // 						[]string{"application/json"},
// // 					),
// // 				},
// // 			},
// // 			Responses: openapi3.NewResponses(
// // 				openapi3.WithStatus(201, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("User created").
// // 						WithContent(
// // 							openapi3.NewContentWithSchemaRef(
// // 								openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
// // 								[]string{"application/json"},
// // 							),
// // 						),
// // 				}),
// // 				openapi3.WithStatus(400, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("Invalid input"),
// // 				}),
// // 			),
// // 		},
// // 	},
// // }

// // func NewHandler(w http.ResponseWriter, r *http.Request) Handler {
// // 	return func(w http.ResponseWriter, r *http.Request)  {

// // 	}
// // }

// // func analyzeFunction(fn interface{}) {
// // 	// Get the function name
// // 	funcValue := reflect.ValueOf(fn)
// // 	funcName := runtime.FuncForPC(funcValue.Pointer()).Name()

// // 	// Parse the file containing the function
// // 	fset := token.NewFileSet()
// // 	file, err := parser.ParseFile(fset, "main.go", nil, 0)
// // 	if err != nil {
// // 		fmt.Printf("Error parsing file: %v\n", err)
// // 		return
// // 	}

// // 	// Find the function in the AST
// // 	var funcDecl *ast.FuncDecl
// // 	ast.Inspect(file, func(n ast.Node) bool {
// // 		if fd, ok := n.(*ast.FuncDecl); ok && fd.Name.Name == funcName {
// // 			funcDecl = fd
// // 			return false
// // 		}
// // 		return true
// // 	})

// // 	if funcDecl == nil {
// // 		fmt.Printf("Function %s not found in AST\n", funcName)
// // 		return
// // 	}

// // 	fmt.Printf("Analyzing function: %s\n", funcName)
// // 	fmt.Printf("Number of parameters: %d\n", len(funcDecl.Type.Params.List))
// // 	fmt.Printf("Number of return values: %d\n", len(funcDecl.Type.Results.List))

// // 	// Analyze return statements
// // 	returnTypes := make(map[string]bool)
// // 	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
// // 		if ret, ok := n.(*ast.ReturnStmt); ok {
// // 			for _, expr := range ret.Results {
// // 				if ident, ok := expr.(*ast.Ident); ok {
// // 					returnTypes[ident.Name] = true
// // 				} else if callExpr, ok := expr.(*ast.CallExpr); ok {
// // 					if ident, ok := callExpr.Fun.(*ast.Ident); ok {
// // 						returnTypes[ident.Name] = true
// // 					}
// // 				}
// // 			}
// // 		}
// // 		return true
// // 	})

// // 	fmt.Println("Possible return types:")
// // 	for returnType := range returnTypes {
// // 		fmt.Printf("  - %s\n", returnType)
// // 	}
// // }

// func identifyPossibleTypes(interfaceType reflect.Type) {
// 	// This is a simplified approach. In a real-world scenario, you'd need a more
// 	// comprehensive way to determine possible implementations.
// 	possibleTypes := []reflect.Type{
// 		reflect.TypeOf(""),
// 		reflect.TypeOf(0),
// 		reflect.TypeOf(struct{}{}),
// 		reflect.TypeOf((*error)(nil)).Elem(),
// 	}

// 	for _, t := range possibleTypes {
// 		if t.Implements(interfaceType) {
// 			fmt.Printf("  - %v\n", t)
// 		}
// 	}
// }

// func tryIt() {
// 	// Sample code to analyze
// 	functionBody := `
// 	func (NICE *http.Request) (responses.IResponse, error) {
// 		if x == "bad" {
// 			return responses.BadRequest("lol really?"), nil
// 		} else if x == "err" {
// 			return responses.InternalServerError("bang"), nil
// 		} else {
// 			pokemon, err := fetchPokemonData("pikachu")
// 			if err != nil {
// 				return responses.InternalServerError("bang"), nil
// 			}
// 			return responses.OK(pokemon).JSON(), nil
// 		}
// 	}
// 	`

// 	// Parse the code
// 	// fset := token.NewFileSet()
// 	node, err := parser.ParseExpr(fmt.Sprintf("func() { %s }", functionBody))
// 	if err != nil {
// 		log.Fatalf("Failed to parse function body: %v", err)
// 	}

// 	// Create a visitor to walk the AST
// 	v := &visitor{
// 		returnTypes: make(map[string]bool),
// 		dataTypes:   make(map[string]bool),
// 	}

// 	// Walk the AST
// 	ast.Walk(v, node)

// 	// Print the results
// 	fmt.Println("Return types:")
// 	for t := range v.returnTypes {
// 		fmt.Println("-", t)
// 	}

// 	fmt.Println("\nData types:")
// 	for t := range v.dataTypes {
// 		fmt.Println("-", t)
// 	}
// }

// type visitor struct {
// 	returnTypes map[string]bool
// 	dataTypes   map[string]bool
// }

// func (v *visitor) Visit(node ast.Node) ast.Visitor {
// 	if node == nil {
// 		return nil
// 	}

// 	switch n := node.(type) {
// 	case *ast.ReturnStmt:
// 		for _, result := range n.Results {
// 			if call, ok := result.(*ast.CallExpr); ok {
// 				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
// 					v.returnTypes[sel.Sel.Name] = true
// 				}
// 			}
// 		}
// 	case *ast.AssignStmt:
// 		for i, rhs := range n.Rhs {
// 			if call, ok := rhs.(*ast.CallExpr); ok {
// 				if ident, ok := call.Fun.(*ast.Ident); ok {
// 					v.dataTypes[ident.Name] = true
// 				}
// 			}
// 			// Check if the left-hand side is an identifier
// 			if i < len(n.Lhs) {
// 				if ident, ok := n.Lhs[i].(*ast.Ident); ok {
// 					v.dataTypes[ident.Name] = true
// 				}
// 			}
// 		}
// 	}

// 	return v
// }

// // matchRoute checks if the request URL matches the route pattern
// func matchRoute(path string, pattern string) bool {
// 	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
// 	urlParts := strings.Split(strings.Trim(path, "/"), "/")

// 	if len(patternParts) != len(urlParts) {
// 		return false
// 	}

// 	for i, part := range patternParts {
// 		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
// 			// This is a path parameter, it matches any value
// 			continue
// 		}
// 		if part != urlParts[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func generateSchema(v interface{}) *openapi3.Schema {
// 	t := reflect.TypeOf(v)
// 	schema := &openapi3.Schema{
// 		Type:       &openapi3.Types{"object"},
// 		Properties: make(openapi3.Schemas),
// 	}

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		jsonTag := field.Tag.Get("json")
// 		if jsonTag == "" {
// 			jsonTag = field.Name
// 		}

// 		var fieldSchema *openapi3.Schema
// 		switch field.Type.Kind() {
// 		case reflect.Int:
// 			fieldSchema = openapi3.NewIntegerSchema()
// 		case reflect.String:
// 			fieldSchema = openapi3.NewStringSchema()
// 		// Add more cases here for other types as needed
// 		default:
// 			// For unsupported types, you might want to log a warning or handle differently
// 			fieldSchema = openapi3.NewSchema()
// 		}

// 		schema.Properties[jsonTag] = openapi3.NewSchemaRef("", fieldSchema)
// 	}

// 	return schema
// }

// // setupOpenAPI initializes the OpenAPI documentation
// // func setupOpenAPI() *openapi3.T {
// // 	openapi := openapi3.T{}
// // 	openapi.OpenAPI = "3.0.0"
// // 	openapi.Info = &openapi3.Info{
// // 		Title:   "Example API",
// // 		Version: "0.0.1",
// // 	}
// // 	openapi.Servers = openapi3.Servers{
// // 		&openapi3.Server{
// // 			URL: "http://localhost:8080",
// // 		},
// // 	}

// // 	// Generate schema for User
// // 	userSchema := generateSchema(User{})

// // 	openapi.Components = &openapi3.Components{
// // 		Schemas: openapi3.Schemas{
// // 			"User": openapi3.NewSchemaRef("", userSchema),
// // 		},
// // 	}

// // 	openapi.Paths = openapi3.NewPaths()
// // 	// for _, route := range routes {
// // 	// 	pathItem := openapi3.PathItem{}
// // 	// 	switch route.Method {
// // 	// 	case http.MethodGet:
// // 	// 		pathItem.Get = route.Operation
// // 	// 	case http.MethodPost:
// // 	// 		pathItem.Post = route.Operation
// // 	// 	case http.MethodPut:
// // 	// 		pathItem.Put = route.Operation
// // 	// 	case http.MethodDelete:
// // 	// 		pathItem.Delete = route.Operation
// // 	// 	}

// // 	// 	openapi.Paths.Set(route.Path, &pathItem)
// // 	// }

// // 	return &openapi
// // }

// // func (api *Client) Put(config Config, handler Handler) Route {
// // 	api.routes[config.Path] = Route{
// // 		Method:  http.MethodPut,
// // 		Path:    config.Path,
// // 		Handler: handler,
// // 	}
// // 	route := api.routes[config.Path]

// // 	// operation := openapi3.Operation{}
// // 	// Responses: openapi3.NewResponses(
// // 	// 	openapi3.WithStatus(200, &openapi3.ResponseRef{
// // 	// 		Value: openapi3.NewResponse().
// // 	// 			WithDescription("OK").
// // 	// 			WithContent(
// // 	// 				openapi3.NewContentWithSchema(
// // 	// 					openapi3.NewStringSchema(),
// // 	// 					[]string{"text/plain"},
// // 	// 				),
// // 	// 			),
// // 	// 	}),
// // 	// ),
// // 	return route
// // }

// // func (api *Client) Get(
// // 	path string,
// // 	handler Handler,
// // ) Route {
// // 	api.routes[path] = Route{
// // 		Method:  http.MethodGet,
// // 		Path:    path,
// // 		Handler: handler,
// // 		Operation: &openapi3.Operation{
// // 			Responses: openapi3.NewResponses(
// // 				openapi3.WithStatus(200, &openapi3.ResponseRef{
// // 					Value: openapi3.NewResponse().
// // 						WithDescription("OK").
// // 						WithContent(
// // 							openapi3.NewContentWithSchema(
// // 								openapi3.NewStringSchema(),
// // 								[]string{"text/plain"},
// // 							),
// // 						),
// // 				}),
// // 			),
// // 		},
// // 	}

// // 	route := api.routes[path]
// // 	api.RegisterHandler(route.Path, route.Method, route.Handler)

// // 	return route
// // }

// type responseWriter struct {
// 	http.ResponseWriter
// 	status int
// 	size   int
// }

// func (rw *responseWriter) WriteHeader(status int) {
// 	rw.status = status
// 	rw.ResponseWriter.WriteHeader(status)
// }

// func (rw *responseWriter) Write(b []byte) (int, error) {
// 	size, err := rw.ResponseWriter.Write(b)
// 	rw.size += size
// 	return size, err
// }

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
// 		next.ServeHTTP(rw, r)

// 		duration := time.Since(start)

// 		remoteAddr := r.RemoteAddr
// 		if ip := r.Header.Get("X-Real-IP"); ip != "" {
// 			remoteAddr = ip
// 		} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
// 			remoteAddr = strings.Split(ip, ", ")[0]
// 		}

// 		log.Printf(
// 			`"%s %s %s" from %s - %d %dB in %v`,
// 			r.Method,
// 			r.URL.RequestURI(),
// 			r.Proto,
// 			remoteAddr,
// 			rw.status,
// 			rw.size,
// 			duration,
// 		)
// 	})
// }
