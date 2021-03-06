package utils

import (
	"github.com/jacob-elektronik/gofact/segment/types"
)

const SubSetDefault = "edifact"
const SubSetEancom = "eancom"

var SegmentTypeFoString = map[string]int{
	"AAI": types.AAI,
	"ADR": types.ADR,
	"ADS": types.ADS,
	"AGR": types.AGR,
	"AJT": types.AJT,
	"ALC": types.ALC,
	"ALI": types.ALI,
	"ALS": types.ALS,
	"APD": types.APD,
	"APP": types.APP,
	"APR": types.APR,
	"ARD": types.ARD,
	"ARR": types.ARR,
	"ASD": types.ASD,
	"ASI": types.ASI,
	"ATI": types.ATI,
	"ATR": types.ATR,
	"ATT": types.ATT,
	"AUT": types.AUT,
	"BAS": types.BAS,
	"BCD": types.BCD,
	"BGM": types.BGM,
	"BII": types.BII,
	"BLI": types.BLI,
	"BUS": types.BUS,
	"CAV": types.CAV,
	"CCD": types.CCD,
	"CCI": types.CCI,
	"CDI": types.CDI,
	"CDS": types.CDS,
	"CDV": types.CDV,
	"CED": types.CED,
	"CIN": types.CIN,
	"CLA": types.CLA,
	"CLI": types.CLI,
	"CLT": types.CLT,
	"CMN": types.CMN,
	"CMP": types.CMP,
	"CNI": types.CNI,
	"CNT": types.CNT,
	"CNX": types.CNX,
	"CNY": types.CNY,
	"COD": types.COD,
	"COM": types.COM,
	"CON": types.CON,
	"COT": types.COT,
	"CPI": types.CPI,
	"CPS": types.CPS,
	"CPT": types.CPT,
	"CRI": types.CRI,
	"CST": types.CST,
	"CTA": types.CTA,
	"CUR": types.CUR,
	"CUX": types.CUX,
	"DAM": types.DAM,
	"DAV": types.DAV,
	"DFN": types.DFN,
	"DGS": types.DGS,
	"DII": types.DII,
	"DIM": types.DIM,
	"DIS": types.DIS,
	"DLI": types.DLI,
	"DLM": types.DLM,
	"DMS": types.DMS,
	"DNT": types.DNT,
	"DOC": types.DOC,
	"DRD": types.DRD,
	"DSG": types.DSG,
	"DSI": types.DSI,
	"DTI": types.DTI,
	"DTM": types.DTM,
	"EDT": types.EDT,
	"EFI": types.EFI,
	"ELM": types.ELM,
	"ELU": types.ELU,
	"ELV": types.ELV,
	"EMP": types.EMP,
	"EQA": types.EQA,
	"EQD": types.EQD,
	"EQN": types.EQN,
	"ERC": types.ERC,
	"ERI": types.ERI,
	"ERP": types.ERP,
	"EVE": types.EVE,
	"FCA": types.FCA,
	"FII": types.FII,
	"FNS": types.FNS,
	"FNT": types.FNT,
	"FOR": types.FOR,
	"FRM": types.FRM,
	"FRQ": types.FRQ,
	"FSQ": types.FSQ,
	"FTI": types.FTI,
	"FTX": types.FTX,
	"GDS": types.GDS,
	"GEI": types.GEI,
	"GID": types.GID,
	"GIN": types.GIN,
	"GIR": types.GIR,
	"GOR": types.GOR,
	"GPO": types.GPO,
	"GRU": types.GRU,
	"HAN": types.HAN,
	"HDI": types.HDI,
	"HDR": types.HDR,
	"HDS": types.HDS,
	"HYN": types.HYN,
	"ICD": types.ICD,
	"ICI": types.ICI,
	"IDE": types.IDE,
	"IFD": types.IFD,
	"IFT": types.IFT,
	"IHC": types.IHC,
	"IMD": types.IMD,
	"IND": types.IND,
	"INP": types.INP,
	"INV": types.INV,
	"IRQ": types.IRQ,
	"ITC": types.ITC,
	"ITD": types.ITD,
	"ITM": types.ITM,
	"LAN": types.LAN,
	"LIN": types.LIN,
	"LKP": types.LKP,
	"LNG": types.LNG,
	"LOC": types.LOC,
	"MAP": types.MAP,
	"MEA": types.MEA,
	"MEM": types.MEM,
	"MES": types.MES,
	"MKS": types.MKS,
	"MOA": types.MOA,
	"MOV": types.MOV,
	"MSD": types.MSD,
	"MSG": types.MSG,
	"MTD": types.MTD,
	"NAA": types.NAA,
	"NAD": types.NAD,
	"NAT": types.NAT,
	"NME": types.NME,
	"NUN": types.NUN,
	"ODI": types.ODI,
	"ODS": types.ODS,
	"ORG": types.ORG,
	"OTI": types.OTI,
	"PAC": types.PAC,
	"PAI": types.PAI,
	"PAS": types.PAS,
	"PCC": types.PCC,
	"PCD": types.PCD,
	"PCI": types.PCI,
	"PDI": types.PDI,
	"PDT": types.PDT,
	"PER": types.PER,
	"PGI": types.PGI,
	"PIA": types.PIA,
	"PLI": types.PLI,
	"PMT": types.PMT,
	"PNA": types.PNA,
	"POC": types.POC,
	"POP": types.POP,
	"POR": types.POR,
	"POS": types.POS,
	"PRC": types.PRC,
	"PRD": types.PRD,
	"PRE": types.PRE,
	"PRI": types.PRI,
	"PRO": types.PRO,
	"PRT": types.PRT,
	"PRV": types.PRV,
	"PSD": types.PSD,
	"PSI": types.PSI,
	"PTY": types.PTY,
	"PYT": types.PYT,
	"QRS": types.QRS,
	"QTI": types.QTI,
	"QTY": types.QTY,
	"QUA": types.QUA,
	"QVR": types.QVR,
	"RCI": types.RCI,
	"RCS": types.RCS,
	"REL": types.REL,
	"RFF": types.RFF,
	"RFR": types.RFR,
	"RJL": types.RJL,
	"RLS": types.RLS,
	"RNG": types.RNG,
	"ROD": types.ROD,
	"RPI": types.RPI,
	"RSL": types.RSL,
	"RTC": types.RTC,
	"RTE": types.RTE,
	"RTI": types.RTI,
	"RUL": types.RUL,
	"SAL": types.SAL,
	"SCC": types.SCC,
	"SCD": types.SCD,
	"SDT": types.SDT,
	"SEG": types.SEG,
	"SEL": types.SEL,
	"SEQ": types.SEQ,
	"SER": types.SER,
	"SFI": types.SFI,
	"SGP": types.SGP,
	"SGU": types.SGU,
	"SPR": types.SPR,
	"SPS": types.SPS,
	"SSR": types.SSR,
	"STA": types.STA,
	"STC": types.STC,
	"STG": types.STG,
	"STS": types.STS,
	"TAX": types.TAX,
	"TCC": types.TCC,
	"TCE": types.TCE,
	"TDI": types.TDI,
	"TDT": types.TDT,
	"TEM": types.TEM,
	"TFF": types.TFF,
	"TIF": types.TIF,
	"TIZ": types.TIZ,
	"TMD": types.TMD,
	"TMP": types.TMP,
	"TOD": types.TOD,
	"TPL": types.TPL,
	"TRF": types.TRF,
	"TRU": types.TRU,
	"TSR": types.TSR,
	"TVL": types.TVL,
	"TXS": types.TXS,
	"UNA": types.UNA,
	"UNB": types.UNB,
	"UNG": types.UNG,
	"UNH": types.UNH,
	"UGH": types.UGH,
	"UGT": types.UGT,
	"UIH": types.UIH,
	"UIT": types.UIT,
	"UNT": types.UNT,
	"UNE": types.UNE,
	"UNS": types.UNS,
	"UNZ": types.UNZ,
	"VEH": types.VEH,
	"VLI": types.VLI,
}
