Why using fiber framework?
- beacuse of its various benefits
- Zero Allocation
- As a rule of thumb, you must only use context values within the handler, and you must not keep any references. 
- the wrong code :
- ``` func handler(c *fiber.Ctx) error {
    go asyncFunc(c) // ❌ BAD! Do not do this
    return nil
}

func asyncFunc(c *fiber.Ctx) {
    fmt.Println(c.Params("id")) // Might panic or give wrong data
}```
- the right way: 
- ```func handler(c *fiber.Ctx) error {
    id := c.Params("id") // ✅ Safe to store
    go asyncFunc(id)     // ✅ Now you're using plain string, not context
    return nil
}

func asyncFunc(id string) {
    fmt.Println("Async task for:", id)
}```