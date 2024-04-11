**How to implement error handling**

*steps*
     -Saving error on context
     -Sending error flag from a function and return 
     -wating for error flag and if error occur call formerror function for error handling
     -format error handling read error from the context and create appropriate http response for user.


// ethio-phone pattern  
`^(/+251)?(97)[1-59]/d{8}$`
regexp.Compile(phonepatter) // regexp.MustCompile



**How to implement pagenation**
*steps*
    -create a get request that accept page and per-page query parameter.
    -Binding a queryparmeter to a struct object (customer metadata struct )
    -per-page=limit and page=page number
     
     limit=per_page 
     offset=(page-1)*per_page


**Field validation**
*steps*
    - validate a user field using a reg.exp
    - username `^[a-zA-Z0-9_]{3,20}$`
    - password `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`