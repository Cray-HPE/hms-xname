# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2022-06-23

### Added

- Added the `GenericXname` interface to the `xnames` package. It is meant to provide a generic way to store or pass a xnames structure, such as to a function that can handle multiple different xnames types.

  All of the xnames structures now implement the following functions:
  - `String() string` - build the xname string for the given structure
  - `Type() xnametypes.HMSType` - get the Type enumeration of the structure.
  - `ParentGeneric() GenericXname` - get the xname 
  - `IsController() bool` - returns true if the structure is a controller type, i.e. that  would host a Redfish entry point

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

