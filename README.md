**English** | [Français](README_FR.md)

# Snapchat Memories Processor

Easily process all your exported Snapchat Memories while preserving overlays, timestamps and locations.

Snapchat exports Memories as raw media files and a separate metadata file. This tool reconstructs your memories by restoring overlays and optionally writing GPS information back into your photos and videos.

## Export your Snapchat Memories

> This section will explain how to export your Snapchat Memories and metadata from Snapchat.

First, login to your snapchat account (phone or pc) and head to `Settings` -> `My Data`

You should then check `Export your Memories` **AND** `Export JSON files`. 

![select items to export](./assets/snap_exports_options.png)

After that you will be prompted to select the time range of the export, this is up to you.

Depending on the size of the export, you should pretty quickly receive an email from snapchat, saying that your export is ready.
You can then head to the same page and you'll se the `Your exports` section. When unfolding it, you should see something similar to this:

![download exports](./assets/snap_exports_download.png)

Download **all** the `zip` archives and you'll be ready to start!

---

## Installation

Download the latest release for your operating system from the [Releases page](https://github.com/EliasLd/snap-memories-processor/releases).

Prebuilt binaries are available for:

* Windows
* macOS
* Linux

> [!NOTE]
> No installation is required. Simply download the appropriate executable and place it wherever you want to use it.

---

## Requirements

### FFmpeg

**FFmpeg** is required for processing videos and overlays. If you don't have it installed, you can't use this tool

#### Windows

The easiest and fastest way to install it is running this command in a command line or powershell:

```shell
winget install ffmpeg
```

> [!WARNING]
> If winget is not recognize as a command, don't panic and install it via the **Microsoft Store**. Just search for "App Installer".

#### macOS

```bash
brew install ffmpeg
```

#### Linux

Ubuntu / Debian:

```bash
sudo apt install ffmpeg
```

Arch Linux:

```bash
sudo pacman -S ffmpeg
```

Verify the installation:

```bash
ffmpeg -version
```

---

## Optional: GPS Metadata Support

By default, GPS coordinates are not written back into the generated media files.

If you want to preserve location data, install ExifTool and use the `--gps` flag.

### Windows

Again, the easiest way is to install it via `winget`, running this command:

```
winget install -e --id OliverBetz.ExifTool
```

### macOS

```bash
brew install exiftool
```

### Linux

Ubuntu / Debian:

```bash
sudo apt install libimage-exiftool-perl
```

Arch Linux:

```bash
sudo pacman -S perl-image-exiftool
```

Verify the installation:

```bash
exiftool -ver
```

---

## Preparing Your Export

Place all Snapchat export `ZIP` files in a directory.

Example:

```text
exports/
├── mydata.zip
├── mydata-2.zip
├── mydata-3.zip
└── mydata-4.zip
```

> [!NOTE]
> The archives may be stored anywhere on your system.

---

## Usage

### Command Line (inside a terminal) 

#### Basic Usage

```bash
smp process -i ./exports
```

The `-i` flag specifies the directory containing your Snapchat export archives.

#### Preserve GPS Metadata

```bash
smp process -i ./exports --gps
```

> [!NOTE]
> This requires ExifTool to be installed.

#### Custom Worker Count

```bash
smp process -i ./exports -w 8
```

The worker count controls how many files are processed simultaneously.

> [!TIP]
> In most cases, the default value is sufficient.

---

## Example Output

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

## Output Directory

Processed media files are written to the `output` directory by default.

Example:

```text
output/
├── 2020-07-24_094EC87A-main.jpg
├── 2020-07-24_42180C76-main.mp4
└── ...
```

---

## Error Logs

If one or more files fail to process, an error log is generated automatically:

```text
output/errors.log
```

This file contains detailed information about each failure and can be useful for troubleshooting.

## License

This project is under MIT License

## Contributing

Want to make this tool better ? Don't hesistate to fork it and open a PR. I'll be happy to read it :)
