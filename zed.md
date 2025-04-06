# Zed Editor - Горячие клавиши

## Основные операции редактирования

| Действие                  | macOS       | Windows/Linux | Команда Zed                  |
|---------------------------|-------------|---------------|------------------------------|
| Отмена действия           | `⌘ + Z`     | `Ctrl + Z`    | `editor::Undo`               |
| Повтор действия           | `⌘ + ⇧ + Z` | `Ctrl + ⇧ + Z`| `editor::Redo`               |
| Вырезать                  | `⌘ + X`     | `Ctrl + X`    | `editor::Cut`                |
| Копировать                | `⌘ + C`     | `Ctrl + C`    | `editor::Copy`               |
| Вставить                  | `⌘ + V`     | `Ctrl + V`    | `editor::Paste`              |
| Удалить                   | `Delete`    | `Delete`      | `editor::Delete`             |
| Удалить назад             | `Backspace` | `Backspace`   | `editor::Backspace`          |

## Навигация по тексту

| Действие                  | macOS       | Windows/Linux | Команда Zed                  |
|---------------------------|-------------|---------------|------------------------------|
| В начало строки           | `⌘ + ←`     | `Home`        | `editor::MoveToBeginningOfLine` |
| В конец строки            | `⌘ + →`     | `End`         | `editor::MoveToEndOfLine`    |
| На слово влево            | `⌥ + ←`     | `Ctrl + ←`    | `editor::MoveToPreviousWordStart` |
| На слово вправо           | `⌥ + →`     | `Ctrl + →`    | `editor::MoveToNextWordEnd`  |
| В начало файла            | `⌘ + ↑`     | `Ctrl + Home` | `editor::MoveToBeginning`    |
| В конец файла             | `⌘ + ↓`     | `Ctrl + End`  | `editor::MoveToEnd`          |
| Прокрутка к курсору       | `^ + L`     | `Ctrl + L`    | `editor::ScrollCursorCenter` |

## Git-интеграция

| Действие                  | macOS         | Windows/Linux   | Команда Zed          |
|---------------------------|---------------|-----------------|----------------------|
| Отменить изменения        | `⌥ + ⌘ + Z`   | `Alt + Ctrl + Z`| `git::Restore`       |
| Добавить в индекс         | `⌘ + Y`       | `Ctrl + Y`      | `git::StageAndNext`  |
| Убрать из индекса         | `⌘ + ⇧ + Y`   | `Ctrl + ⇧ + Y`  | `git::UnstageAndNext`|
| Показать blame            | `⌥ + ⌘ + G B` | `Alt + Ctrl + G B` | `editor::ToggleGitBlame` |

## Как использовать этот файл

1. Сохраните как `zed-hotkeys.md`
2. Для конвертации в PDF:
   - Через VS Code с плагином Markdown PDF
   - На [dillinger.io](https://dillinger.io/)
   - Через pandoc: `pandoc zed-hotkeys.md -o zed-hotkeys.pdf`

> Примечания:  
> - `⌘` = Command (Mac)  
> - `⌥` = Option (Mac) / Alt (Win)  
> - `⇧` = Shift  
> - `^` = Control
