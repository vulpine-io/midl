package midl

// EmptyAdapter creates an Adapter with no default settings or serializers
func EmptyAdapter() Adapter {
	return &adapter{}
}
