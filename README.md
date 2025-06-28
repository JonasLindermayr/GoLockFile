# GoLockFile 🔒

Ein einfaches CLI-Tool zum Verschlüsseln und Entschlüsseln von Dateien mit AES-GCM Verschlüsselung.  
---

## 📦 Features

- AES-256 GCM Dateiverschlüsselung
- Verschlüsselung und Entschlüsselung via CLI
- Möglichkeit, die Originaldatei nach der Verschlüsselung/Entschlüsselung zu löschen
- Vergabe eines neuen Dateinamens für verschlüsselte Dateien
- Einfache Bedienung über Befehlszeilenparameter

---

## 🛠️ Installation

### Voraussetzungen

- [Go](https://go.dev/doc/install) Version 1.18 oder höher

---

## 📑 Build-Anleitung

Download
```bash
https://github.com/JonasLindermayr/GoLockFile/releases/tag/v1.0.0
```

### macOS / Linux

```bash
make build
```

### Windows

```bash
go build -o build/GoLockFile.exe ./cmd
```

## 🚀 Verwendung

### Hilfe anzeigen

```bash
./golockfile help
```

### Datei verschlüsseln

```bash
./golockfile encrypt -file <Dateipfad> -p <Passwort>
```

#### Optionale Parameter:

```bash
-n <Name> → Neuen Dateinamen vergeben
-w → Originaldatei nach der Verschlüsselung löschen

./golockfile encrypt -file geheim.txt -p meinpasswort -n verschluesselt -w
```

### Datei entschlüsseln

```bash
./golockfile decrypt -file <Dateipfad> -p <Passwort>
```

#### Optionale Parameter:

```bash
-w → Verschlüsselte Datei nach der Entschlüsselung löschen

./golockfile decrypt -file geheim.txt.locked -p meinpasswort -w
```

## 📌 Hinweise

-- Verschlüsselte Dateien erhalten die Endung .locked

-- Die Verschlüsselung verwendet AES-256 GCM mit einem SHA-256 Hash des Passworts als Key

-- Passwort darf nicht vergessen werden — es gibt keine Möglichkeit zur Wiederherstellung!

## Autor

Jonas Lindermayr
GitHub: https://github.com/JonasLindermayr
