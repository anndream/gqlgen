// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package chat

import (
	context "context"
	strconv "strconv"
	sync "sync"
	time "time"

	graphql "github.com/vektah/gqlgen/graphql"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers}
}

type Resolvers interface {
	Mutation_post(ctx context.Context, text string, username string, roomName string) (Message, error)
	Query_room(ctx context.Context, name string) (*Chatroom, error)

	Subscription_messageAdded(ctx context.Context, roomName string) (<-chan Message, error)
}

type Message struct {
	ID        string
	Text      string
	CreatedBy string
	CreatedAt time.Time
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	data := ec._query(op.Selections)
	ec.wg.Wait()

	return &graphql.Response{
		Data:   data,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	data := ec._mutation(op.Selections)
	ec.wg.Wait()

	return &graphql.Response{
		Data:   data,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) <-chan *graphql.Response {
	events := make(chan *graphql.Response, 10)

	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	eventData := ec._subscription(op.Selections)
	if ec.Errors != nil {
		events <- &graphql.Response{
			Data:   graphql.Null,
			Errors: ec.Errors,
		}
		close(events)
	} else {
		go func() {
			for data := range eventData {
				ec.wg.Wait()
				events <- &graphql.Response{
					Data:   data,
					Errors: ec.Errors,
				}
				time.Sleep(20 * time.Millisecond)
			}
		}()
	}
	return events
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	wg        sync.WaitGroup
}

var chatroomImplementors = []string{"Chatroom"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _chatroom(sel []query.Selection, it *Chatroom) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, chatroomImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Chatroom")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name

			out.Values[i] = graphql.MarshalString(res)
		case "messages":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Messages

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler
				tmp1 = ec._message(field.Selections, &res[idx1])
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var messageImplementors = []string{"Message"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _message(sel []query.Selection, it *Message) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, messageImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Message")
		case "id":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.ID

			out.Values[i] = graphql.MarshalID(res)
		case "text":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Text

			out.Values[i] = graphql.MarshalString(res)
		case "createdBy":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.CreatedBy

			out.Values[i] = graphql.MarshalString(res)
		case "createdAt":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.CreatedAt

			out.Values[i] = graphql.MarshalTime(res)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var mutationImplementors = []string{"Mutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _mutation(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, mutationImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "post":
			badArgs := false
			var arg0 string
			if tmp, ok := field.Args["text"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			var arg1 string
			if tmp, ok := field.Args["username"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg1 = tmp2
			}
			var arg2 string
			if tmp, ok := field.Args["roomName"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg2 = tmp2
			}
			if badArgs {
				continue
			}
			res, err := ec.resolvers.Mutation_post(ec.ctx, arg0, arg1, arg2)
			if err != nil {
				ec.Error(err)
				continue
			}

			out.Values[i] = ec._message(field.Selections, &res)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _query(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, queryImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "room":
			badArgs := false
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.Query_room(ec.ctx, arg0)
				if err != nil {
					ec.Error(err)
					return
				}

				if res == nil {
					out.Values[i] = graphql.Null
				} else {
					out.Values[i] = ec._chatroom(field.Selections, res)
				}
			}(i, field)
		case "__schema":
			badArgs := false
			if badArgs {
				continue
			}
			res := ec.introspectSchema()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Schema(field.Selections, res)
			}
		case "__type":
			badArgs := false
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				tmp2, err := graphql.UnmarshalString(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := ec.introspectType(arg0)

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var subscriptionImplementors = []string{"Subscription"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _subscription(sel []query.Selection) <-chan graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, subscriptionImplementors, ec.variables)

	if len(fields) != 1 {
		ec.Errorf("must subscribe to exactly one stream")
		return nil
	}

	var field = fields[0]
	channel := make(chan graphql.Marshaler, 1)
	switch field.Name {
	case "messageAdded":
		badArgs := false
		var arg0 string
		if tmp, ok := field.Args["roomName"]; ok {
			tmp2, err := graphql.UnmarshalString(tmp)
			if err != nil {
				badArgs = true
			}
			arg0 = tmp2
		}
		if badArgs {
			return nil
		}
		results, err := ec.resolvers.Subscription_messageAdded(ec.ctx, arg0)
		if err != nil {
			ec.Error(err)
			return nil
		}

		go func() {
			for res := range results {
				var out graphql.OrderedMap
				var messageRes graphql.Marshaler
				messageRes = ec._message(field.Selections, &res)
				out.Add(field.Alias, messageRes)
				channel <- &out
			}
		}()
	default:
		panic("unknown field " + strconv.Quote(field.Name))
	}

	return channel
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, it *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __DirectiveImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "locations":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Locations()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler
				tmp1 = graphql.MarshalString(res[idx1])
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "args":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, it *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __EnumValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "isDeprecated":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, it *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __FieldImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "args":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "type":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "isDeprecated":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, it *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __InputValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "type":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "defaultValue":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.DefaultValue()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, it *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __SchemaImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Types()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "queryType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.QueryType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "mutationType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.MutationType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "subscriptionType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.SubscriptionType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "directives":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Directives()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Directive(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, it *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __TypeImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Kind()

			out.Values[i] = graphql.MarshalString(res)
		case "name":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Name()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "description":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "fields":
			badArgs := false
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := graphql.UnmarshalBoolean(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := it.Fields(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Field(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "interfaces":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.Interfaces()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "possibleTypes":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.PossibleTypes()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "enumValues":
			badArgs := false
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				tmp2, err := graphql.UnmarshalBoolean(tmp)
				if err != nil {
					badArgs = true
				}
				arg0 = tmp2
			}
			if badArgs {
				continue
			}
			res := it.EnumValues(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___EnumValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "inputFields":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.InputFields()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "ofType":
			badArgs := false
			if badArgs {
				continue
			}
			res := it.OfType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var parsedSchema = schema.MustParse("type Chatroom {\n    name: String!\n    messages: [Message!]!\n}\n\ntype Message {\n    id: ID!\n    text: String!\n    createdBy: String!\n    createdAt: Time!\n}\n\ntype Query {\n    room(name:String!): Chatroom\n}\n\ntype Mutation {\n    post(text: String!, username: String!, roomName: String!): Message!\n}\n\ntype Subscription {\n    messageAdded(roomName: String!): Message!\n}\n\nschema {\n    query: Query\n    mutation: Mutation\n    subscription: Subscription\n}\n\nscalar Time\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}
