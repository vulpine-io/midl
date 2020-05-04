/*
Package midlmock provides a set of configurable mock
implementations of all the interface types present in the
midl package.

Usage

Each mock implementation has a property for each function
defined on the interface it mocks.  For example, the
midl.Serializer interface defines the function "Serialize".
The midlmock.Serializer implementation has the public
property "SerializeFunc" which will be called by the mock's
Serialize function.  This allows the use of any arbitrary
test functionality within the mock.

Mock implementations which have self returning or builder
style function calls will automatically return a reference
to their current instance.
*/
package midlmock
