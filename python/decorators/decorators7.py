def without_decorator():

    def add_messages(start_message, end_message):
        def amend_func(old_add_func):
            def new_add_func(x, y):
                print(start_message)
                result = old_add_func(x, y)
                print(end_message)
                return result
            return new_add_func
        return amend_func
    
    def add(x, y):
        return x + y
    
    decorator_func = add_messages("Starting add!", "Add completed!")
    add = decorator_func(add)

    print("1st result is", add(2, 8))

def with_decorator():

    def add_messages(start_message, end_message):
        def amend_func(func):
            def new_add_func(x, y):
                print(start_message)
                result = func(x, y)
                print(end_message)
                return result
            return new_add_func
        return amend_func
    
    @add_messages("Starting add!!", "Add completed!!")
    def add(x, y):
        return x + y

    print("2nd result is", add(7, 4))

print("Run first...")
without_decorator()

print("Run second...")
with_decorator()
