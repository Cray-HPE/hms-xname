# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.0] - 2024-11-22

### Changed

- Updated go to 1.23

## [1.3.0] - 2023-08-29

### Added
* Added xname support for virtual nodes

## [1.2.0] - 2023-08-29

### Changed
* Switched to Go 1.20

### Added
* Add FromStringToStruct helper function that utilizes generics to reduce boilerplate code when parsing xnames.

## [1.1.0] - 2022-06-23

### Added

- Added the `Xname` interface to the `xnames` package. It is meant to provide a generic way to store or pass a xnames structure, such as to a function that can handle multiple different xnames types.

  All of the xnames structures now implement the following functions:
  - `String() string` - build the xname string for the given structure
  - `Type() xnametypes.HMSType` - get the Type enumeration of the structure.
  - `ParentInterface() Xname` - get the parent structure but with the `Xname` interface.

### Changed

- Corrected the name of the `xnames` package, as it was originally set to `xname`.
- Changed the `xnames.FromString` method to use the GenericInterface type. 

## [1.0.2] - 2022-02-09

### Changed

- Fixed number of args in xname for ChassisBMCNic

## [1.0.1] - 2022-01-21

### Changed

- purged hms-base references.

## [1.0.0] - 2021-12-13

### Added

- CASMHMS-5180: Added new xnames package to manipulate xnames via Go structures, instead of using `fmt.Sprintf` or regular expressions to manipulate the string representation of xnames. 

### Changed

- CASMHMS-5180: Moved HMSTypes and related functions to the new hms-xname repo under the xnametypes package.
- CASMHMS-5180: Added functions to expose the contents of the hmsCompRecognitionTable
- CASMHMS-4667: GetHMSCompParent now returns s0 for CDUs and Cabinets.
- CASMHMS-4668: Added missing MgmtHLSwitchEnclosure type
- CASMHMS-4668: Corrected the parent type for MgmtHLSwitch to be MgmtHLSwitchEnclosure
- CASMHMS-4668: Corrected the parent type for RouterTORFpga to be RouterTOR

