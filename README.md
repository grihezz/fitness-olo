***Интернет-порал для любителей здоровго образа жтзни "OLO"***

**Приступая к работе**
Инструкция о том, как получить копию этого ПО и запустить его на локальном компьютере с целью разработки и тестирования. Подробную информацию о развертывании ПО в условиях эксплуатации см. в разделе «Развертывание».

Предварительные условия
Чтобы запустить сервисы нужно:
1. Скачать IDE дял работы с Golang(рекомендуем Goland от JetBrains)
2. Скачть компилятор go
3. Скачть утилиту make, protoc.git.docker,docker-compose 

Установка
Ссылки на ПО приведенные выше:

Windows:
1. https://www.jetbrains.com/ru-ru/go/ - Goland
2. https://gcc.gnu.org/ - компилятор gcc
3. https://gnuwin32.sourceforge.net/packages/make.htm - утилита Make 
4. https://grpc.io/docs/protoc-installation/ - утилита protoc
5. https://git-scm.com/downloads - git
6. https://docs.docker.com/get-docker/ - Docker
7. https://docs.docker.com/compose/install/ - docker-compose 

MacOS:
1. https://www.jetbrains.com/ru-ru/go/ - Goland
2. В терминале: brew install gcc - компилятор gcc
3. В терминале: brew install make - утилита Make
4. В терминале: brew install make - утилита Make
5. https://git-scm.com/downloads - git
6. https://docs.docker.com/get-docker/ - Docker
7. В треминале: brew install docker-compose - docker-compose 

Пошаговая инструкция, которая поможет войти в среду разработки.

1. Скачать IDE дял работы с Golang(рекомендуем Goland от JetBrains)
2. На экране выбора или создания проека выбираем "Get from VCS"
3. Вводим ссылку на репозиторий (git@github.com:IKBO-13-21/OLO-backend.git)
   
Развертывание
1. Выполнить команды из Makefile в следующем порядке:
   1. install-deps
   2. get-deps
   3. vendor-proto
   4. generate
2. Есть 2 способа запуститиь проект
   * Локально, запистув каждый микросервис через go run main.go, находясь в дериктории cmd или нажав зеленый треугольник в верхней правой части IDE
   * Через Docker с помощью docker-compose up --build находсь в корне проекта

Создано с помощью
* golang - Это современный, статически типизированный, компилируемый язык программирования, разработанный в Google, который обеспечивает высокую производительность и поддержку параллельных вычислений.
* grpc - Это высокопроизводительный, открытый и общий фреймворк для удаленного вызова процедур (RPC), который позволяет разработчикам создавать масштабируемые и эффективные сервисы.
* make - Это система автоматизации, которая управляет процессом сборки программного обеспечения, позволяя разработчикам автоматизировать компиляцию и связывание исходного кода.
* jwt - Это открытый стандарт (RFC 7519), который определяет способ безопасной передачи информации между двумя сторонами в виде JSON объекта. JWT часто используется для аутентификации и передачи информации между сервером и клиентом.
* crypto - Это пакет в Go, предоставляющий криптографические примитивы, такие как хеширование, шифрование и цифровые подписи, для обеспечения безопасности данных.
* mysql - Это система управления реляционными базами данных, которая использует язык SQL для управления данными. MySQL широко используется для разработки веб-приложений и других приложений, требующих хранения и обработки больших объемов данных.
* slog - Это библиотека для логирования в Go, предназначенная для упрощения процесса логирования в приложениях, позволяя разработчикам легко добавлять логирование с различными уровнями детализации и форматами вывода.
* protobuf - Это библиотека для сериализации структурированных данных, разработанная Google. Она позволяет определить структуру данных в простом текстовом формате, который затем может быть использован для генерации кода на различных языках программирования. Protobuf обеспечивает эффективную сериализацию данных, что делает его идеальным для передачи данных между серверами и клиентами в распределенных системах

Управление версиями
Для управления версиями мы используем SemVer. Информацию о доступных версиях см. в тегах в этом репозитории.

## Авторы
* Бабоха Григорий - golang-developer /teamlead
* Агафонов Андрей - golang-developer /dev-ops

Copyright &copy; Бабоха Григорий, 2024

The MIT License (MIT)
