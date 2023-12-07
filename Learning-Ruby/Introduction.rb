# path: 1. Print Hello World
puts("Hello World!")

p "Hello World!"

# Path: 2.variable.rb
name = "shivam"
puts("Hello " + name + "!")
puts("Hello #{name}!")

age = 18
puts("I am " + age.to_s + " years old.")
puts("I am #{age} years old.")

# Path: 3 User Input
puts("What is your name?")
name = gets()
puts("Hello " + name + "!")
puts("Hello #{name}!")
enrrollement = gets.chomp.to_i
puts("I am " + enrrollement.to_s + " years old.")
puts("I am #{enrrollement} years old.")

# Path: 4 String
puts("Hello World!".length)
puts("Hello World!".include?("World"))
puts("Hello World!".include?("world"))
puts("Hello World!".upcase())
puts("Hello World!".downcase())
puts("Hello World!".capitalize())
puts("Hello World!".reverse())
puts("Hello World!".sub("World", "Ruby"))
puts("Hello World!".concat(" I am Ruby."))
# Multiline string
puts("Hello
World
!")
# Freeze a string
puts("Hello World!".freeze())
# compare string
puts("Hello World!" == "Hello World!")
puts("Hello World!" == "Hello world!")
# String of index
puts("Hello World!"[0])
puts("Hello World!"[1])
puts("Hello World!"[2])
puts("Hello World!"[0, 5])
puts("Hello World!"[0..5])
puts("Hello World!"[0...5])
puts("Hello World!"[-1])
puts("Hello World!"[-2])

# Path: 5 Number
puts(1 + 2)
puts(1 - 2)
puts(1 * 2)
puts(1 / 2)
puts(1.0 / 2)
puts(1 ** 2)
puts(1 % 2)
puts(1 + 2 * 3)

# Path: 6 compare operator
puts(1 == 1)
puts(1 != 1)
puts(1 > 1)
puts(1 >= 1)
puts(1 < 1)
puts(1 <= 1)
# Ternaly operator
puts(1 == 1 ? "Yes" : "No")
puts(1 != 1 ? "Yes" : "No")
puts(1 > 1 ? "Yes" : "No")
puts(1 >= 1 ? "Yes" : "No")
puts(1 < 1 ? "Yes" : "No")
puts(1 <= 1 ? "Yes" : "No")


# Path: 7.array.rb
names = ["shivam", "Bob", "Alice"]
puts(names[0])
puts(names[1])
puts(names[2])
puts(names.length)
names.push("Tom")
puts(names.length)
puts(names[3])
enrrollements = Array.new(3)
puts(enrrollements[0])
puts(enrrollements[1])
puts(enrrollements[2])
puts(enrrollements.length)
enrrollements.push(21)
puts(enrrollements.length)
puts(enrrollements[3])
puts(enrrollements.include?(18))
puts(enrrollements.first())
puts(enrrollements.last())
puts(enrrollements.reverse())
puts(enrrollements.unshift(17))
puts(enrrollements.shift())
puts(enrrollements.pop())
puts(enrrollements.pop())
# delete
enrrollements.delete(18)
# delete_at
enrrollements.delete_at(0)
# each
enrrollements.each{|enrrollement| puts(enrrollement)}
# each_with_index
enrrollements.each_with_index{|enrrollement, index| puts("#{index}: #{enrrollement}")}




# Path: 8.hash.rb
john = {"name" => "shivam", "age" => 18}
puts(john["name"])
puts(john["age"])
john["age"] = 20
puts(john["age"])


# Path: 9.loop.rb
i = 0
while i < 10
    puts(i)
    i = i + 1
    end

# Path: 10 for loop
for i in 0..9
    puts(i)
    end 

# Path: 11 untill and unless
i = 0
until i == 10
    puts(i)
    i = i + 1
    end

# Path: 12 each
names = ["shivam", "Bob", "Alice"]
# names.each do |name|
    puts(name)

# Path: 13 break and next
i = 0
while i < 10
    puts(i)
    i = i + 1
    break if i == 5
    end

# Path: 14 next
i = 0
while i < 10
    i = i + 1
    next if i == 5
    puts(i)
    end

# Path: 15 redo and retry
i = 0
while i < 10
    i = i + 1
    puts(i)
    redo if i == 5
    end

# retry
i = 0
while i < 10
    i = i + 1
    puts(i)
    # retry if i == 5
    end
    

# Path: 16 if else
age = 18
if age >= 18
    puts("You are an adult.")
    else
    puts("You are a child.")
    end

# Path: 17 case
age = 18
case age
    when 0..18
    puts("You are a child.")
    when 18..60
    puts("You are an adult.")
    else
    puts("You are a senior.")
    end

# Path: 18.function.rb
def add(a, b)
    return a + b
    end
puts(add(1, 2))

# Path: 19 switch case
day = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]
case day[0]
    when "Monday"
    puts("Today is Monday.")
    when "Tuesday"
    puts("Today is Tuesday.")
    when "Wednesday"
    puts("Today is Wednesday.")
    when "Thursday"
    puts("Today is Thursday.")
    when "Friday"
    puts("Today is Friday.")
    when "Saturday"
    puts("Today is Saturday.")
    when "Sunday"
    puts("Today is Sunday.")
    end
     


# Path: 20.function.rb
def add(a, b)
    return a + b
    end
puts(add(1, 2))


# Path: 21.class.rb
class Person
    def initialize(name, age)
        @name = name
        @age = age
        end
    def say_hello
        puts("Hello, my name is #{@name}, I am #{@age} years old.")
        end
    end
john = Person.new("shivam", 18)
john.say_hello()


# Path: 22.class_inheritance.rb
class Person
    def initialize(name, age)
        @name = name
        @age = age
        end
    def say_hello
        puts("Hello, my name is #{@name}, I am #{@age} years old.")
        end
    end
class Student < Person
    def initialize(name, age, school)
        super(name, age)
        @school = school
        end
    def say_hello
        super()
        puts("I am studying at #{@school}.")
        end
    end
john = Student.new("shivam", 21, "IIT")
john.say_hello()


# Path: 23.class_inheritance.rb
class Person
    def initialize(name, age)
        @name = name
        @age = age
        end
    def say_hello
        puts("Hello, my name is #{@name}, I am #{@age} years old.")
        end
    end
class Student < Person
    def initialize(name, age, school)
        super(name, age)
        @school = school
        end
    def say_hello
        super()
        puts("I am studying at #{@school}.")
        end
    end
class Teacher < Person
    def initialize(name, age, school)
        super(name, age)
        @school = school
        end
    def say_hello
        super()
        puts("I am teaching at #{@school}.")
        end
    end
john = Student.new("shivam", 21, "IIT")
john.say_hello()
bob = Teacher.new("Shyam", 20, "IIT")
bob.say_hello()

