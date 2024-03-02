- The **Domain** represents your primary business model or entity
- Define your main object models or properties for your business here, including database models, DTOs (Data Transfer Objects), etc
- Keep this package as straightforward as possible. Avoid including any code that is not directly related to the model itself
- If a **Feature** imports anything from this location, and you want the **Feature** to be accessible through the `importFeature` command 
without the risk of missing package errors, **DON'T FORGET** to include them in the `features/yourFeature/dependency.json` file
