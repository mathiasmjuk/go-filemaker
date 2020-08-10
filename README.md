# Getting started

## Starting/destroying session
``` go
conn, err := filemaker.Connect("https://example.com", "database", "username", "password")
if err != nil {
  fmt.Println("Error:", err.Error())
  return
}
defer conn.Close()
```

## Perform find
``` go
command := filemaker.NewFindCommand(
  filemaker.NewFindRequest(
    filemaker.NewFindCriterion("fieldname", "=matchthis"),
  ),
  //...
)

records, err := conn.PerformFind("layoutname", command)
if err != nil {
  switch err.(type) {
  case *ErrorNotFound:
    fmt.Println("Records not found!")
  default:
    fmt.Println("Unknown error:", err.Error())
  }

  return
}

for _, record := range records {
  fmt.Println(record.GetField("fieldname"))
}
```

### Omit
``` go
command := filemaker.NewFindCommand(
  filemaker.NewFindRequest(
    filemaker.NewFindCriterion("fieldname", "somethinglikethis"),
  ),
  filemaker.NewFindRequest(
    filemaker.NewFindCriterion("otherfieldname", "=notsomethinglikethis"),
  ).Omit(), //Omit request
)
```

### Limit
``` go
command := filemaker.NewFindCommand(
  //...
).SetLimit(10)
```

### Offset
``` go
command := filemaker.NewFindCommand(
  //...
).SetOffset(10)
```

### Limit and offset (chaining)
``` go
command := filemaker.NewFindCommand(
  //...
).SetLimit(10).SetOffset(10)
```

## Records

### Create
``` go
record := filemaker.CreateRecord("layoutname")
record.SetField("fieldname", "data")

err = conn.Commit(&record) //Need to pass record by pointer
if err != nil {
  fmt.Println("Error:", err.Error())
}

fmt.Println("Record ID:", record.ID)
```

### Edit
``` go
record.SetField("fieldname", "new data")

err := conn.Commit(&record) //Need to pass record by pointer
if err != nil {
  fmt.Println("Error:", err.Error())
}
```

### Delete
``` go
err = conn.Delete(record.Layout, record.ID)
if err != nil {
  fmt.Println("Error:", err.Error())
}
```