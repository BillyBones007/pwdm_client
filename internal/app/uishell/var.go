package uishell

// Main page
const mainPage string = `
Please use 'exit' or Ctrl+D to exit this program.
Please use 'register' to create a new account.
Please use 'login' to log in to your account.
And use 'help' for more information.
-------------------------------------------------`

// Help text
const helpText string = `*** Клиент менеджера паролей ***

*** Шифрует полученные данные и сохраняет их на сервере.

*** Позволяет работать со следующими данными:
	1. Пары логин/пароль
	2. Данные банковских карт
	3. Произвольные текстовые данные
	4. Произвольные бинарные данные


*** Поддерживает автодополнение команд. Выбрать команду из списка предложенных
    можно с помощью TAB.
	
*** Для входа в приложение необходимо авторизоваться/зарегестрироваться.

*** Адрес сервера задается в конфигурационном файле config.json

*** Основные команды:

	Для неавторизованного пользователя: 
		register         - регистрация нового пользователя
		login            - войти в учетную запись
		exit             - выход из приложения
		
	Для авторизованного пользователя:
		logout           - выйти из учетной записи
		exit             - выход из приложения
		save
		   |
		   |-lp          - сохранить пару login/password
		   |-card        - сохранить данные банковской карты
		   |-text        - сохранить произвольный текст (вводится в поле ввода)
		   |    |
		   |    - [-f]   - сохранить текстовые данные из файла
		   |-binary      - сохранить бинарные данные из файла

		edit
		   |
		   |-lp          - редактировать пару login/password
		   |-card        - редактировать данные банковской карты
		   |-text        - редактировать текстовые данные (вводится в поле ввода)
		   |    |
		   |    - [-f]   - редактировать текстовые данные из файла
		   |-binary      - редактировать бинарные данные из файла
		
		show
		   |
		   |-lp          - вывести на экран список сохраненных данных с парами login/password
		   |-card        - вывести на экран список сохраненных данных по банковским картам
		   |-text        - вывести на экран список сохраненных текстовых данных
		   |-binary      - вывести на экран список сохраненных бинарных данных
		   
		get
		   |
		   |-lp [id]     - получить данные пары login/password по id записи
		   |-card [id]   - получить данные банковской карты по id записи
		   |-text [id]   - получить текстовые данные по id записи (сохраняется в файл)
		   |-binary [id] - получить бинарные данные по id записи (сохраняется в файл)
		   
		delete
		   |
		   |-lp [id]     - удалить запись с парой login/password по id записи
		   |-card [id]   - удалить запись с банковской картой по id записи
		   |-text [id]   - удалить запись с текстовыми данными по id записи
		   |-binary [id] - удалить запись с бинарными данными по id записи`

// Hint
const hint string = `NOTE:  The Title, Comment and Tag fields are optional.`
