// Code generated by tempo, DO NOT EDIT.
package choreography

import runtime "tempo/runtime"

// Projection of choreography foo
func foo_A(env *runtime.Env) {
	env.Send(10, "B")
}
func foo_B(env *runtime.Env) {
	var value *runtime.Async = env.Recv("A")
	_ = value
	var another int = value.Get().(int)
	_ = another
}
