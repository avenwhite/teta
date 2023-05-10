# Компонентная архитектура
<!-- Состав и взаимосвязи компонентов системы между собой и внешними системами с указанием протоколов, ключевые технологии, используемые для реализации компонентов.
Диаграмма контейнеров C4 и текстовое описание. 
Подробнее: https://confluence.mts.ru/pages/viewpage.action?pageId=375783368
-->
## Контейнерная диаграмма

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()
AddElementTag("microService", $shape=EightSidedShape(), $bgColor="CornflowerBlue", $fontColor="white", $legendText="microservice")
AddElementTag("storage", $shape=RoundedBoxShape(), $bgColor="lightSkyBlue", $fontColor="white")


System_Boundary(c, "MTS helloconf") {
   Container(app, "веб-приложение", "html, JavaScript, Angular", "Портал системы конфиеренций helloconf")
   Container(facade_service, "Фасадный интеграционный сервис", "Java, Spring Boot", "Сервис управления поиском и хранением данных системы", $tags = "microService")
   Container(storage_service, "Сервис хранения данных", "Java, Spring Boot", "Сервис ", $tags = "microService")
   Container(сonference_service, "Сервис расписаний", "Java, Spring Boot", "Сервис ", $tags = "microService")    
   Container(storage_service_db, "Storage Service", "PostgreSQL", "Сервис ", $tags = "storage")            
   Container(сonference_service_db, "Storage Service", "PostgreSQL", "Сервис ", $tags = "storage")
   Container(feedback_service, "Сервис обратной связи", "Java, Spring Boot", "Сервис ", $tags = "microService")    
   Container(feedback_service_db, "Storage Service", "PostgreSQL", "Сервис ", $tags = "storage")   
   Container(program_service, "Сервис бронирования конференций", "Java, Spring Boot", "Сервис ", $tags = "microService")    
   Container(program_service_db, "Storage Service", "PostgreSQL", "Сервис ", $tags = "storage")          
}
Person(usr, "Пользователи", "Участники конферениции\Администраторы\Модераторы")
System_Ext(online, "Сервис онлайн транслирования", "Хранение и запись конференций")

Rel(usr, app, "Работа с порталом", "HTTPS")
Rel(app,facade_service,"REST")
Rel(facade_service,сonference_service,"REST")
Rel(facade_service,storage_service,"REST")
Rel(facade_service,feedback_service,"REST")
Rel(facade_service,program_service,"REST")
Rel(program_service,program_service_db,"SQL")
Rel(feedback_service,feedback_service_db,"SQL")
Rel(storage_service,storage_service_db,"SQL")
Rel(сonference_service,сonference_service_db,"SQL")


Rel(c, online, "Транслирование потового видео для онлайн участников", "HTTPS")
@enduml
```
