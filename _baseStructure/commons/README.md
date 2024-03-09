- In this place, you can implement various functions to assist you in performing common tasks—consider them as helpers
- Common functions can be directly called from any location
- If a **Domain** or **Feature** imports anything from this location, and you want the **Feature** to be accessible through 
the `importFeature` command without the risk of missing package errors, **DON'T FORGET** to include them in the `features/yourFeature/dependency.json` file
