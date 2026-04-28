[Русский](#about-ru) | [English](#about-en)

<DIV ALIGN="CENTER" ID="about-ru"><H1>О программе</H1></DIV>

`Picdoc` (английская аббревиатура "**PIC**ture **DO**minating **C**olor" - "преобладающий цвет в картинке") - это программа, позволяющая найти преобладающий текст в картинке. Это очень полезный инструмент для тех, кто составляет легенды и присваивает там цвета. Поддерживаются форматы JPG, PNG и ICO.

<DIV ALIGN="CENTER"><H1>Требования</H1></DIV>

* Windows 11

Linux, MacOS и Android (с версии программы v2.0.0) не поддерживаются, однако вы можете собрать `picdoc` из исходного кода.

<DIV ALIGN="CENTER"><h1>Деомнстрация работы</h1></div>

<DIV ALIGN="center"><h3>Картинка для теста:</h3>

<img src="https://github.com/LazataknesSoftware/picdoc/blob/main/resources/fidsico.ico">

</div>

На Windows с `cmd.exe/conhost.exe`:

![cmd.exe/conhost.exe](https://github.com/LazataknesSoftware/picdoc/blob/main/resources/picdoc_proof_windows.PNG)

**Важно!** `conhost.exe` поддерживает только 16 цветов, из-за чего результирующий цвет может быть неточным.


На Android с `Termux`, к которому доступ получен с помощью `PuTTY`:

![PuTTY/Termux](https://github.com/LazataknesSoftware/picdoc/blob/main/resources/picdoc_proof.PNG)

<DIV ALIGN="CENTER"><h1>Опции этой программы</h1></div>

`Picdoc` имеет только одну опцию: `-C`. Эта опция указывает `picdoc` пропускать оттенки белого, серого и черного цветов (`44,44,44`, `200,200,200`,`0,0,0`).

<DIV ALIGN="CENTER"><h1>Благодарности</h1></DIV>

* Огромное спасибо разработчикам библиотеки [color-extractor](https://github.com/marekm4/color-extractor). `Picdoc` не существовал бы без нее!
* Спасибо разработчикам библиотеки [gookit/color](https://github.com/gookit/color). Моя программа без нее не смогла бы отображать цвета в консоли Windows.
* Спасибо разработчикам библиотеки [bestico/ico](https://pkg.go.dev/github.com/mat/besticon/ico). Благодаря ей `picdoc` может обрабатывать ICO-файлы (иконки)!

---

<DIV ALIGN="CENTER" ID="about-en"><H1>About this tool</H1></DIV>

`Picdoc` (abbreviation from **PIC**ture **DO**minating **C**olor) is program that allows to find dominated color in picture. It is very useful tool for those who make legends and need to give to item its colour. It allows to process JPG, PNG and ICO.

<DIV ALIGN="CENTER"><H1>Requirements</H1></DIV>

* Windows 11

Linux, MacOS and Android (since program version v2.0.0) are not supported. However, you can build `picdoc` from source.

<DIV ALIGN="CENTER"><h1>Demo</h1></div>

<DIV ALIGN="center"><h3>Testing picture:</h3>

<img src="https://github.com/LazataknesSoftware/picdoc/blob/main/resources/fidsico.ico">

</div>

On Windows with `cmd.exe/conhost.exe`:

![cmd.exe/conhost.exe](https://github.com/LazataknesSoftware/picdoc/blob/main/resources/picdoc_proof_windows.PNG)

**Important!** `conhost.exe` supports 16 colors only, so color image may be incorrect!


On Android with `Termux`, which accessed by `PuTTY`:

![PuTTY/Termux](https://github.com/LazataknesSoftware/picdoc/blob/main/resources/picdoc_proof.PNG)

<DIV ALIGN="CENTER"><h1>Flags of this tool</h1></div>

`Picdoc` has only 1 flag: `-C`. That flag tells `picdoc` to skip grey, black and white colors (`44,44,44`, `200,200,200`,`0,0,0`).

<DIV ALIGN="CENTER"><h1>Credits</h1></DIV>

* Thanks very much to developers of library called [color-extractor](https://github.com/marekm4/color-extractor). `Picdoc` wouldn't exist without it!
* Thanks to developers of great library called [gookit/color](https://github.com/gookit/color). Without it my program wouldn't show colors in Windows console.
* Thanks to developers of library called [bestico/ico](https://pkg.go.dev/github.com/mat/besticon/ico). Thanks to it `picdoc` can process `.ico` files!
