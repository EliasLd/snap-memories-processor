[English](README.md) | **Français**

# Snapchat Memories Processor

Traitez facilement toutes vos Memories Snapchat exportées tout en conservant les overlays, les dates et les données de localisation.

Snapchat exporte les Memories sous forme de fichiers multimédias bruts accompagnés d'un fichier de métadonnées séparé. Cet outil reconstruit vos Memories en restaurant les overlays et, si vous le souhaitez, en réécrivant les informations GPS directement dans vos photos et vidéos.

## Exporter vos Memories Snapchat

> Cette section expliquera comment exporter vos Memories et leurs métadonnées depuis Snapchat.

Commencez par vous connecter à votre compte Snapchat (sur téléphone ou ordinateur), puis rendez-vous dans `Paramètres` -> `Mes données`.

Vous devez ensuite cocher `Exporter mes souvenirs` **ET** `Exporter des fichiers JSON`.

![sélection des éléments à exporter](./assets/snap_exports_options_fr.png)

Vous serez ensuite invité à choisir la période à exporter. Ce choix dépend entièrement de vos besoins.

Selon la taille de l'export, vous devriez recevoir assez rapidement un e-mail de Snapchat indiquant que votre export est prêt.

Vous pourrez alors retourner sur la même page et ouvrir la section `Vos exports`. Vous devriez voir quelque chose de similaire à ceci :

![téléchargement des exports](./assets/snap_exports_download_fr.png)

Téléchargez **toutes** les archives `zip` et vous serez prêt à commencer.

---

## Installation

Téléchargez la dernière version de l'outil correspondant à votre système d'exploitation depuis la [page des releases](https://github.com/EliasLd/snap-memories-processor/releases).

Des exécutables précompilés sont disponibles pour :

* Windows
* macOS
* Linux

> [!NOTE]
> Aucune installation n'est nécessaire. Téléchargez simplement l'exécutable correspondant à votre système et placez-le où vous souhaitez l'utiliser.

---

## Prérequis

### Installation de `FFmpeg`

**FFmpeg** est nécessaire pour traiter les vidéos et les overlays. Si FFmpeg n'est pas installé, cet outil ne pourra pas fonctionner.

#### Sur Windows

Le moyen le plus simple et le plus rapide consiste à exécuter cette commande dans un terminal ou PowerShell :

```shell
winget install ffmpeg
```

> [!WARNING]
> Si la commande `winget` n'est pas reconnue, ne vous inquiétez pas. Installez-la via le **Microsoft Store** en recherchant simplement « App Installer ».

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

## Optionnel : Conservation des données GPS

Par défaut, les coordonnées GPS ne sont pas réécrites dans les fichiers générés.

Si vous souhaitez conserver les données de localisation, installez ExifTool et utilisez l'option `--gps`.

### Sur Windows

Là encore, la méthode la plus simple consiste à l'installer avec `winget` :

```bash
winget install -e --id OliverBetz.ExifTool
```

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

## Préparer votre export

Placez toutes les archives Snapchat `ZIP` dans un dossier.

Exemple :

```text
exports/
├── mydata.zip
├── mydata-2.zip
├── mydata-3.zip
└── mydata-4.zip
```

> [!NOTE]
> Les archives peuvent être stockées n'importe où sur votre système.

---

## Utilisation

### Ligne de commande (dans un terminal)

#### Utilisation de base

```bash
smp process -i ./exports
```

L'option `-i` permet d'indiquer le dossier contenant vos archives d'export Snapchat.

#### Conserver les données GPS

```bash
smp process -i ./exports --gps
```

> [!NOTE]
> Cette fonctionnalité nécessite qu'ExifTool soit installé.

#### Nombre personnalisé de threads

```bash
smp process -i ./exports -w 8
```

Le nombre de threads détermine combien de fichiers peuvent être traités simultanément.

> [!TIP]
> Dans la plupart des cas, la valeur par défaut est suffisante.

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

Les fichiers traités sont écrits dans le dossier `output` par défaut.

Exemple :

```text
output/
├── 2020-07-24_094EC87A-main.jpg
├── 2020-07-24_42180C76-main.mp4
└── ...
```

---

## Journal d'erreurs

Si un ou plusieurs fichiers ne peuvent pas être traités, un journal d'erreurs est automatiquement généré :

```text
output/errors.log
```

Ce fichier contient des informations détaillées sur chaque erreur et peut être utile pour diagnostiquer un problème.

## Licence

Ce projet est distribué sous licence MIT.

## Contribution

Vous souhaitez améliorer cet outil ? N'hésitez pas à le fork et à ouvrir une Pull Request. Je serai ravi de la lire :)

