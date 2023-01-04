# MiniBinMac

## Мини корзина в трее для MacOs

![Preview](./docs/img/emptybin_panel.png)

![Preview](./docs/img/img1.png)

## Меню приложения

![Menu](docs/img/menubin.png)

### Кнопка "очистить" очищает корзину

![Menu-Clear](docs/img/menu_bin_clear.png)

### Кнопка "Выход" закрывает приложение

![Menu-Quit](docs/img/menubin_exit.png)


## Состояния

![полная корзина](docs/img/bin.png)
![пустая корзина](docs/img/empty_bin.png)


## Установка

### Автоматическая установка

```bash
git clone https://github.com/MiniBinMac
cd MiniBinMac
make install
```

### Ручная установка

```zsh
git clone https://github.com/MiniBinMac
cd MiniBinMac
chmod +x install.sh
./install.sh
```

## Запуск

### Исполняемый файл назодится в `MiniBinMac/bin/minibinmac`

### Установщик `install.sh` переносит его по пути `/usr/local/bin/` и `opt/local/bin/` , вы можете добавить исполняемый фал в **объекты входа**

![Объекты входа](docs/img/autostart_settings.png)

![Добавление исполняемого файлв в объекты входа](docs/img/autostart_objects.png)

### теперь приложение будет автоматически запускаться при входе в систему

### Поддержка


| OS        | Tested    |  
| ----------- | -----------|
| **Mac OS Ventura**      | :white_check_mark:   |
| **cat**     | :x:   |
| **clear**   | :x:    |

<p align="center">2023 © <a href="https://github.com/Avdushin" target="_blank">AVDUSHIN</a></p>