package main

import "fmt"

var uniqueID int

type Object struct {
	id int
}

type ObjectPool struct {
	unlocked []*Object
	locked   []*Object
}

func (self *ObjectPool) AcquireObject() *Object {
	var object *Object
	if len(self.unlocked) <= 0 {
		object = &Object{uniqueID}
		uniqueID++
	} else {
		object, self.unlocked = self.unlocked[len(self.unlocked)-1], self.unlocked[:len(self.unlocked)-1]
	}
	self.locked = append(self.locked, object)
	return object
}

func (self *ObjectPool) ReleaseObject(object *Object) {
	for i, lockedObject := range self.locked {
		if lockedObject.id == object.id {
			self.locked = append(self.locked[:i], self.locked[i+1:]...)
			self.unlocked = append(self.unlocked, object)
			return
		}
	}
}

func main() {
	objectPool := &ObjectPool{}
	object1 := objectPool.AcquireObject()
	fmt.Println(object1.id)

	object2 := objectPool.AcquireObject()
	fmt.Println(object2.id)
	objectPool.ReleaseObject(object2)

	object3 := objectPool.AcquireObject()
	fmt.Println(object3.id)
}
