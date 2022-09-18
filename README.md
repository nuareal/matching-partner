# Matching Partner

An api for finding partner company information.

# Rest Api

Current api functions:
- GetMatchingPartners
- GetPartner

# Usage

- Run main.go
- Open http://localhost:8080/swagger/index.html

Consumer example:
```json
{
  "address": {
    "lat": 52.50,
    "lon": 13.37
  },
  "floorArea": 10,
  "material": "wood",
  "phoneNumber": "+4203874573949 "
}
```

# Future Improvements
- Add register costumer / save costumer info
- Use database for costumer and partner info
- Add unit tests

# Extra Information
Assumption:
- Company name is unique
