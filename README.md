
```markdown
Agency CLI

مدیریت آژانس‌ها از طریق خط فرمان.
```
## اجرا

```bash
go run main.go
```

## دستورات

| دستور | توضیح |
|-------|-------|
| `list` | لیست آژانس‌ها |
| `get` | مشاهده آژانس |
| `create` | ایجاد آژانس |
| `edit` | ویرایش آژانس |
| `status` | وضعیت سیستم |
| `exit` | خروج |

## فیلتر منطقه

```bash
go run main.go list -region=Tehran
```

## ساختار پروژه

```
├── main.go
├── model/
├── repository/
├── service/
├── handler/
└── utils/
```
