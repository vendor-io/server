## 🎹 Keyboardify - full-stack e-commerce application 
[![Release](https://img.shields.io/github/v/tag/FoxSaysDerp/keyboardify-server?sort=semver)](https://github.com/FoxSaysDerp/keyboardify-server/releases)

#### Environments
![Go - 1.18.3](https://img.shields.io/badge/Go-1.18.3-00a3cc?style=flat&logo=go&logoColor=white)
![Echo - 4.7.2](https://img.shields.io/badge/Echo-4.7.2-00a6c6?style=flat)
![Gorm - 1.23.8](https://img.shields.io/badge/Gorm-1.23.8-529ee2?style=flat)

---
### Repos ⛓
- [Client](https://github.com/foxsaysderp/keyboardify-client)
- [Server](https://github.com/foxsaysderp/keyboardify-server)

### Docs 🗃

todo 💀

### TODO 📝
- Move to PostgreSQL/MySQL
- Implement Loadbalancer
- Implement Redis
- Create rank system (User/Producer/Admin)
- Overhaul current controllers to include a dedicated error system, including internal code and message: 
```
{
   "error": 520021,
   "message": "You have insufficient permissions to view this order"
}
```