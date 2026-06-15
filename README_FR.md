[English](README.md) | **Français**

# Snapchat Memories Processor

Traitez facilement toutes vos Memories Snapchat exportées tout en conservant les overlays, les dates et les données de localisation.

Snapchat exporte les Memories sous forme de fichiers multimédias accompagnés d'un fichier de métadonnées séparé. Cet outil reconstruit vos memories en restaurant les overlays et en réinjectant, si vous le souhaitez, les coordonnées GPS dans les photos et vidéos générées.

## Exporter ses Memories Snapchat

Cette section expliquera comment exporter vos Memories et métadonnées depuis Snapchat.

*Captures d'écran et GIFs à venir.*

---

## Installation

Téléchargez la dernière version depuis la page Releases.

Des binaires sont disponibles pour :

* Windows
* macOS
* Linux

Aucune installation n'est nécessaire. Téléchargez simplement l'exécutable correspondant à votre système.

---

## Prérequis

### FFmpeg

FFmpeg est nécessaire pour traiter les vidéos et les overlays.

#### Windows

Téléchargez FFmpeg depuis le site officiel puis ajoutez-le au PATH.

#### macOS

```bash
brew install ffmpeg
```

#### Linux

Ubuntu / Debian :

```bash
sudo apt install ffmpeg
```

Arch Linux :

```bash
sudo pacman -S ffmpeg
```

Vérifiez l'installation :

```bash
ffmpeg -version
```

---

## Optionnel : conservation des données GPS

Par défaut, les coordonnées GPS ne sont pas réécrites dans les fichiers générés.

Pour conserver les données de localisation, installez ExifTool puis utilisez l'option `--gps`.

### Windows

Téléchargez ExifTool depuis le site officiel puis ajoutez-le au PATH.

### macOS

```bash
brew install exiftool
```

### Linux

Ubuntu / Debian :

```bash
sudo apt install libimage-exiftool-perl
```

Arch Linux :

```bash
sudo pacman -S perl-image-exiftool
```

Vérifiez l'installation :

```bash
exiftool -ver
```

---

## Préparer l'export

Placez toutes les archives ZIP Snapchat dans un même dossier.

Exemple :

```text
exports/
├── mydata.zip
├── mydata-2.zip
├── mydata-3.zip
└── mydata-4.zip
```

Les archives peuvent être placées n'importe où sur votre ordinateur.

---

## Utilisation

### Utilisation simple

```bash
smp process -i ./exports
```

L'option `-i` permet de spécifier le dossier contenant les archives Snapchat.

### Conserver les données GPS

```bash
smp process -i ./exports --gps
```

Cette fonctionnalité nécessite ExifTool.

### Nombre de workers personnalisé

```bash
smp process -i ./exports -w 8
```

Le nombre de workers détermine combien de fichiers sont traités simultanément.

Dans la plupart des cas, la valeur par défaut est suffisante.

---

## Exemple de sortie

```text
Total media  : 2502
Videos       : 1898
Images       : 604
With overlay : 435

Processed    : 2502
Failed       : 0

Completed in 42.7s
```

---

## Dossier de sortie

Par défaut, tous les médias traités sont enregistrés dans le dossier `output`.

Exemple :

```text
output/
├── 2020-07-24_094EC87A-main.jpg
├── 2020-07-24_42180C76-main.mp4
└── ...
```

---

## Traitement des erreurs

Si certains fichiers ne peuvent pas être traités, un fichier de log est automatiquement généré :

```text
output/errors.log
```

Ce fichier contient les détails des erreurs rencontrées.

---

## Licence

Ce projet est sous Licence MIT.
