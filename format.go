//Package cff provides a parser for the CFF format.
//Go to https://alexander-lindner.github.io/go-cff/ for a full documentation.
package cff

import (
	"net/url"
	"time"
)

//URL represents a basic url. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsurl for more information.
type URL struct {
	*url.URL
}

//Date represents a date
type Date time.Time

//Country represents a country
type Country string

const (
	COUNTRY_AD Country = "AD"
	COUNTRY_AE Country = "AE"
	COUNTRY_AF Country = "AF"
	COUNTRY_AG Country = "AG"
	COUNTRY_AI Country = "AI"
	COUNTRY_AL Country = "AL"
	COUNTRY_AM Country = "AM"
	COUNTRY_AO Country = "AO"
	COUNTRY_AQ Country = "AQ"
	COUNTRY_AR Country = "AR"
	COUNTRY_AS Country = "AS"
	COUNTRY_AT Country = "AT"
	COUNTRY_AU Country = "AU"
	COUNTRY_AW Country = "AW"
	COUNTRY_AX Country = "AX"
	COUNTRY_AZ Country = "AZ"
	COUNTRY_BA Country = "BA"
	COUNTRY_BB Country = "BB"
	COUNTRY_BD Country = "BD"
	COUNTRY_BE Country = "BE"
	COUNTRY_BF Country = "BF"
	COUNTRY_BG Country = "BG"
	COUNTRY_BH Country = "BH"
	COUNTRY_BI Country = "BI"
	COUNTRY_BJ Country = "BJ"
	COUNTRY_BL Country = "BL"
	COUNTRY_BM Country = "BM"
	COUNTRY_BN Country = "BN"
	COUNTRY_BO Country = "BO"
	COUNTRY_BQ Country = "BQ"
	COUNTRY_BR Country = "BR"
	COUNTRY_BS Country = "BS"
	COUNTRY_BT Country = "BT"
	COUNTRY_BV Country = "BV"
	COUNTRY_BW Country = "BW"
	COUNTRY_BY Country = "BY"
	COUNTRY_BZ Country = "BZ"
	COUNTRY_CA Country = "CA"
	COUNTRY_CC Country = "CC"
	COUNTRY_CD Country = "CD"
	COUNTRY_CF Country = "CF"
	COUNTRY_CG Country = "CG"
	COUNTRY_CH Country = "CH"
	COUNTRY_CI Country = "CI"
	COUNTRY_CK Country = "CK"
	COUNTRY_CL Country = "CL"
	COUNTRY_CM Country = "CM"
	COUNTRY_CN Country = "CN"
	COUNTRY_CO Country = "CO"
	COUNTRY_CR Country = "CR"
	COUNTRY_CU Country = "CU"
	COUNTRY_CV Country = "CV"
	COUNTRY_CW Country = "CW"
	COUNTRY_CX Country = "CX"
	COUNTRY_CY Country = "CY"
	COUNTRY_CZ Country = "CZ"
	COUNTRY_DE Country = "DE"
	COUNTRY_DJ Country = "DJ"
	COUNTRY_DK Country = "DK"
	COUNTRY_DM Country = "DM"
	COUNTRY_DO Country = "DO"
	COUNTRY_DZ Country = "DZ"
	COUNTRY_EC Country = "EC"
	COUNTRY_EE Country = "EE"
	COUNTRY_EG Country = "EG"
	COUNTRY_EH Country = "EH"
	COUNTRY_ER Country = "ER"
	COUNTRY_ES Country = "ES"
	COUNTRY_ET Country = "ET"
	COUNTRY_FI Country = "FI"
	COUNTRY_FJ Country = "FJ"
	COUNTRY_FK Country = "FK"
	COUNTRY_FM Country = "FM"
	COUNTRY_FO Country = "FO"
	COUNTRY_FR Country = "FR"
	COUNTRY_GA Country = "GA"
	COUNTRY_GB Country = "GB"
	COUNTRY_GD Country = "GD"
	COUNTRY_GE Country = "GE"
	COUNTRY_GF Country = "GF"
	COUNTRY_GG Country = "GG"
	COUNTRY_GH Country = "GH"
	COUNTRY_GI Country = "GI"
	COUNTRY_GL Country = "GL"
	COUNTRY_GM Country = "GM"
	COUNTRY_GN Country = "GN"
	COUNTRY_GP Country = "GP"
	COUNTRY_GQ Country = "GQ"
	COUNTRY_GR Country = "GR"
	COUNTRY_GS Country = "GS"
	COUNTRY_GT Country = "GT"
	COUNTRY_GU Country = "GU"
	COUNTRY_GW Country = "GW"
	COUNTRY_GY Country = "GY"
	COUNTRY_HK Country = "HK"
	COUNTRY_HM Country = "HM"
	COUNTRY_HN Country = "HN"
	COUNTRY_HR Country = "HR"
	COUNTRY_HT Country = "HT"
	COUNTRY_HU Country = "HU"
	COUNTRY_ID Country = "ID"
	COUNTRY_IE Country = "IE"
	COUNTRY_IL Country = "IL"
	COUNTRY_IM Country = "IM"
	COUNTRY_IN Country = "IN"
	COUNTRY_IO Country = "IO"
	COUNTRY_IQ Country = "IQ"
	COUNTRY_IR Country = "IR"
	COUNTRY_IS Country = "IS"
	COUNTRY_IT Country = "IT"
	COUNTRY_JE Country = "JE"
	COUNTRY_JM Country = "JM"
	COUNTRY_JO Country = "JO"
	COUNTRY_JP Country = "JP"
	COUNTRY_KE Country = "KE"
	COUNTRY_KG Country = "KG"
	COUNTRY_KH Country = "KH"
	COUNTRY_KI Country = "KI"
	COUNTRY_KM Country = "KM"
	COUNTRY_KN Country = "KN"
	COUNTRY_KP Country = "KP"
	COUNTRY_KR Country = "KR"
	COUNTRY_KW Country = "KW"
	COUNTRY_KY Country = "KY"
	COUNTRY_KZ Country = "KZ"
	COUNTRY_LA Country = "LA"
	COUNTRY_LB Country = "LB"
	COUNTRY_LC Country = "LC"
	COUNTRY_LI Country = "LI"
	COUNTRY_LK Country = "LK"
	COUNTRY_LR Country = "LR"
	COUNTRY_LS Country = "LS"
	COUNTRY_LT Country = "LT"
	COUNTRY_LU Country = "LU"
	COUNTRY_LV Country = "LV"
	COUNTRY_LY Country = "LY"
	COUNTRY_MA Country = "MA"
	COUNTRY_MC Country = "MC"
	COUNTRY_MD Country = "MD"
	COUNTRY_ME Country = "ME"
	COUNTRY_MF Country = "MF"
	COUNTRY_MG Country = "MG"
	COUNTRY_MH Country = "MH"
	COUNTRY_MK Country = "MK"
	COUNTRY_ML Country = "ML"
	COUNTRY_MM Country = "MM"
	COUNTRY_MN Country = "MN"
	COUNTRY_MO Country = "MO"
	COUNTRY_MP Country = "MP"
	COUNTRY_MQ Country = "MQ"
	COUNTRY_MR Country = "MR"
	COUNTRY_MS Country = "MS"
	COUNTRY_MT Country = "MT"
	COUNTRY_MU Country = "MU"
	COUNTRY_MV Country = "MV"
	COUNTRY_MW Country = "MW"
	COUNTRY_MX Country = "MX"
	COUNTRY_MY Country = "MY"
	COUNTRY_MZ Country = "MZ"
	COUNTRY_NA Country = "NA"
	COUNTRY_NC Country = "NC"
	COUNTRY_NE Country = "NE"
	COUNTRY_NF Country = "NF"
	COUNTRY_NG Country = "NG"
	COUNTRY_NI Country = "NI"
	COUNTRY_NL Country = "NL"
	COUNTRY_NO Country = "NO"
	COUNTRY_NP Country = "NP"
	COUNTRY_NR Country = "NR"
	COUNTRY_NU Country = "NU"
	COUNTRY_NZ Country = "NZ"
	COUNTRY_OM Country = "OM"
	COUNTRY_PA Country = "PA"
	COUNTRY_PE Country = "PE"
	COUNTRY_PF Country = "PF"
	COUNTRY_PG Country = "PG"
	COUNTRY_PH Country = "PH"
	COUNTRY_PK Country = "PK"
	COUNTRY_PL Country = "PL"
	COUNTRY_PM Country = "PM"
	COUNTRY_PN Country = "PN"
	COUNTRY_PR Country = "PR"
	COUNTRY_PS Country = "PS"
	COUNTRY_PT Country = "PT"
	COUNTRY_PW Country = "PW"
	COUNTRY_PY Country = "PY"
	COUNTRY_QA Country = "QA"
	COUNTRY_RE Country = "RE"
	COUNTRY_RO Country = "RO"
	COUNTRY_RS Country = "RS"
	COUNTRY_RU Country = "RU"
	COUNTRY_RW Country = "RW"
	COUNTRY_SA Country = "SA"
	COUNTRY_SB Country = "SB"
	COUNTRY_SC Country = "SC"
	COUNTRY_SD Country = "SD"
	COUNTRY_SE Country = "SE"
	COUNTRY_SG Country = "SG"
	COUNTRY_SH Country = "SH"
	COUNTRY_SI Country = "SI"
	COUNTRY_SJ Country = "SJ"
	COUNTRY_SK Country = "SK"
	COUNTRY_SL Country = "SL"
	COUNTRY_SM Country = "SM"
	COUNTRY_SN Country = "SN"
	COUNTRY_SO Country = "SO"
	COUNTRY_SR Country = "SR"
	COUNTRY_SS Country = "SS"
	COUNTRY_ST Country = "ST"
	COUNTRY_SV Country = "SV"
	COUNTRY_SX Country = "SX"
	COUNTRY_SY Country = "SY"
	COUNTRY_SZ Country = "SZ"
	COUNTRY_TC Country = "TC"
	COUNTRY_TD Country = "TD"
	COUNTRY_TF Country = "TF"
	COUNTRY_TG Country = "TG"
	COUNTRY_TH Country = "TH"
	COUNTRY_TJ Country = "TJ"
	COUNTRY_TK Country = "TK"
	COUNTRY_TL Country = "TL"
	COUNTRY_TM Country = "TM"
	COUNTRY_TN Country = "TN"
	COUNTRY_TO Country = "TO"
	COUNTRY_TR Country = "TR"
	COUNTRY_TT Country = "TT"
	COUNTRY_TV Country = "TV"
	COUNTRY_TW Country = "TW"
	COUNTRY_TZ Country = "TZ"
	COUNTRY_UA Country = "UA"
	COUNTRY_UG Country = "UG"
	COUNTRY_UM Country = "UM"
	COUNTRY_US Country = "US"
	COUNTRY_UY Country = "UY"
	COUNTRY_UZ Country = "UZ"
	COUNTRY_VA Country = "VA"
	COUNTRY_VC Country = "VC"
	COUNTRY_VE Country = "VE"
	COUNTRY_VG Country = "VG"
	COUNTRY_VI Country = "VI"
	COUNTRY_VN Country = "VN"
	COUNTRY_VU Country = "VU"
	COUNTRY_WF Country = "WF"
	COUNTRY_WS Country = "WS"
	COUNTRY_YE Country = "YE"
	COUNTRY_YT Country = "YT"
	COUNTRY_ZA Country = "ZA"
	COUNTRY_ZM Country = "ZM"
	COUNTRY_ZW Country = "ZW"
)

//License represents the license type
type License string

const (
	LICENSE_0BSD                              License = "0BSD"
	LICENSE_AAL                               License = "AAL"
	LICENSE_Abstyles                          License = "Abstyles"
	LICENSE_Adobe2006                         License = "Adobe-2006"
	LICENSE_AdobeGlyph                        License = "Adobe-Glyph"
	LICENSE_ADSL                              License = "ADSL"
	LICENSE_AFL11                             License = "AFL-1.1"
	LICENSE_AFL12                             License = "AFL-1.2"
	LICENSE_AFL20                             License = "AFL-2.0"
	LICENSE_AFL21                             License = "AFL-2.1"
	LICENSE_AFL30                             License = "AFL-3.0"
	LICENSE_Afmparse                          License = "Afmparse"
	LICENSE_AGPL10                            License = "AGPL-1.0"
	LICENSE_AGPL10only                        License = "AGPL-1.0-only"
	LICENSE_AGPL10orlater                     License = "AGPL-1.0-or-later"
	LICENSE_AGPL30only                        License = "AGPL-3.0-only"
	LICENSE_AGPL30orlater                     License = "AGPL-3.0-or-later"
	LICENSE_AGPL30                            License = "AGPL-3.0"
	LICENSE_Aladdin                           License = "Aladdin"
	LICENSE_AMDPLPA                           License = "AMDPLPA"
	LICENSE_AML                               License = "AML"
	LICENSE_AMPAS                             License = "AMPAS"
	LICENSE_ANTLRPD                           License = "ANTLR-PD"
	LICENSE_ANTLRPDfallback                   License = "ANTLR-PD-fallback"
	LICENSE_Apache10                          License = "Apache-1.0"
	LICENSE_Apache11                          License = "Apache-1.1"
	LICENSE_Apache20                          License = "Apache-2.0"
	LICENSE_APAFML                            License = "APAFML"
	LICENSE_APL10                             License = "APL-1.0"
	LICENSE_Apps2p                            License = "App-s2p"
	LICENSE_APSL10                            License = "APSL-1.0"
	LICENSE_APSL11                            License = "APSL-1.1"
	LICENSE_APSL12                            License = "APSL-1.2"
	LICENSE_APSL20                            License = "APSL-2.0"
	LICENSE_Artistic10                        License = "Artistic-1.0"
	LICENSE_Artistic10cl8                     License = "Artistic-1.0-cl8"
	LICENSE_Artistic10Perl                    License = "Artistic-1.0-Perl"
	LICENSE_Artistic20                        License = "Artistic-2.0"
	LICENSE_Bahyph                            License = "Bahyph"
	LICENSE_Barr                              License = "Barr"
	LICENSE_Beerware                          License = "Beerware"
	LICENSE_BitTorrent10                      License = "BitTorrent-1.0"
	LICENSE_BitTorrent11                      License = "BitTorrent-1.1"
	LICENSE_blessing                          License = "blessing"
	LICENSE_BlueOak100                        License = "BlueOak-1.0.0"
	LICENSE_Borceux                           License = "Borceux"
	LICENSE_BSD1Clause                        License = "BSD-1-Clause"
	LICENSE_BSD2Clause                        License = "BSD-2-Clause"
	LICENSE_BSD2ClausePatent                  License = "BSD-2-Clause-Patent"
	LICENSE_BSD2ClauseViews                   License = "BSD-2-Clause-Views"
	LICENSE_BSD2ClauseFreeBSD                 License = "BSD-2-Clause-FreeBSD"
	LICENSE_BSD2ClauseNetBSD                  License = "BSD-2-Clause-NetBSD"
	LICENSE_BSD3Clause                        License = "BSD-3-Clause"
	LICENSE_BSD3ClauseAttribution             License = "BSD-3-Clause-Attribution"
	LICENSE_BSD3ClauseClear                   License = "BSD-3-Clause-Clear"
	LICENSE_BSD3ClauseLBNL                    License = "BSD-3-Clause-LBNL"
	LICENSE_BSD3ClauseModification            License = "BSD-3-Clause-Modification"
	LICENSE_BSD3ClauseNoMilitaryLicense       License = "BSD-3-Clause-No-Military-License"
	LICENSE_BSD3ClauseNoNuclearLicense        License = "BSD-3-Clause-No-Nuclear-License"
	LICENSE_BSD3ClauseNoNuclearLicense2014    License = "BSD-3-Clause-No-Nuclear-License-2014"
	LICENSE_BSD3ClauseNoNuclearWarranty       License = "BSD-3-Clause-No-Nuclear-Warranty"
	LICENSE_BSD3ClauseOpenMPI                 License = "BSD-3-Clause-Open-MPI"
	LICENSE_BSD4Clause                        License = "BSD-4-Clause"
	LICENSE_BSD4ClauseShortened               License = "BSD-4-Clause-Shortened"
	LICENSE_BSD4ClauseUC                      License = "BSD-4-Clause-UC"
	LICENSE_BSDProtection                     License = "BSD-Protection"
	LICENSE_BSDSourceCode                     License = "BSD-Source-Code"
	LICENSE_BSL10                             License = "BSL-1.0"
	LICENSE_BUSL11                            License = "BUSL-1.1"
	LICENSE_bzip2105                          License = "bzip2-1.0.5"
	LICENSE_bzip2106                          License = "bzip2-1.0.6"
	LICENSE_CUDA10                            License = "C-UDA-1.0"
	LICENSE_CAL10                             License = "CAL-1.0"
	LICENSE_CAL10CombinedWorkException        License = "CAL-1.0-Combined-Work-Exception"
	LICENSE_Caldera                           License = "Caldera"
	LICENSE_CATOSL11                          License = "CATOSL-1.1"
	LICENSE_CCBY10                            License = "CC-BY-1.0"
	LICENSE_CCBY20                            License = "CC-BY-2.0"
	LICENSE_CCBY25                            License = "CC-BY-2.5"
	LICENSE_CCBY25AU                          License = "CC-BY-2.5-AU"
	LICENSE_CCBY30                            License = "CC-BY-3.0"
	LICENSE_CCBY30AT                          License = "CC-BY-3.0-AT"
	LICENSE_CCBY30DE                          License = "CC-BY-3.0-DE"
	LICENSE_CCBY30NL                          License = "CC-BY-3.0-NL"
	LICENSE_CCBY30US                          License = "CC-BY-3.0-US"
	LICENSE_CCBY40                            License = "CC-BY-4.0"
	LICENSE_CCBYNC10                          License = "CC-BY-NC-1.0"
	LICENSE_CCBYNC20                          License = "CC-BY-NC-2.0"
	LICENSE_CCBYNC25                          License = "CC-BY-NC-2.5"
	LICENSE_CCBYNC30                          License = "CC-BY-NC-3.0"
	LICENSE_CCBYNC30DE                        License = "CC-BY-NC-3.0-DE"
	LICENSE_CCBYNC40                          License = "CC-BY-NC-4.0"
	LICENSE_CCBYNCND10                        License = "CC-BY-NC-ND-1.0"
	LICENSE_CCBYNCND20                        License = "CC-BY-NC-ND-2.0"
	LICENSE_CCBYNCND25                        License = "CC-BY-NC-ND-2.5"
	LICENSE_CCBYNCND30                        License = "CC-BY-NC-ND-3.0"
	LICENSE_CCBYNCND30DE                      License = "CC-BY-NC-ND-3.0-DE"
	LICENSE_CCBYNCND30IGO                     License = "CC-BY-NC-ND-3.0-IGO"
	LICENSE_CCBYNCND40                        License = "CC-BY-NC-ND-4.0"
	LICENSE_CCBYNCSA10                        License = "CC-BY-NC-SA-1.0"
	LICENSE_CCBYNCSA20                        License = "CC-BY-NC-SA-2.0"
	LICENSE_CCBYNCSA20FR                      License = "CC-BY-NC-SA-2.0-FR"
	LICENSE_CCBYNCSA20UK                      License = "CC-BY-NC-SA-2.0-UK"
	LICENSE_CCBYNCSA25                        License = "CC-BY-NC-SA-2.5"
	LICENSE_CCBYNCSA30                        License = "CC-BY-NC-SA-3.0"
	LICENSE_CCBYNCSA30DE                      License = "CC-BY-NC-SA-3.0-DE"
	LICENSE_CCBYNCSA30IGO                     License = "CC-BY-NC-SA-3.0-IGO"
	LICENSE_CCBYNCSA40                        License = "CC-BY-NC-SA-4.0"
	LICENSE_CCBYND10                          License = "CC-BY-ND-1.0"
	LICENSE_CCBYND20                          License = "CC-BY-ND-2.0"
	LICENSE_CCBYND25                          License = "CC-BY-ND-2.5"
	LICENSE_CCBYND30                          License = "CC-BY-ND-3.0"
	LICENSE_CCBYND30DE                        License = "CC-BY-ND-3.0-DE"
	LICENSE_CCBYND40                          License = "CC-BY-ND-4.0"
	LICENSE_CCBYSA10                          License = "CC-BY-SA-1.0"
	LICENSE_CCBYSA20                          License = "CC-BY-SA-2.0"
	LICENSE_CCBYSA20UK                        License = "CC-BY-SA-2.0-UK"
	LICENSE_CCBYSA21JP                        License = "CC-BY-SA-2.1-JP"
	LICENSE_CCBYSA25                          License = "CC-BY-SA-2.5"
	LICENSE_CCBYSA30                          License = "CC-BY-SA-3.0"
	LICENSE_CCBYSA30AT                        License = "CC-BY-SA-3.0-AT"
	LICENSE_CCBYSA30DE                        License = "CC-BY-SA-3.0-DE"
	LICENSE_CCBYSA40                          License = "CC-BY-SA-4.0"
	LICENSE_CCPDDC                            License = "CC-PDDC"
	LICENSE_CC010                             License = "CC0-1.0"
	LICENSE_CDDL10                            License = "CDDL-1.0"
	LICENSE_CDDL11                            License = "CDDL-1.1"
	LICENSE_CDL10                             License = "CDL-1.0"
	LICENSE_CDLAPermissive10                  License = "CDLA-Permissive-1.0"
	LICENSE_CDLAPermissive20                  License = "CDLA-Permissive-2.0"
	LICENSE_CDLASharing10                     License = "CDLA-Sharing-1.0"
	LICENSE_CECILL10                          License = "CECILL-1.0"
	LICENSE_CECILL11                          License = "CECILL-1.1"
	LICENSE_CECILL20                          License = "CECILL-2.0"
	LICENSE_CECILL21                          License = "CECILL-2.1"
	LICENSE_CECILLB                           License = "CECILL-B"
	LICENSE_CECILLC                           License = "CECILL-C"
	LICENSE_CERNOHL11                         License = "CERN-OHL-1.1"
	LICENSE_CERNOHL12                         License = "CERN-OHL-1.2"
	LICENSE_CERNOHLP20                        License = "CERN-OHL-P-2.0"
	LICENSE_CERNOHLS20                        License = "CERN-OHL-S-2.0"
	LICENSE_CERNOHLW20                        License = "CERN-OHL-W-2.0"
	LICENSE_ClArtistic                        License = "ClArtistic"
	LICENSE_CNRIJython                        License = "CNRI-Jython"
	LICENSE_CNRIPython                        License = "CNRI-Python"
	LICENSE_CNRIPythonGPLCompatible           License = "CNRI-Python-GPL-Compatible"
	LICENSE_COIL10                            License = "COIL-1.0"
	LICENSE_CommunitySpec10                   License = "Community-Spec-1.0"
	LICENSE_Condor11                          License = "Condor-1.1"
	LICENSE_copyleftnext030                   License = "copyleft-next-0.3.0"
	LICENSE_copyleftnext031                   License = "copyleft-next-0.3.1"
	LICENSE_CPAL10                            License = "CPAL-1.0"
	LICENSE_CPL10                             License = "CPL-1.0"
	LICENSE_CPOL102                           License = "CPOL-1.02"
	LICENSE_Crossword                         License = "Crossword"
	LICENSE_CrystalStacker                    License = "CrystalStacker"
	LICENSE_CUAOPL10                          License = "CUA-OPL-1.0"
	LICENSE_Cube                              License = "Cube"
	LICENSE_curl                              License = "curl"
	LICENSE_DFSL10                            License = "D-FSL-1.0"
	LICENSE_diffmark                          License = "diffmark"
	LICENSE_DLDEBY20                          License = "DL-DE-BY-2.0"
	LICENSE_DOC                               License = "DOC"
	LICENSE_Dotseqn                           License = "Dotseqn"
	LICENSE_DRL10                             License = "DRL-1.0"
	LICENSE_DSDP                              License = "DSDP"
	LICENSE_dvipdfm                           License = "dvipdfm"
	LICENSE_ECL10                             License = "ECL-1.0"
	LICENSE_ECL20                             License = "ECL-2.0"
	LICENSE_eCos20                            License = "eCos-2.0"
	LICENSE_EFL10                             License = "EFL-1.0"
	LICENSE_EFL20                             License = "EFL-2.0"
	LICENSE_eGenix                            License = "eGenix"
	LICENSE_Elastic20                         License = "Elastic-2.0"
	LICENSE_Entessa                           License = "Entessa"
	LICENSE_EPICS                             License = "EPICS"
	LICENSE_EPL10                             License = "EPL-1.0"
	LICENSE_EPL20                             License = "EPL-2.0"
	LICENSE_ErlPL11                           License = "ErlPL-1.1"
	LICENSE_etalab20                          License = "etalab-2.0"
	LICENSE_EUDatagrid                        License = "EUDatagrid"
	LICENSE_EUPL10                            License = "EUPL-1.0"
	LICENSE_EUPL11                            License = "EUPL-1.1"
	LICENSE_EUPL12                            License = "EUPL-1.2"
	LICENSE_Eurosym                           License = "Eurosym"
	LICENSE_Fair                              License = "Fair"
	LICENSE_FDKAAC                            License = "FDK-AAC"
	LICENSE_Frameworx10                       License = "Frameworx-1.0"
	LICENSE_FreeBSDDOC                        License = "FreeBSD-DOC"
	LICENSE_FreeImage                         License = "FreeImage"
	LICENSE_FSFAP                             License = "FSFAP"
	LICENSE_FSFUL                             License = "FSFUL"
	LICENSE_FSFULLR                           License = "FSFULLR"
	LICENSE_FTL                               License = "FTL"
	LICENSE_GD                                License = "GD"
	LICENSE_GFDL11                            License = "GFDL-1.1"
	LICENSE_GFDL11invariantsonly              License = "GFDL-1.1-invariants-only"
	LICENSE_GFDL11invariantsorlater           License = "GFDL-1.1-invariants-or-later"
	LICENSE_GFDL11noinvariantsonly            License = "GFDL-1.1-no-invariants-only"
	LICENSE_GFDL11noinvariantsorlater         License = "GFDL-1.1-no-invariants-or-later"
	LICENSE_GFDL11only                        License = "GFDL-1.1-only"
	LICENSE_GFDL11orlater                     License = "GFDL-1.1-or-later"
	LICENSE_GFDL12invariantsonly              License = "GFDL-1.2-invariants-only"
	LICENSE_GFDL12invariantsorlater           License = "GFDL-1.2-invariants-or-later"
	LICENSE_GFDL12noinvariantsonly            License = "GFDL-1.2-no-invariants-only"
	LICENSE_GFDL12noinvariantsorlater         License = "GFDL-1.2-no-invariants-or-later"
	LICENSE_GFDL12only                        License = "GFDL-1.2-only"
	LICENSE_GFDL12orlater                     License = "GFDL-1.2-or-later"
	LICENSE_GFDL12                            License = "GFDL-1.2"
	LICENSE_GFDL13invariantsonly              License = "GFDL-1.3-invariants-only"
	LICENSE_GFDL13invariantsorlater           License = "GFDL-1.3-invariants-or-later"
	LICENSE_GFDL13noinvariantsonly            License = "GFDL-1.3-no-invariants-only"
	LICENSE_GFDL13noinvariantsorlater         License = "GFDL-1.3-no-invariants-or-later"
	LICENSE_GFDL13only                        License = "GFDL-1.3-only"
	LICENSE_GFDL13orlater                     License = "GFDL-1.3-or-later"
	LICENSE_GFDL13                            License = "GFDL-1.3"
	LICENSE_Giftware                          License = "Giftware"
	LICENSE_GL2PS                             License = "GL2PS"
	LICENSE_Glide                             License = "Glide"
	LICENSE_Glulxe                            License = "Glulxe"
	LICENSE_GLWTPL                            License = "GLWTPL"
	LICENSE_gnuplot                           License = "gnuplot"
	LICENSE_GPL10                             License = "GPL-1.0"
	LICENSE_GPL10only                         License = "GPL-1.0-only"
	LICENSE_GPL10orlater                      License = "GPL-1.0-or-later"
	LICENSE_GPL10PLUS                         License = "GPL-1.0+"
	LICENSE_GPL20                             License = "GPL-2.0"
	LICENSE_GPL20only                         License = "GPL-2.0-only"
	LICENSE_GPL20orlater                      License = "GPL-2.0-or-later"
	LICENSE_GPL20withautoconfexception        License = "GPL-2.0-with-autoconf-exception"
	LICENSE_GPL20withbisonexception           License = "GPL-2.0-with-bison-exception"
	LICENSE_GPL20withclasspathexception       License = "GPL-2.0-with-classpath-exception"
	LICENSE_GPL20withfontexception            License = "GPL-2.0-with-font-exception"
	LICENSE_GPL20withGCCexception             License = "GPL-2.0-with-GCC-exception"
	LICENSE_GPL20PLUS                         License = "GPL-2.0+"
	LICENSE_GPL30                             License = "GPL-3.0"
	LICENSE_GPL30only                         License = "GPL-3.0-only"
	LICENSE_GPL30orlater                      License = "GPL-3.0-or-later"
	LICENSE_GPL30withautoconfexception        License = "GPL-3.0-with-autoconf-exception"
	LICENSE_GPL30withGCCexception             License = "GPL-3.0-with-GCC-exception"
	LICENSE_GPL30PLUS                         License = "GPL-3.0+"
	LICENSE_gSOAP13b                          License = "gSOAP-1.3b"
	LICENSE_HaskellReport                     License = "HaskellReport"
	LICENSE_Hippocratic21                     License = "Hippocratic-2.1"
	LICENSE_HPND                              License = "HPND"
	LICENSE_HPNDsellvariant                   License = "HPND-sell-variant"
	LICENSE_HTMLTIDY                          License = "HTMLTIDY"
	LICENSE_IBMpibs                           License = "IBM-pibs"
	LICENSE_ICU                               License = "ICU"
	LICENSE_IJG                               License = "IJG"
	LICENSE_ImageMagick                       License = "ImageMagick"
	LICENSE_iMatix                            License = "iMatix"
	LICENSE_Imlib2                            License = "Imlib2"
	LICENSE_InfoZIP                           License = "Info-ZIP"
	LICENSE_Intel                             License = "Intel"
	LICENSE_IntelACPI                         License = "Intel-ACPI"
	LICENSE_Interbase10                       License = "Interbase-1.0"
	LICENSE_IPA                               License = "IPA"
	LICENSE_IPL10                             License = "IPL-1.0"
	LICENSE_ISC                               License = "ISC"
	LICENSE_Jam                               License = "Jam"
	LICENSE_JasPer20                          License = "JasPer-2.0"
	LICENSE_JPNIC                             License = "JPNIC"
	LICENSE_JSON                              License = "JSON"
	LICENSE_LAL12                             License = "LAL-1.2"
	LICENSE_LAL13                             License = "LAL-1.3"
	LICENSE_Latex2e                           License = "Latex2e"
	LICENSE_Leptonica                         License = "Leptonica"
	LICENSE_LGPL20                            License = "LGPL-2.0"
	LICENSE_LGPL20PLUS                        License = "LGPL-2.0+"
	LICENSE_LGPL20only                        License = "LGPL-2.0-only"
	LICENSE_LGPL20orlater                     License = "LGPL-2.0-or-later"
	LICENSE_LGPL21                            License = "LGPL-2.1"
	LICENSE_LGPL21only                        License = "LGPL-2.1-only"
	LICENSE_LGPL21PLUS                        License = "LGPL-2.1+"
	LICENSE_LGPL21orlater                     License = "LGPL-2.1-or-later"
	LICENSE_LGPL30                            License = "LGPL-3.0"
	LICENSE_LGPL30only                        License = "LGPL-3.0-only"
	LICENSE_LGPL30orlater                     License = "LGPL-3.0-or-later"
	LICENSE_LGPL30PLUS                        License = "LGPL-3.0+"
	LICENSE_LGPLLR                            License = "LGPLLR"
	LICENSE_Libpng                            License = "Libpng"
	LICENSE_libpng20                          License = "libpng-2.0"
	LICENSE_libselinux10                      License = "libselinux-1.0"
	LICENSE_libtiff                           License = "libtiff"
	LICENSE_LiLiQP11                          License = "LiLiQ-P-1.1"
	LICENSE_LiLiQR11                          License = "LiLiQ-R-1.1"
	LICENSE_LiLiQRplus11                      License = "LiLiQ-Rplus-1.1"
	LICENSE_Linuxmanpagescopyleft             License = "Linux-man-pages-copyleft"
	LICENSE_LinuxOpenIB                       License = "Linux-OpenIB"
	LICENSE_LPL10                             License = "LPL-1.0"
	LICENSE_LPL102                            License = "LPL-1.02"
	LICENSE_LPPL10                            License = "LPPL-1.0"
	LICENSE_LPPL11                            License = "LPPL-1.1"
	LICENSE_LPPL12                            License = "LPPL-1.2"
	LICENSE_LPPL13a                           License = "LPPL-1.3a"
	LICENSE_LPPL13c                           License = "LPPL-1.3c"
	LICENSE_MakeIndex                         License = "MakeIndex"
	LICENSE_MirOS                             License = "MirOS"
	LICENSE_MIT                               License = "MIT"
	LICENSE_MIT0                              License = "MIT-0"
	LICENSE_MITadvertising                    License = "MIT-advertising"
	LICENSE_MITCMU                            License = "MIT-CMU"
	LICENSE_MITenna                           License = "MIT-enna"
	LICENSE_MITfeh                            License = "MIT-feh"
	LICENSE_MITModernVariant                  License = "MIT-Modern-Variant"
	LICENSE_MITopengroup                      License = "MIT-open-group"
	LICENSE_MITNFA                            License = "MITNFA"
	LICENSE_Motosoto                          License = "Motosoto"
	LICENSE_mpich2                            License = "mpich2"
	LICENSE_MPL10                             License = "MPL-1.0"
	LICENSE_MPL11                             License = "MPL-1.1"
	LICENSE_MPL20                             License = "MPL-2.0"
	LICENSE_MPL20nocopyleftexception          License = "MPL-2.0-no-copyleft-exception"
	LICENSE_MSPL                              License = "MS-PL"
	LICENSE_MSRL                              License = "MS-RL"
	LICENSE_MTLL                              License = "MTLL"
	LICENSE_MulanPSL10                        License = "MulanPSL-1.0"
	LICENSE_MulanPSL20                        License = "MulanPSL-2.0"
	LICENSE_Multics                           License = "Multics"
	LICENSE_Mup                               License = "Mup"
	LICENSE_NAIST2003                         License = "NAIST-2003"
	LICENSE_NASA13                            License = "NASA-1.3"
	LICENSE_Naumen                            License = "Naumen"
	LICENSE_NBPL10                            License = "NBPL-1.0"
	LICENSE_NCGLUK20                          License = "NCGL-UK-2.0"
	LICENSE_NCSA                              License = "NCSA"
	LICENSE_NetSNMP                           License = "Net-SNMP"
	LICENSE_NetCDF                            License = "NetCDF"
	LICENSE_Newsletr                          License = "Newsletr"
	LICENSE_NGPL                              License = "NGPL"
	LICENSE_NISTPD                            License = "NIST-PD"
	LICENSE_NISTPDfallback                    License = "NIST-PD-fallback"
	LICENSE_NLOD10                            License = "NLOD-1.0"
	LICENSE_NLOD20                            License = "NLOD-2.0"
	LICENSE_Nunit                             License = "Nunit"
	LICENSE_NLPL                              License = "NLPL"
	LICENSE_Nokia                             License = "Nokia"
	LICENSE_NOSL                              License = "NOSL"
	LICENSE_Noweb                             License = "Noweb"
	LICENSE_NPL10                             License = "NPL-1.0"
	LICENSE_NPL11                             License = "NPL-1.1"
	LICENSE_NPOSL30                           License = "NPOSL-3.0"
	LICENSE_NRL                               License = "NRL"
	LICENSE_NTP                               License = "NTP"
	LICENSE_NTP0                              License = "NTP-0"
	LICENSE_OUDA10                            License = "O-UDA-1.0"
	LICENSE_OCCTPL                            License = "OCCT-PL"
	LICENSE_OCLC20                            License = "OCLC-2.0"
	LICENSE_ODbL10                            License = "ODbL-1.0"
	LICENSE_ODCBy10                           License = "ODC-By-1.0"
	LICENSE_OFL10                             License = "OFL-1.0"
	LICENSE_OFL10noRFN                        License = "OFL-1.0-no-RFN"
	LICENSE_OFL10RFN                          License = "OFL-1.0-RFN"
	LICENSE_OFL11                             License = "OFL-1.1"
	LICENSE_OFL11noRFN                        License = "OFL-1.1-no-RFN"
	LICENSE_OFL11RFN                          License = "OFL-1.1-RFN"
	LICENSE_OGC10                             License = "OGC-1.0"
	LICENSE_OGDLTaiwan10                      License = "OGDL-Taiwan-1.0"
	LICENSE_OGLCanada20                       License = "OGL-Canada-2.0"
	LICENSE_OGLUK10                           License = "OGL-UK-1.0"
	LICENSE_OGLUK20                           License = "OGL-UK-2.0"
	LICENSE_OGLUK30                           License = "OGL-UK-3.0"
	LICENSE_OGTSL                             License = "OGTSL"
	LICENSE_OLDAP11                           License = "OLDAP-1.1"
	LICENSE_OLDAP12                           License = "OLDAP-1.2"
	LICENSE_OLDAP13                           License = "OLDAP-1.3"
	LICENSE_OLDAP14                           License = "OLDAP-1.4"
	LICENSE_OLDAP20                           License = "OLDAP-2.0"
	LICENSE_OLDAP201                          License = "OLDAP-2.0.1"
	LICENSE_OLDAP21                           License = "OLDAP-2.1"
	LICENSE_OLDAP22                           License = "OLDAP-2.2"
	LICENSE_OLDAP221                          License = "OLDAP-2.2.1"
	LICENSE_OLDAP222                          License = "OLDAP-2.2.2"
	LICENSE_OLDAP23                           License = "OLDAP-2.3"
	LICENSE_OLDAP24                           License = "OLDAP-2.4"
	LICENSE_OLDAP25                           License = "OLDAP-2.5"
	LICENSE_OLDAP26                           License = "OLDAP-2.6"
	LICENSE_OLDAP27                           License = "OLDAP-2.7"
	LICENSE_OLDAP28                           License = "OLDAP-2.8"
	LICENSE_OML                               License = "OML"
	LICENSE_OpenSSL                           License = "OpenSSL"
	LICENSE_OPL10                             License = "OPL-1.0"
	LICENSE_OPUBL10                           License = "OPUBL-1.0"
	LICENSE_OSETPL21                          License = "OSET-PL-2.1"
	LICENSE_OSL10                             License = "OSL-1.0"
	LICENSE_OSL11                             License = "OSL-1.1"
	LICENSE_OSL20                             License = "OSL-2.0"
	LICENSE_OSL21                             License = "OSL-2.1"
	LICENSE_OSL30                             License = "OSL-3.0"
	LICENSE_Parity600                         License = "Parity-6.0.0"
	LICENSE_Parity700                         License = "Parity-7.0.0"
	LICENSE_PDDL10                            License = "PDDL-1.0"
	LICENSE_PHP30                             License = "PHP-3.0"
	LICENSE_PHP301                            License = "PHP-3.01"
	LICENSE_Plexus                            License = "Plexus"
	LICENSE_PolyFormNoncommercial100          License = "PolyForm-Noncommercial-1.0.0"
	LICENSE_PolyFormSmallBusiness100          License = "PolyForm-Small-Business-1.0.0"
	LICENSE_PostgreSQL                        License = "PostgreSQL"
	LICENSE_PSF20                             License = "PSF-2.0"
	LICENSE_psfrag                            License = "psfrag"
	LICENSE_psutils                           License = "psutils"
	LICENSE_Python20                          License = "Python-2.0"
	LICENSE_Qhull                             License = "Qhull"
	LICENSE_QPL10                             License = "QPL-1.0"
	LICENSE_Rdisc                             License = "Rdisc"
	LICENSE_RHeCos11                          License = "RHeCos-1.1"
	LICENSE_RPL11                             License = "RPL-1.1"
	LICENSE_RPL15                             License = "RPL-1.5"
	LICENSE_RPSL10                            License = "RPSL-1.0"
	LICENSE_RSAMD                             License = "RSA-MD"
	LICENSE_RSCPL                             License = "RSCPL"
	LICENSE_Ruby                              License = "Ruby"
	LICENSE_SAXPD                             License = "SAX-PD"
	LICENSE_Saxpath                           License = "Saxpath"
	LICENSE_SCEA                              License = "SCEA"
	LICENSE_SchemeReport                      License = "SchemeReport"
	LICENSE_StandardMLNJ                      License = "StandardML-NJ"
	LICENSE_Sendmail                          License = "Sendmail"
	LICENSE_Sendmail823                       License = "Sendmail-8.23"
	LICENSE_SGIB10                            License = "SGI-B-1.0"
	LICENSE_SGIB11                            License = "SGI-B-1.1"
	LICENSE_SGIB20                            License = "SGI-B-2.0"
	LICENSE_SHL05                             License = "SHL-0.5"
	LICENSE_SHL051                            License = "SHL-0.51"
	LICENSE_SimPL20                           License = "SimPL-2.0"
	LICENSE_SISSL                             License = "SISSL"
	LICENSE_SISSL12                           License = "SISSL-1.2"
	LICENSE_Sleepycat                         License = "Sleepycat"
	LICENSE_SMLNJ                             License = "SMLNJ"
	LICENSE_SMPPL                             License = "SMPPL"
	LICENSE_SNIA                              License = "SNIA"
	LICENSE_Spencer86                         License = "Spencer-86"
	LICENSE_Spencer94                         License = "Spencer-94"
	LICENSE_Spencer99                         License = "Spencer-99"
	LICENSE_SPL10                             License = "SPL-1.0"
	LICENSE_SSHOpenSSH                        License = "SSH-OpenSSH"
	LICENSE_SSHshort                          License = "SSH-short"
	LICENSE_SSPL10                            License = "SSPL-1.0"
	LICENSE_SugarCRM113                       License = "SugarCRM-1.1.3"
	LICENSE_SWL                               License = "SWL"
	LICENSE_TAPROHL10                         License = "TAPR-OHL-1.0"
	LICENSE_TCL                               License = "TCL"
	LICENSE_TCPwrappers                       License = "TCP-wrappers"
	LICENSE_TMate                             License = "TMate"
	LICENSE_TORQUE11                          License = "TORQUE-1.1"
	LICENSE_TOSL                              License = "TOSL"
	LICENSE_TUBerlin10                        License = "TU-Berlin-1.0"
	LICENSE_TUBerlin20                        License = "TU-Berlin-2.0"
	LICENSE_UCL10                             License = "UCL-1.0"
	LICENSE_UnicodeDFS2015                    License = "Unicode-DFS-2015"
	LICENSE_UnicodeDFS2016                    License = "Unicode-DFS-2016"
	LICENSE_UnicodeTOU                        License = "Unicode-TOU"
	LICENSE_Unlicense                         License = "Unlicense"
	LICENSE_UPL10                             License = "UPL-1.0"
	LICENSE_Vim                               License = "Vim"
	LICENSE_VOSTROM                           License = "VOSTROM"
	LICENSE_VSL10                             License = "VSL-1.0"
	LICENSE_W3C                               License = "W3C"
	LICENSE_W3C19980720                       License = "W3C-19980720"
	LICENSE_W3C20150513                       License = "W3C-20150513"
	LICENSE_Watcom10                          License = "Watcom-1.0"
	LICENSE_Wsuipa                            License = "Wsuipa"
	LICENSE_WTFPL                             License = "WTFPL"
	LICENSE_X11                               License = "X11"
	LICENSE_X11distributemodificationsvariant License = "X11-distribute-modifications-variant"
	LICENSE_wxWindows                         License = "wx-Windows"
	LICENSE_Xerox                             License = "Xerox"
	LICENSE_XFree8611                         License = "XFree86-1.1"
	LICENSE_xinetd                            License = "xinetd"
	LICENSE_Xnet                              License = "Xnet"
	LICENSE_xpp                               License = "xpp"
	LICENSE_XSkat                             License = "XSkat"
	LICENSE_YPL10                             License = "YPL-1.0"
	LICENSE_YPL11                             License = "YPL-1.1"
	LICENSE_Zed                               License = "Zed"
	LICENSE_Zend20                            License = "Zend-2.0"
	LICENSE_Zimbra13                          License = "Zimbra-1.3"
	LICENSE_Zimbra14                          License = "Zimbra-1.4"
	LICENSE_Zlib                              License = "Zlib"
	LICENSE_zlibacknowledgement               License = "zlib-acknowledgement"
	LICENSE_ZPL11                             License = "ZPL-1.1"
	LICENSE_ZPL20                             License = "ZPL-2.0"
	LICENSE_ZPL21                             License = "ZPL-2.1"
)

//LicenseContainer is used to store the licenses as an array
type LicenseContainer struct {
	Data []License
}

//Language represents a languages identifier
type Language string

//ORCID represents a ORCID identifier
type ORCID string

//ISBN represents a ISBN
type ISBN string

//ISSN represents a ISSN
type ISSN string

//PMCID represents a PMC ID
type PMCID string

//SWHID represents a SWH ID
type SWHID string

//TYPE represents the type of the reference
type TYPE string

const (
	TYPE_art                    TYPE = "art"
	TYPE_article                TYPE = "article"
	TYPE_audiovisual            TYPE = "audiovisual"
	TYPE_bill                   TYPE = "bill"
	TYPE_blog                   TYPE = "blog"
	TYPE_book                   TYPE = "book"
	TYPE_catalogue              TYPE = "catalogue"
	TYPE_conferencePaper        TYPE = "conference-paper"
	TYPE_conference             TYPE = "conference"
	TYPE_data                   TYPE = "data"
	TYPE_database               TYPE = "database"
	TYPE_dictionary             TYPE = "dictionary"
	TYPE_editedWork             TYPE = "edited-work"
	TYPE_encyclopedia           TYPE = "encyclopedia"
	TYPE_filmBroadcast          TYPE = "film-broadcast"
	TYPE_generic                TYPE = "generic"
	TYPE_governmentDocument     TYPE = "government-document"
	TYPE_grant                  TYPE = "grant"
	TYPE_hearing                TYPE = "hearing"
	TYPE_historicalWork         TYPE = "historical-work"
	TYPE_legalCase              TYPE = "legal-case"
	TYPE_legalRule              TYPE = "legal-rule"
	TYPE_magazineArticle        TYPE = "magazine-article"
	TYPE_manual                 TYPE = "manual"
	TYPE_map                    TYPE = "map"
	TYPE_multimedia             TYPE = "multimedia"
	TYPE_music                  TYPE = "music"
	TYPE_newspaperArticle       TYPE = "newspaper-article"
	TYPE_pamphlet               TYPE = "pamphlet"
	TYPE_patent                 TYPE = "patent"
	TYPE_personalCommunication  TYPE = "personal-communication"
	TYPE_proceedings            TYPE = "proceedings"
	TYPE_report                 TYPE = "report"
	TYPE_serial                 TYPE = "serial"
	TYPE_slides                 TYPE = "slides"
	TYPE_softwareCode           TYPE = "software-code"
	TYPE_softwareContainer      TYPE = "software-container"
	TYPE_softwareExecutable     TYPE = "software-executable"
	TYPE_softwareVirtualMachine TYPE = "software-virtual-machine"
	TYPE_software               TYPE = "software"
	TYPE_soundRecording         TYPE = "sound-recording"
	TYPE_standard               TYPE = "standard"
	TYPE_statute                TYPE = "statute"
	TYPE_thesis                 TYPE = "thesis"
	TYPE_unpublished            TYPE = "unpublished"
	TYPE_video                  TYPE = "video"
	TYPE_website                TYPE = "website"
)

//CFFTYPE represents the type of the work
type CFFTYPE string

const (
	CFFTYPE_Software CFFTYPE = "software"
	CFFTYPE_Dataset  CFFTYPE = "dataset"
)

//STATUS represents the status of the work
type STATUS string

const (
	STATUS_abstract      STATUS = "abstract"
	STATUS_advanceOnline STATUS = "advance-online"
	STATUS_inPreparation STATUS = "in-preparation"
	STATUS_inPress       STATUS = "in-press"
	STATUS_preprint      STATUS = "preprint"
	STATUS_submitted     STATUS = "submitted"
)

//ReferenceType represents the type of a reference
type ReferenceType string

const (
	ReferenceType_art                    ReferenceType = "art"
	ReferenceType_article                ReferenceType = "article"
	ReferenceType_audiovisual            ReferenceType = "audiovisual"
	ReferenceType_bill                   ReferenceType = "bill"
	ReferenceType_blog                   ReferenceType = "blog"
	ReferenceType_book                   ReferenceType = "book"
	ReferenceType_catalogue              ReferenceType = "catalogue"
	ReferenceType_conferencePaper        ReferenceType = "conference-paper"
	ReferenceType_conference             ReferenceType = "conference"
	ReferenceType_data                   ReferenceType = "data"
	ReferenceType_database               ReferenceType = "database"
	ReferenceType_dictionary             ReferenceType = "dictionary"
	ReferenceType_editedWork             ReferenceType = "edited-work"
	ReferenceType_encyclopedia           ReferenceType = "encyclopedia"
	ReferenceType_filmBroadcast          ReferenceType = "film-broadcast"
	ReferenceType_generic                ReferenceType = "generic"
	ReferenceType_governmentDocument     ReferenceType = "government-document"
	ReferenceType_grant                  ReferenceType = "grant"
	ReferenceType_hearing                ReferenceType = "hearing"
	ReferenceType_historicalWork         ReferenceType = "historical-work"
	ReferenceType_legalCase              ReferenceType = "legal-case"
	ReferenceType_legalRule              ReferenceType = "legal-rule"
	ReferenceType_magazineArticle        ReferenceType = "magazine-article"
	ReferenceType_manual                 ReferenceType = "manual"
	ReferenceType_map                    ReferenceType = "map"
	ReferenceType_multimedia             ReferenceType = "multimedia"
	ReferenceType_music                  ReferenceType = "music"
	ReferenceType_newspaperArticle       ReferenceType = "newspaper-article"
	ReferenceType_pamphlet               ReferenceType = "pamphlet"
	ReferenceType_patent                 ReferenceType = "patent"
	ReferenceType_personalCommunication  ReferenceType = "personal-communication"
	ReferenceType_proceedings            ReferenceType = "proceedings"
	ReferenceType_report                 ReferenceType = "report"
	ReferenceType_serial                 ReferenceType = "serial"
	ReferenceType_slides                 ReferenceType = "slides"
	ReferenceType_softwareCode           ReferenceType = "software-code"
	ReferenceType_softwareContainer      ReferenceType = "software-container"
	ReferenceType_softwareExecutable     ReferenceType = "software-executable"
	ReferenceType_softwareVirtualMachine ReferenceType = "software-virtual-machine"
	ReferenceType_software               ReferenceType = "software"
	ReferenceType_soundRecording         ReferenceType = "sound-recording"
	ReferenceType_standard               ReferenceType = "standard"
	ReferenceType_statute                ReferenceType = "statute"
	ReferenceType_thesis                 ReferenceType = "thesis"
	ReferenceType_unpublished            ReferenceType = "unpublished"
	ReferenceType_video                  ReferenceType = "video"
	ReferenceType_website                ReferenceType = "website"
)

//DOI represents a CFF DOI. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsdoi for more information.
type DOI struct {
	General            string
	DirectoryIndicator string
	RegistrantCode     string
}

//Entity represents an CFF entity. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsentity for more information.
type Entity struct {
	Address   string  `yaml:"address,omitempty"`
	Alias     string  `yaml:"alias,omitempty"`
	City      string  `yaml:"city,omitempty"`
	Country   Country `yaml:"country,omitempty"`
	DateEnd   Date    `yaml:"date-end,omitempty"`
	DateStart Date    `yaml:"date-start,omitempty"`
	Email     string  `yaml:"email,omitempty"`
	Fax       string  `yaml:"fax,omitempty"`
	Location  string  `yaml:"location,omitempty"`
	Name      string  `yaml:"name,omitempty"`
	Orcid     ORCID   `yaml:"orcid,omitempty"`
	PostCode  string  `yaml:"post-code,omitempty"`
	Region    string  `yaml:"region,omitempty"`
	Tel       string  `yaml:"tel,omitempty"`
	Website   URL     `yaml:"website,omitempty"`
}

//Person represents an CFF person. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsperson for more information.
type Person struct {
	Address      string  `yaml:"address,omitempty"`
	Affiliation  string  `yaml:"affiliation,omitempty"`
	Alias        string  `yaml:"alias,omitempty"`
	City         string  `yaml:"city,omitempty"`
	Country      Country `yaml:"country,omitempty"`
	Email        string  `yaml:"email,omitempty"`
	Family       string  `yaml:"family-names,omitempty"`
	Fax          string  `yaml:"fax,omitempty"`
	GivenNames   string  `yaml:"given-names,omitempty"`
	NameParticle string  `yaml:"name-particle,omitempty"`
	NameSuffix   string  `yaml:"name-suffix,omitempty"`
	Orcid        ORCID   `yaml:"orcid,omitempty"`
	PostCode     string  `yaml:"post-code,omitempty"`
	Region       string  `yaml:"region,omitempty"`
	Tel          string  `yaml:"tel,omitempty"`
	Website      URL     `yaml:"website,omitempty"`
}

//PersonEntity a proxy struct which holds either a Person or an Entity.
//Use the bool field to determine which one is set.
type PersonEntity struct {
	Person   Person
	Entity   Entity
	IsEntity bool
	IsPerson bool
}
type IdentifierTest struct {
	IdentifiersType string `yaml:"type,omitempty"`
}
type IdentifierDOI struct {
	IdentifiersType string `yaml:"type,omitempty"`
	Value           DOI    `yaml:"value,omitempty"`
	Description     string `yaml:"description,omitempty"`
}
type IdentifierURL struct {
	IdentifiersType string `yaml:"type,omitempty"`
	Value           URL    `yaml:"value,omitempty"`
	Description     string `yaml:"description,omitempty"`
}
type IdentifierSWH struct {
	IdentifiersType string `yaml:"type,omitempty"`
	Value           SWHID  `yaml:"value,omitempty"`
	Description     string `yaml:"description,omitempty"`
}
type IdentifierOther struct {
	IdentifiersType string `yaml:"type,omitempty"`
	Value           string `yaml:"value,omitempty"`
	Description     string `yaml:"description,omitempty"`
}

//Identifier represents an CFF identifier. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsidentifier for more information.
type Identifier struct {
	DOI     IdentifierDOI
	URL     IdentifierURL
	SWH     IdentifierSWH
	Other   IdentifierOther
	IsDOI   bool
	IsURL   bool
	IsSWH   bool
	IsOther bool
}

//Reference represents an CFF references. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsreference for more information.
type Reference struct {
	Abbreviation       string           `yaml:"abbreviation,omitempty"`
	Abstract           string           `yaml:"abstract,omitempty"`
	Authors            []PersonEntity   `yaml:"authors,omitempty"`
	CollectionDoi      DOI              `yaml:"collection-doi,omitempty"`
	CollectionTitle    string           `yaml:"collection-title,omitempty"`
	CollectionType     string           `yaml:"collection-type,omitempty"`
	Commit             string           `yaml:"commit,omitempty"`
	Conference         Entity           `yaml:"conference,omitempty"`
	Contact            []PersonEntity   `yaml:"contact,omitempty"`
	Copyright          string           `yaml:"copyright,omitempty"`
	DataType           string           `yaml:"data-type,omitempty"`
	DatabaseProvider   Entity           `yaml:"database-provider,omitempty"`
	Database           string           `yaml:"database,omitempty"`
	DateAccessed       Date             `yaml:"date-accessed,omitempty"`
	DateDownloaded     Date             `yaml:"date-downloaded,omitempty"`
	DatePublished      Date             `yaml:"date-published,omitempty"`
	DateReleased       Date             `yaml:"date-released,omitempty"`
	Department         string           `yaml:"department,omitempty"`
	Doi                DOI              `yaml:"doi,omitempty"`
	Edition            string           `yaml:"edition,omitempty"`
	Editors            []PersonEntity   `yaml:"editors,omitempty"`
	EditorsSeries      []PersonEntity   `yaml:"editors-series,omitempty"`
	End                string           `yaml:"end,omitempty"`
	Entry              string           `yaml:"entry,omitempty"`
	Filename           string           `yaml:"filename,omitempty"`
	Format             string           `yaml:"format,omitempty"`
	Identifiers        []Identifier     `yaml:"identifiers,omitempty"`
	Institution        Entity           `yaml:"institution,omitempty"`
	Isbn               ISBN             `yaml:"isbn,omitempty"`
	Issn               ISSN             `yaml:"issn,omitempty"`
	Issue              string           `yaml:"issue,omitempty"`
	IssueDate          string           `yaml:"issue-date,omitempty"`
	IssueTitle         string           `yaml:"issue-title,omitempty"`
	Journal            string           `yaml:"journal,omitempty"`
	Keywords           []string         `yaml:"keywords,omitempty"`
	Languages          []Language       `yaml:"languages,omitempty"`
	License            LicenseContainer `yaml:"license,omitempty"`
	LicenseUrl         URL              `yaml:"license-url,omitempty"`
	LocEnd             string           `yaml:"loc-end,omitempty"`
	LocStart           string           `yaml:"loc-start,omitempty"`
	Location           Entity           `yaml:"location,omitempty"`
	Medium             string           `yaml:"medium,omitempty"`
	Month              int              `yaml:"month,omitempty"`
	Nihmsid            string           `yaml:"nihmsid,omitempty"`
	Notes              string           `yaml:"notes,omitempty"`
	Number             string           `yaml:"number,omitempty"`
	NumberVolumes      string           `yaml:"number-volumes,omitempty"`
	Pages              string           `yaml:"pages,omitempty"`
	PatentStates       []string         `yaml:"patent-states,omitempty"`
	Pmcid              PMCID            `yaml:"pmcid,omitempty"`
	Publisher          Entity           `yaml:"publisher,omitempty"`
	Recipients         []PersonEntity   `yaml:"recipients,omitempty"`
	Repository         URL              `yaml:"repository,omitempty"`
	RepositoryArtifact URL              `yaml:"repository-artifact,omitempty"`
	RepositoryCode     URL              `yaml:"repository-code,omitempty"`
	Scope              string           `yaml:"scope,omitempty"`
	Section            string           `yaml:"section,omitempty"`
	Senders            []PersonEntity   `yaml:"senders,omitempty"`
	Start              string           `yaml:"start,omitempty"`
	Status             STATUS           `yaml:"status,omitempty"`
	Term               string           `yaml:"term,omitempty"`
	ThesisType         string           `yaml:"thesis-type,omitempty"`
	Title              string           `yaml:"title,omitempty"`
	Translators        []PersonEntity   `yaml:"translators,omitempty"`
	ReferenceType      ReferenceType    `yaml:"type,omitempty"`
	Url                URL              `yaml:"url,omitempty"`
	Version            string           `yaml:"version,omitempty"`
	Volume             int              `yaml:"volume,omitempty"`
	VolumeTitle        string           `yaml:"volume-title,omitempty"`
	Year               int              `yaml:"year,omitempty"`
	YearOriginal       int              `yaml:"year-original,omitempty"`
}

//Cff is the representation of a CFF file
type Cff struct {
	Abstract           string           `yaml:"abstract,omitempty"`
	Authors            []PersonEntity   `yaml:"authors,omitempty"`
	CffVersion         string           `yaml:"cff-version,omitempty"`
	Commit             string           `yaml:"commit,omitempty"`
	Contact            []PersonEntity   `yaml:"contact,omitempty"`
	DateReleased       Date             `yaml:"date-released,omitempty"`
	Doi                DOI              `yaml:"doi,omitempty"`
	Identifiers        []Identifier     `yaml:"identifiers,omitempty"`
	Keywords           []string         `yaml:"keywords,omitempty"`
	License            LicenseContainer `yaml:"license,omitempty"`
	LicenseUtl         URL              `yaml:"license-url,omitempty"`
	Message            string           `yaml:"message,omitempty"`
	PreferredCitation  Reference        `yaml:"preferred-citation,omitempty"`
	References         []Reference      `yaml:"references,omitempty"`
	Repository         URL              `yaml:"repository,omitempty"`
	RepositoryArtifact URL              `yaml:"repository-artifact,omitempty"`
	RepositoryCode     URL              `yaml:"repository-code,omitempty"`
	Title              string           `yaml:"title,omitempty"`
	CffType            CFFTYPE          `yaml:"type,omitempty"`
	Url                URL              `yaml:"url,omitempty"`
	Version            string           `yaml:"version,omitempty"`
}
