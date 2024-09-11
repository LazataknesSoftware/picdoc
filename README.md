<DIV ALIGN="CENTER"><H1>About this tool</H1></DIV>

`Picdoc` (abbreviation from **PIC**ture **DO**minating **C**olor) is program that allows to find dominated color in picture. It is very useful tool for those who make legends and need to give to item its colour.

<DIV ALIGN="CENTER"><H1>Acceptable Formats</H1></DIV>

It allows to process `JPG (JPEG)`, `PNG` and `ICO`.

<DIV ALIGN="CENTER"><H1>Requirements</H1></DIV>

* Windows 10 1607
* Android: [Termux](https://github.com/termux-app/termux)

Linux and MacOS are not supported. However, you can build `picdoc` from source.

<DIV ALIGN="CENTER"><h1>Demo</h1></div>

<DIV ALIGN="center"><h3>Testing picture:</h3>

<img src="D:\MYSOFT~1\USEFUL~1\PICDOC\FIDSICO.ICO">

</div>

On Windows with `cmd.exe/conhost.exe`:

![cmd.exe/conhost.exe](D:\MYSOFT~1\USEFUL~1\PICDOC\picdoc_proof_windows.png)

**IMPORTANT!** `conhost.exe` supports 16 colors only, so color image may be incorrect!


On Android `(PuTTY/Termux)`:

![PuTTY/Termux](D:\MYSOFT~1\USEFUL~1\PICDOC\picdoc_proof.png)

<DIV ALIGN="CENTER"><h1>Flags of this tool</h1></div>

`Picdoc` has only 1 flag: `-C`. That flag tells `picdoc` to skip grey, black and white colors (`44,44,44`, `200,200,200`,`0,0,0`).

<DIV ALIGN="CENTER"><H1>Errors explaination</H1></DIV>

* `I did not find file. Run picdoc with filename.` - this error means that file not found. Also that error shows when you ran `picdoc` without filename.
* `I can't recognize file. Is it picture (JPG, JPEG, PNG, ICO)?` - probably, you sent non-image file (for example, `.js`, `.txt`, `.mp4`) to `picdoc`.

<DIV ALIGN="CENTER"><h1>Credits</h1></DIV>

* Thanks very much to developers of library called [color-extractor](https://github.com/marekm4/color-extractor). `Picdoc` wouldn't exist without it!
* Thanks to developers of great library called [gookit/color](https://github.com/gookit/color). Without it my program wouldn't show colors in Windows console.
* Thanks to developers of library called [bestico/ico](https://pkg.go.dev/github.com/mat/besticon/ico). Thanks to it `picdoc` can process `.ico` files!

<DIV ALIGN="CENTER"><h1>Do you have idea?</h1></DIV>

Great! You can make pull request (if you have idea) or issue (if you face problem).
