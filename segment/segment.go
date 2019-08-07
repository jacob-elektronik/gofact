package segment

import (
	"gofact/segmenttype"
)

// Segment edi segments
type Segment struct {
	SType int
	Tag   string
	Data  string
}

func (s Segment) PrintSegment() string {
	ret := ""
	switch s.SType {
	case segmenttype.ServiceSegment:
		ret += "Segmenttype: ServiceSegment"
	case segmenttype.AAI:
		ret += "Segmenttype: Accommodation allocation information"
	case segmenttype.ADI:
		ret += "Segmenttype: Health care claim adjudication information"
	case segmenttype.ADR:
		ret += "Segmenttype: Address"
	case segmenttype.ADS:
		ret += "Segmenttype: Address"
	case segmenttype.AGR:
		ret += "Segmenttype: Agreement identification"
	case segmenttype.AJT:
		ret += "Segmenttype: Adjustment details"
	case segmenttype.ALC:
		ret += "Segmenttype: Allowance or charge"
	case segmenttype.ALI:
		ret += "Segmenttype: Additional information"
	case segmenttype.ALS:
		ret += "Segmenttype: Additional location information"
	case segmenttype.APD:
		ret += "Segmenttype: Additional transport details"
	case segmenttype.APP:
		ret += "Segmenttype: Applicability"
	case segmenttype.APR:
		ret += "Segmenttype: Additional price information"
	case segmenttype.ARD:
		ret += "Segmenttype: Monetary amount function"
	case segmenttype.ARR:
		ret += "Segmenttype: Array information"
	case segmenttype.ASD:
		ret += "Segmenttype: Service details"
	case segmenttype.ASI:
		ret += "Segmenttype: Array structure identification"
	case segmenttype.ATI:
		ret += "Segmenttype: Tour information"
	case segmenttype.ATR:
		ret += "Segmenttype: Attribute"
	case segmenttype.ATT:
		ret += "Segmenttype: Attribute"
	case segmenttype.AUT:
		ret += "Segmenttype: Authentication result"
	case segmenttype.BAS:
		ret += "Segmenttype: Basis"
	case segmenttype.BCD:
		ret += "Segmenttype: Benefit and coverage detail"
	case segmenttype.BGM:
		ret += "Segmenttype: Beginning of message"
	case segmenttype.BII:
		ret += "Segmenttype: Structure identification"
	case segmenttype.BLI:
		ret += "Segmenttype: Billable information"
	case segmenttype.BUS:
		ret += "Segmenttype: Business function"
	case segmenttype.CAV:
		ret += "Segmenttype: Characteristic value"
	case segmenttype.CCD:
		ret += "Segmenttype: Credit cover details"
	case segmenttype.CCI:
		ret += "Segmenttype: Characteristic/class id"
	case segmenttype.CDI:
		ret += "Segmenttype: Physical or logical state"
	case segmenttype.CDS:
		ret += "Segmenttype: Code set identification"
	case segmenttype.CDV:
		ret += "Segmenttype: Code value definition"
	case segmenttype.CED:
		ret += "Segmenttype: Computer environment details"
	case segmenttype.CIN:
		ret += "Segmenttype: Clinical information"
	case segmenttype.CLA:
		ret += "Segmenttype: Clause identification"
	case segmenttype.CLI:
		ret += "Segmenttype: Clinical intervention"
	case segmenttype.CLT:
		ret += "Segmenttype: Clear terminate information"
	case segmenttype.CMN:
		ret += "Segmenttype: Commission information"
	case segmenttype.CMP:
		ret += "Segmenttype: Composite data element identification"
	case segmenttype.CNI:
		ret += "Segmenttype: Consignment information"
	case segmenttype.CNT:
		ret += "Segmenttype: Control total"
	case segmenttype.CNX:
		ret += "Segmenttype: Connection details"
	case segmenttype.CNY:
		ret += "Segmenttype: Country information"
	case segmenttype.COD:
		ret += "Segmenttype: Component details"
	case segmenttype.COM:
		ret += "Segmenttype: Communication contact"
	case segmenttype.CON:
		ret += "Segmenttype: Contact information"
	case segmenttype.COT:
		ret += "Segmenttype: Contribution details"
	case segmenttype.CPI:
		ret += "Segmenttype: Charge payment instructions"
	case segmenttype.CPS:
		ret += "Segmenttype: Consignment packing sequence"
	case segmenttype.CPT:
		ret += "Segmenttype: Account identification"
	case segmenttype.CRI:
		ret += "Segmenttype: Consumer reference information"
	case segmenttype.CST:
		ret += "Segmenttype: Customs status of goods"
	case segmenttype.CTA:
		ret += "Segmenttype: Contact information"
	case segmenttype.CUR:
		ret += "Segmenttype: Currencies"
	case segmenttype.CUX:
		ret += "Segmenttype: Currencies"
	case segmenttype.DAM:
		ret += "Segmenttype: Damage"
	case segmenttype.DAV:
		ret += "Segmenttype: Daily availability"
	case segmenttype.DFN:
		ret += "Segmenttype: Definition function"
	case segmenttype.DGS:
		ret += "Segmenttype: Dangerous goods"
	case segmenttype.DII:
		ret += "Segmenttype: Directory identification"
	case segmenttype.DIM:
		ret += "Segmenttype: Dimensions"
	case segmenttype.DIS:
		ret += "Segmenttype: Discount information"
	case segmenttype.DLI:
		ret += "Segmenttype: Document line identification"
	case segmenttype.DLM:
		ret += "Segmenttype: Delivery limitations"
	case segmenttype.DMS:
		ret += "Segmenttype: Document/message summary"
	case segmenttype.DNT:
		ret += "Segmenttype: Dental information"
	case segmenttype.DOC:
		ret += "Segmenttype: Document/message details"
	case segmenttype.DRD:
		ret += "Segmenttype: Data representation details"
	case segmenttype.DSG:
		ret += "Segmenttype: Dosage administration"
	case segmenttype.DSI:
		ret += "Segmenttype: Data set identification"
	case segmenttype.DTI:
		ret += "Segmenttype: Date and time information"
	case segmenttype.DTM:
		ret += "Segmenttype: Date/time/period"
	case segmenttype.EDT:
		ret += "Segmenttype: Editing details"
	case segmenttype.EFI:
		ret += "Segmenttype: External file link identification"
	case segmenttype.ELM:
		ret += "Segmenttype: Simple data element details"
	case segmenttype.ELU:
		ret += "Segmenttype: Data element usage details"
	case segmenttype.ELV:
		ret += "Segmenttype: Element value definition"
	case segmenttype.EMP:
		ret += "Segmenttype: Employment details"
	case segmenttype.EQA:
		ret += "Segmenttype: Attached equipment"
	case segmenttype.EQD:
		ret += "Segmenttype: Equipment details"
	case segmenttype.EQN:
		ret += "Segmenttype: Number of units"
	case segmenttype.ERC:
		ret += "Segmenttype: Application error information"
	case segmenttype.ERI:
		ret += "Segmenttype: Application error information"
	case segmenttype.ERP:
		ret += "Segmenttype: Error point details"
	case segmenttype.EVE:
		ret += "Segmenttype: Event"
	case segmenttype.FCA:
		ret += "Segmenttype: Financial charges allocation"
	case segmenttype.FII:
		ret += "Segmenttype: Financial institution information"
	case segmenttype.FNS:
		ret += "Segmenttype: Footnote set"
	case segmenttype.FNT:
		ret += "Segmenttype: Footnote"
	case segmenttype.FOR:
		ret += "Segmenttype: Formula"
	case segmenttype.FRM:
		ret += "Segmenttype: Follow-up action"
	case segmenttype.FRQ:
		ret += "Segmenttype: Frequency"
	case segmenttype.FSQ:
		ret += "Segmenttype: Formula sequence"
	case segmenttype.FTI:
		ret += "Segmenttype: Frequent traveller information"
	case segmenttype.FTX:
		ret += "Segmenttype: Free text"
	case segmenttype.GDS:
		ret += "Segmenttype: Nature of cargo"
	case segmenttype.GEI:
		ret += "Segmenttype: Processing information"
	case segmenttype.GID:
		ret += "Segmenttype: Goods item details"
	case segmenttype.GIN:
		ret += "Segmenttype: Goods identity number"
	case segmenttype.GIR:
		ret += "Segmenttype: Related identification numbers"
	case segmenttype.GOR:
		ret += "Segmenttype: Governmental requirements"
	case segmenttype.GPO:
		ret += "Segmenttype: Geographical position"
	case segmenttype.GRU:
		ret += "Segmenttype: Segment group usage details"
	case segmenttype.HAN:
		ret += "Segmenttype: Handling instructions"
	case segmenttype.HDI:
		ret += "Segmenttype: Hardware device information"
	case segmenttype.HDR:
		ret += "Segmenttype: Header information"
	case segmenttype.HDS:
		ret += "Segmenttype: Health diagnosis service and delivery"
	case segmenttype.HYN:
		ret += "Segmenttype: Hierarchy information"
	case segmenttype.ICD:
		ret += "Segmenttype: Insurance cover description"
	case segmenttype.ICI:
		ret += "Segmenttype: Insurance cover information"
	case segmenttype.IDE:
		ret += "Segmenttype: Identity"
	case segmenttype.IFD:
		ret += "Segmenttype: Information detail"
	case segmenttype.IFT:
		ret += "Segmenttype: Interactive free text"
	case segmenttype.IHC:
		ret += "Segmenttype: Person characteristic"
	case segmenttype.IMD:
		ret += "Segmenttype: Item description"
	case segmenttype.IND:
		ret += "Segmenttype: Index details"
	case segmenttype.INP:
		ret += "Segmenttype: Parties and instruction"
	case segmenttype.INV:
		ret += "Segmenttype: Inventory management related details"
	case segmenttype.IRQ:
		ret += "Segmenttype: Information required"
	case segmenttype.ITC:
		ret += "Segmenttype: Institutional claim"
	case segmenttype.ITD:
		ret += "Segmenttype: Information type data"
	case segmenttype.ITM:
		ret += "Segmenttype: Item number"
	case segmenttype.LAN:
		ret += "Segmenttype: Language"
	case segmenttype.LIN:
		ret += "Segmenttype: Line item"
	case segmenttype.LKP:
		ret += "Segmenttype: Level indication"
	case segmenttype.LNG:
		ret += "Segmenttype: Language"
	case segmenttype.LOC:
		ret += "Segmenttype: Place/location identification"
	case segmenttype.MAP:
		ret += "Segmenttype: Message application product information"
	case segmenttype.MEA:
		ret += "Segmenttype: Measurements"
	case segmenttype.MEM:
		ret += "Segmenttype: Membership details"
	case segmenttype.MES:
		ret += "Segmenttype: Measurements"
	case segmenttype.MKS:
		ret += "Segmenttype: Market/sales channel information"
	case segmenttype.MOA:
		ret += "Segmenttype: Monetary amount"
	case segmenttype.MOV:
		ret += "Segmenttype: Car delivery instruction"
	case segmenttype.MSD:
		ret += "Segmenttype: Message action details"
	case segmenttype.MSG:
		ret += "Segmenttype: Message type identification"
	case segmenttype.MTD:
		ret += "Segmenttype: Maintenance operation details"
	case segmenttype.NAA:
		ret += "Segmenttype: Name and address"
	case segmenttype.NAD:
		ret += "Segmenttype: Name and address"
	case segmenttype.NAT:
		ret += "Segmenttype: Nationality"
	case segmenttype.NME:
		ret += "Segmenttype: Name"
	case segmenttype.NUN:
		ret += "Segmenttype: Number of units"
	case segmenttype.ODI:
		ret += "Segmenttype: Origin and destination details"
	case segmenttype.ODS:
		ret += "Segmenttype: Additional product details"
	case segmenttype.ORG:
		ret += "Segmenttype: Originator of request details"
	case segmenttype.OTI:
		ret += "Segmenttype: Other insurance"
	case segmenttype.PAC:
		ret += "Segmenttype: Package"
	case segmenttype.PAI:
		ret += "Segmenttype: Payment instructions"
	case segmenttype.PAS:
		ret += "Segmenttype: Attendance"
	case segmenttype.PCC:
		ret += "Segmenttype: Premium calculation component details"
	case segmenttype.PCD:
		ret += "Segmenttype: Percentage details"
	case segmenttype.PCI:
		ret += "Segmenttype: Package identification"
	case segmenttype.PDI:
		ret += "Segmenttype: Person demographic information"
	case segmenttype.PDT:
		ret += "Segmenttype: Product information"
	case segmenttype.PER:
		ret += "Segmenttype: Period related details"
	case segmenttype.PGI:
		ret += "Segmenttype: Product group information"
	case segmenttype.PIA:
		ret += "Segmenttype: Additional product id"
	case segmenttype.PLI:
		ret += "Segmenttype: Product location information"
	case segmenttype.PMT:
		ret += "Segmenttype: Payment information"
	case segmenttype.PNA:
		ret += "Segmenttype: Party identification"
	case segmenttype.POC:
		ret += "Segmenttype: Purpose of conveyance call"
	case segmenttype.POP:
		ret += "Segmenttype: Period of operation"
	case segmenttype.POR:
		ret += "Segmenttype: Location and/or related time information"
	case segmenttype.POS:
		ret += "Segmenttype: Point of sale information"
	case segmenttype.PRC:
		ret += "Segmenttype: Process identification"
	case segmenttype.PRD:
		ret += "Segmenttype: Product identification"
	case segmenttype.PRE:
		ret += "Segmenttype: Price details"
	case segmenttype.PRI:
		ret += "Segmenttype: Price details"
	case segmenttype.PRO:
		ret += "Segmenttype: Promotions"
	case segmenttype.PRT:
		ret += "Segmenttype: Party information"
	case segmenttype.PRV:
		ret += "Segmenttype: Proviso details"
	case segmenttype.PSD:
		ret += "Segmenttype: Physical sample description"
	case segmenttype.PSI:
		ret += "Segmenttype: Service information"
	case segmenttype.PTY:
		ret += "Segmenttype: Priority"
	case segmenttype.PYT:
		ret += "Segmenttype: Payment terms"
	case segmenttype.QRS:
		ret += "Segmenttype: Query and response"
	case segmenttype.QTI:
		ret += "Segmenttype: Quantity"
	case segmenttype.QTY:
		ret += "Segmenttype: Quantity"
	case segmenttype.QUA:
		ret += "Segmenttype: Qualification"
	case segmenttype.QVR:
		ret += "Segmenttype: Quantity variances"
	case segmenttype.RCI:
		ret += "Segmenttype: Reservation control information"
	case segmenttype.RCS:
		ret += "Segmenttype: Requirements and conditions"
	case segmenttype.REL:
		ret += "Segmenttype: Relationship"
	case segmenttype.RFF:
		ret += "Segmenttype: Reference"
	case segmenttype.RFR:
		ret += "Segmenttype: Reference"
	case segmenttype.RJL:
		ret += "Segmenttype: Accounting journal identification"
	case segmenttype.RLS:
		ret += "Segmenttype: Relationship"
	case segmenttype.RNG:
		ret += "Segmenttype: Range details"
	case segmenttype.ROD:
		ret += "Segmenttype: Risk object type"
	case segmenttype.RPI:
		ret += "Segmenttype: Quantity and action details"
	case segmenttype.RSL:
		ret += "Segmenttype: Result"
	case segmenttype.RTC:
		ret += "Segmenttype: Rate types"
	case segmenttype.RTE:
		ret += "Segmenttype: Rate details"
	case segmenttype.RTI:
		ret += "Segmenttype: Rate details"
	case segmenttype.RUL:
		ret += "Segmenttype: Rule information"
	case segmenttype.SAL:
		ret += "Segmenttype: Remuneration type identification"
	case segmenttype.SCC:
		ret += "Segmenttype: Scheduling conditions"
	case segmenttype.SCD:
		ret += "Segmenttype: Structure component definition"
	case segmenttype.SDT:
		ret += "Segmenttype: Selection details"
	case segmenttype.SEG:
		ret += "Segmenttype: Segment identification"
	case segmenttype.SEL:
		ret += "Segmenttype: Seal number"
	case segmenttype.SEQ:
		ret += "Segmenttype: Sequence details"
	case segmenttype.SER:
		ret += "Segmenttype: Facility information"
	case segmenttype.SFI:
		ret += "Segmenttype: Safety information"
	case segmenttype.SGP:
		ret += "Segmenttype: Split goods placement"
	case segmenttype.SGU:
		ret += "Segmenttype: Segment usage details"
	case segmenttype.SPR:
		ret += "Segmenttype: Organisation classification details"
	case segmenttype.SPS:
		ret += "Segmenttype: Sampling parameters for summary statistics"
	case segmenttype.SSR:
		ret += "Segmenttype: Special requirement details"
	case segmenttype.STA:
		ret += "Segmenttype: Statistics"
	case segmenttype.STC:
		ret += "Segmenttype: Statistical concept"
	case segmenttype.STG:
		ret += "Segmenttype: Stages"
	case segmenttype.STS:
		ret += "Segmenttype: Status"
	case segmenttype.TAX:
		ret += "Segmenttype: Duty/tax/fee details"
	case segmenttype.TCC:
		ret += "Segmenttype: Charge/rate calculations"
	case segmenttype.TCE:
		ret += "Segmenttype: Time and certainty"
	case segmenttype.TDI:
		ret += "Segmenttype: Traveller document information"
	case segmenttype.TDT:
		ret += "Segmenttype: Transport information"
	case segmenttype.TEM:
		ret += "Segmenttype: Test method"
	case segmenttype.TFF:
		ret += "Segmenttype: Tariff information"
	case segmenttype.TIF:
		ret += "Segmenttype: Traveller information"
	case segmenttype.TIZ:
		ret += "Segmenttype: Time zone information"
	case segmenttype.TMD:
		ret += "Segmenttype: Transport movement details"
	case segmenttype.TMP:
		ret += "Segmenttype: Temperature"
	case segmenttype.TOD:
		ret += "Segmenttype: Terms of delivery or transport"
	case segmenttype.TPL:
		ret += "Segmenttype: Transport placement"
	case segmenttype.TRF:
		ret += "Segmenttype: Traffic restriction details"
	case segmenttype.TRU:
		ret += "Segmenttype: Technical rules"
	case segmenttype.TSR:
		ret += "Segmenttype: Transport service requirements"
	case segmenttype.TVL:
		ret += "Segmenttype: Travel product information"
	case segmenttype.TXS:
		ret += "Segmenttype: Taxes"
	case segmenttype.VEH:
		ret += "Segmenttype: Vehicle"
	case segmenttype.VLI:
		ret += "Segmenttype: Value list identification"
	}
	ret += " \tTag :" + s.Tag + "\tData: " + string(s.Data)
	return ret
}

func (s Segment) String() string {
	return s.PrintSegment()
}
