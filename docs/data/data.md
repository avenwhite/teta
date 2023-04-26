# Модель предметной области
<!-- Логическая модель, содержащая бизнес-сущности предметной области, атрибуты и связи между ними. 
Подробнее: https://confluence.mts.ru/pages/viewpage.action?pageId=375782602

Используется диаграмма классов UML. Документация: https://plantuml.com/class-diagram 
-->
```plantuml
@startuml
class Conference {
  -id: int
  -title: string
  -description: string
  -location: string
  -start_date: datetime
  -end_date: datetime
  -speakers: List[Speaker]
  -program: Program
  -feedback: List[Feedback]
  -partners: List[Partner]
  -audience: List[Audience]
}

class Speaker {
  -id: int
  -name: string
  -email: string
  -phone: string
  -bio: string
  -topics: List[string]
  -presentations: List[Presentation]
}

class Presentation {
  -id: int
  -title: string
  -description: string
  -file: File
  -reviewers: List[Reviewer]
}

class Reviewer {
  -id: int
  -name: string
  -email: string
  -phone: string
  -expertise: List[string]
  -reviews: List[Review]
}

class Review {
  -id: int
  -score: int
  -comments: string
  -reviewer: Reviewer
}

class Program {
  -id: int
  -sessions: List[Session]
}

class Session {
  -id: int
  -title: string
  -start_time: datetime
  -end_time: datetime
  -presentations: List[Presentation]
  -translation: List[Translation]
}

class Feedback {
  -id: int
  -score: int
  -comments: string
  -audience: Audience
}
class Partner {
  -id: int
  -name: string
  -logo: File
  -website: string
}

class Audience {
  -id: int
  -name: string
  -email: string
  -phone: string
  -affiliation: string
  -feedback: List[Feedback]
}

class Translation {
  -id: int
  -name: string
  -type: string
  -size: int
}

Conference -> Speaker
Conference -> Program
Conference -> Feedback
Conference -> Partner
Conference -> Audience
Conference -> Translation
Speaker -> Presentation
Presentation -> Reviewer
Reviewer -> Review
Program -> Session
Session -> Presentation
Feedback -> Audience
@enduml
```
