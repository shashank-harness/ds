package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("ds", func() {
	Title("Directory Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("ds", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})

	})
})

var _ = Service("ds", func() {
	Description("The directory service handles account mgmt")

	HTTP(func() {
		Path("/ds")
	})

	Method("list", func() {
		Description("List all stored bottles")
		Result(CollectionOf(AccountMgmt), func() {
			View("default")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
	//String token, String referer, String gaClientId, String visitorToken
	Method("complete", func() {
		Description("Complete New Signup flow with token")
		Payload(func() {
			Attribute("token", String, "Verification token")
			Attribute("referer", String, "referer")
			Attribute("gaClientId", String, "gaClientId")
			Attribute("visitorToken", String, "visitorToken")
		})
		Result(User, func() {
			View("default")
		})
		HTTP(func() {
			PUT("/complete/{token}")
			Response(StatusOK)
		})
	})
	/*Method("get", func() {
		Description("Get Account by ID")
		Payload(func() {
			Field(1, "id", String, "ID of bottle to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(Account)
		//Error("not_found", NotFound, "Account not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			//Response("not_found", StatusNotFound)
		})
	})
	*/
	Method("demo", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var AccountMgmt = ResultType("application/vnd.harness.ds.accountmgmt", func() {
	Description("AccountMgmt type describes a customer account of company.")
	Reference(Account)
	TypeName("AccountMgmt")

	Attributes(func() {
		Field(1, "id")
		Field(2, "uuid")
		Field(3, "clusterurl")
		Field(4, "accountname")
	})

	View("default", func() {
		Attribute("id")
		Attribute("uuid")
		Attribute("clusterurl")
		Attribute("accountname")
	})

	Required("id", "uuid", "clusterurl", "accountname")
})

var Account = Type("Account", func() {
	Description("Account describes a customer account of Harness.")
	Attribute("id", Int, "ID is an integer type unique id of this DB entity of account.", func() {
		Example(1)
	})
	Attribute("uuid", String, "UUID is the the unique id of this Account", func() {
		Example("kmpySmUISimoRrJL6NL73w")
	})
	Attribute("clusterurl", String, "Tells in which cluster does this account operates in", func() {
		Example("John Doe")
	})
	Attribute("accountname", String, "Name of the Account/Company", func() {
		Example("ABC Pvt Ltd")
	})
	Required("id", "uuid", "clusterurl", "accountname")
})

var User = ResultType("application/vnd.harness.ds.user", func() {
	Description("UserResource describes a user with some account details.")
	Reference(UserType)
	TypeName("UserResource")

	Attributes(func() {
		Field(1, "uuid")
		Field(2, "email")
		Field(3, "name")
		Field(4, "clusterurl")
		Field(5, "accountname")
	})

	View("default", func() {
		Attribute("uuid")
		Attribute("email")
		Attribute("name")
		Attribute("clusterurl")
		Attribute("accountname")
	})

	Required("uuid", "email", "name", "clusterurl", "accountname")
})

var UserType = Type("User", func() {
	Description("User describes a user with some account details.")

	Attribute("uuid", String, "UUID is the the unique id of this Account", func() {
		Example("kmpySmUISimoRrJL6NL73w")
	})
	Attribute("email", String, "email of this user", func() {
		Example("shashank.singh@harness.io")
	})
	Attribute("name", String, "name of this user", func() {
		Example("Shashank Singh")
	})
	Attribute("clusterurl", String, "Tells in which cluster does this account operates in", func() {
		Example("http://localdev.harness.io/prod1")
	})
	Attribute("accountname", String, "Name of the Account/Company", func() {
		Example("Harness")
	})
	Required("uuid", "email", "name", "clusterurl", "accountname")
})
