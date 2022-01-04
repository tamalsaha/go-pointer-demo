package main

import "fmt"

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	// Namespace of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name" protobuf:"bytes,2,opt,name=name"`
}

// COPY on input
// func (ref ObjectReference) WithNamespace(fallback string) ObjectReference {

// Never Copy
// func (ref *ObjectReference) WithNamespace(fallback string) *ObjectReference {

// NO COPY on input (so original changes), but copy on return (so after return the returned one can be changed)
// This is the worst (least expected version).
// func (ref *ObjectReference) WithNamespace(fallback string) ObjectReference {

func (ref *ObjectReference) WithNamespace(fallback string) ObjectReference {
	if ref.Namespace != "" {
		return *ref
	}
	ref.Namespace = fallback
	return *ref
}

func main() {
	r1 := &ObjectReference{
		Namespace: "",
		Name:      "r1",
	}
	b1 := r1.WithNamespace("b1")
	fmt.Println(r1.Namespace)
	fmt.Println(b1.Namespace)
	r1.Namespace = "r2"
	fmt.Println(b1.Namespace)
}
