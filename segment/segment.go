package segment

import (
	"github.com/jacob-elektronik/gofact/segment/types"
)

// Segment edi segments
type Segment struct {
	SType int
	Tag   string
	Data  string
}

// Segment.PrintSegment
func (s Segment) PrintSegment() string {
	ret := ""
	switch s.SType {
	case types.AAI:
		ret += "Segmenttype: Accommodation allocation information"
	case types.ADI:
		ret += "Segmenttype: Health care claim adjudication information"
	case types.ADR:
		ret += "Segmenttype: Address"
	case types.ADS:
		ret += "Segmenttype: Address"
	case types.AGR:
		ret += "Segmenttype: Agreement identification"
	case types.AJT:
		ret += "Segmenttype: Adjustment details"
	case types.ALC:
		ret += "Segmenttype: Allowance or charge"
	case types.ALI:
		ret += "Segmenttype: Additional information"
	case types.ALS:
		ret += "Segmenttype: Additional location information"
	case types.APD:
		ret += "Segmenttype: Additional transport details"
	case types.APP:
		ret += "Segmenttype: Applicability"
	case types.APR:
		ret += "Segmenttype: Additional price information"
	case types.ARD:
		ret += "Segmenttype: Monetary amount function"
	case types.ARR:
		ret += "Segmenttype: Array information"
	case types.ASD:
		ret += "Segmenttype: Service details"
	case types.ASI:
		ret += "Segmenttype: Array structure identification"
	case types.ATI:
		ret += "Segmenttype: Tour information"
	case types.ATR:
		ret += "Segmenttype: Attribute"
	case types.ATT:
		ret += "Segmenttype: Attribute"
	case types.AUT:
		ret += "Segmenttype: Authentication result"
	case types.BAS:
		ret += "Segmenttype: Basis"
	case types.BCD:
		ret += "Segmenttype: Benefit and coverage detail"
	case types.BGM:
		ret += "Segmenttype: Beginning of message"
	case types.BII:
		ret += "Segmenttype: Structure identification"
	case types.BLI:
		ret += "Segmenttype: Billable information"
	case types.BUS:
		ret += "Segmenttype: Business function"
	case types.CAV:
		ret += "Segmenttype: Characteristic value"
	case types.CCD:
		ret += "Segmenttype: Credit cover details"
	case types.CCI:
		ret += "Segmenttype: Characteristic/class id"
	case types.CDI:
		ret += "Segmenttype: Physical or logical state"
	case types.CDS:
		ret += "Segmenttype: Code set identification"
	case types.CDV:
		ret += "Segmenttype: Code value definition"
	case types.CED:
		ret += "Segmenttype: Computer environment details"
	case types.CIN:
		ret += "Segmenttype: Clinical information"
	case types.CLA:
		ret += "Segmenttype: Clause identification"
	case types.CLI:
		ret += "Segmenttype: Clinical intervention"
	case types.CLT:
		ret += "Segmenttype: Clear terminate information"
	case types.CMN:
		ret += "Segmenttype: Commission information"
	case types.CMP:
		ret += "Segmenttype: Composite data element identification"
	case types.CNI:
		ret += "Segmenttype: Consignment information"
	case types.CNT:
		ret += "Segmenttype: Control total"
	case types.CNX:
		ret += "Segmenttype: Connection details"
	case types.CNY:
		ret += "Segmenttype: Country information"
	case types.COD:
		ret += "Segmenttype: Component details"
	case types.COM:
		ret += "Segmenttype: Communication contact"
	case types.CON:
		ret += "Segmenttype: Contact information"
	case types.COT:
		ret += "Segmenttype: Contribution details"
	case types.CPI:
		ret += "Segmenttype: Charge payment instructions"
	case types.CPS:
		ret += "Segmenttype: Consignment packing sequence"
	case types.CPT:
		ret += "Segmenttype: Account identification"
	case types.CRI:
		ret += "Segmenttype: Consumer reference information"
	case types.CST:
		ret += "Segmenttype: Customs status of goods"
	case types.CTA:
		ret += "Segmenttype: Contact information"
	case types.CUR:
		ret += "Segmenttype: Currencies"
	case types.CUX:
		ret += "Segmenttype: Currencies"
	case types.DAM:
		ret += "Segmenttype: Damage"
	case types.DAV:
		ret += "Segmenttype: Daily availability"
	case types.DFN:
		ret += "Segmenttype: Definition function"
	case types.DGS:
		ret += "Segmenttype: Dangerous goods"
	case types.DII:
		ret += "Segmenttype: Directory identification"
	case types.DIM:
		ret += "Segmenttype: Dimensions"
	case types.DIS:
		ret += "Segmenttype: Discount information"
	case types.DLI:
		ret += "Segmenttype: Document line identification"
	case types.DLM:
		ret += "Segmenttype: Delivery limitations"
	case types.DMS:
		ret += "Segmenttype: Document/message summary"
	case types.DNT:
		ret += "Segmenttype: Dental information"
	case types.DOC:
		ret += "Segmenttype: Document/message details"
	case types.DRD:
		ret += "Segmenttype: Data representation details"
	case types.DSG:
		ret += "Segmenttype: Dosage administration"
	case types.DSI:
		ret += "Segmenttype: Data set identification"
	case types.DTI:
		ret += "Segmenttype: Date and time information"
	case types.DTM:
		ret += "Segmenttype: Date/time/period"
	case types.EDT:
		ret += "Segmenttype: Editing details"
	case types.EFI:
		ret += "Segmenttype: External file link identification"
	case types.ELM:
		ret += "Segmenttype: Simple data element details"
	case types.ELU:
		ret += "Segmenttype: Data element usage details"
	case types.ELV:
		ret += "Segmenttype: Element value definition"
	case types.EMP:
		ret += "Segmenttype: Employment details"
	case types.EQA:
		ret += "Segmenttype: Attached equipment"
	case types.EQD:
		ret += "Segmenttype: Equipment details"
	case types.EQN:
		ret += "Segmenttype: Number of units"
	case types.ERC:
		ret += "Segmenttype: Application error information"
	case types.ERI:
		ret += "Segmenttype: Application error information"
	case types.ERP:
		ret += "Segmenttype: Error point details"
	case types.EVE:
		ret += "Segmenttype: Event"
	case types.FCA:
		ret += "Segmenttype: Financial charges allocation"
	case types.FII:
		ret += "Segmenttype: Financial institution information"
	case types.FNS:
		ret += "Segmenttype: Footnote set"
	case types.FNT:
		ret += "Segmenttype: Footnote"
	case types.FOR:
		ret += "Segmenttype: Formula"
	case types.FRM:
		ret += "Segmenttype: Follow-up action"
	case types.FRQ:
		ret += "Segmenttype: Frequency"
	case types.FSQ:
		ret += "Segmenttype: Formula sequence"
	case types.FTI:
		ret += "Segmenttype: Frequent traveller information"
	case types.FTX:
		ret += "Segmenttype: Free text"
	case types.GDS:
		ret += "Segmenttype: Nature of cargo"
	case types.GEI:
		ret += "Segmenttype: Processing information"
	case types.GID:
		ret += "Segmenttype: Goods item details"
	case types.GIN:
		ret += "Segmenttype: Goods identity number"
	case types.GIR:
		ret += "Segmenttype: Related identification numbers"
	case types.GOR:
		ret += "Segmenttype: Governmental requirements"
	case types.GPO:
		ret += "Segmenttype: Geographical position"
	case types.GRU:
		ret += "Segmenttype: Segment group usage details"
	case types.HAN:
		ret += "Segmenttype: Handling instructions"
	case types.HDI:
		ret += "Segmenttype: Hardware device information"
	case types.HDR:
		ret += "Segmenttype: Header information"
	case types.HDS:
		ret += "Segmenttype: Health diagnosis service and delivery"
	case types.HYN:
		ret += "Segmenttype: Hierarchy information"
	case types.ICD:
		ret += "Segmenttype: Insurance cover description"
	case types.ICI:
		ret += "Segmenttype: Insurance cover information"
	case types.IDE:
		ret += "Segmenttype: Identity"
	case types.IFD:
		ret += "Segmenttype: Information detail"
	case types.IFT:
		ret += "Segmenttype: Interactive free text"
	case types.IHC:
		ret += "Segmenttype: Person characteristic"
	case types.IMD:
		ret += "Segmenttype: Item description"
	case types.IND:
		ret += "Segmenttype: Index details"
	case types.INP:
		ret += "Segmenttype: Parties and instruction"
	case types.INV:
		ret += "Segmenttype: Inventory management related details"
	case types.IRQ:
		ret += "Segmenttype: Information required"
	case types.ITC:
		ret += "Segmenttype: Institutional claim"
	case types.ITD:
		ret += "Segmenttype: Information type data"
	case types.ITM:
		ret += "Segmenttype: Item number"
	case types.LAN:
		ret += "Segmenttype: Language"
	case types.LIN:
		ret += "Segmenttype: Line item"
	case types.LKP:
		ret += "Segmenttype: Level indication"
	case types.LNG:
		ret += "Segmenttype: Language"
	case types.LOC:
		ret += "Segmenttype: Place/location identification"
	case types.MAP:
		ret += "Segmenttype: Message application product information"
	case types.MEA:
		ret += "Segmenttype: Measurements"
	case types.MEM:
		ret += "Segmenttype: Membership details"
	case types.MES:
		ret += "Segmenttype: Measurements"
	case types.MKS:
		ret += "Segmenttype: Market/sales channel information"
	case types.MOA:
		ret += "Segmenttype: Monetary amount"
	case types.MOV:
		ret += "Segmenttype: Car delivery instruction"
	case types.MSD:
		ret += "Segmenttype: Message action details"
	case types.MSG:
		ret += "Segmenttype: Message type identification"
	case types.MTD:
		ret += "Segmenttype: Maintenance operation details"
	case types.NAA:
		ret += "Segmenttype: Name and address"
	case types.NAD:
		ret += "Segmenttype: Name and address"
	case types.NAT:
		ret += "Segmenttype: Nationality"
	case types.NME:
		ret += "Segmenttype: Name"
	case types.NUN:
		ret += "Segmenttype: Number of units"
	case types.ODI:
		ret += "Segmenttype: Origin and destination details"
	case types.ODS:
		ret += "Segmenttype: Additional product details"
	case types.ORG:
		ret += "Segmenttype: Originator of request details"
	case types.OTI:
		ret += "Segmenttype: Other insurance"
	case types.PAC:
		ret += "Segmenttype: Package"
	case types.PAI:
		ret += "Segmenttype: Payment instructions"
	case types.PAS:
		ret += "Segmenttype: Attendance"
	case types.PCC:
		ret += "Segmenttype: Premium calculation component details"
	case types.PCD:
		ret += "Segmenttype: Percentage details"
	case types.PCI:
		ret += "Segmenttype: Package identification"
	case types.PDI:
		ret += "Segmenttype: Person demographic information"
	case types.PDT:
		ret += "Segmenttype: Product information"
	case types.PER:
		ret += "Segmenttype: Period related details"
	case types.PGI:
		ret += "Segmenttype: Product group information"
	case types.PIA:
		ret += "Segmenttype: Additional product id"
	case types.PLI:
		ret += "Segmenttype: Product location information"
	case types.PMT:
		ret += "Segmenttype: Payment information"
	case types.PNA:
		ret += "Segmenttype: Party identification"
	case types.POC:
		ret += "Segmenttype: Purpose of conveyance call"
	case types.POP:
		ret += "Segmenttype: Period of operation"
	case types.POR:
		ret += "Segmenttype: Location and/or related time information"
	case types.POS:
		ret += "Segmenttype: Point of sale information"
	case types.PRC:
		ret += "Segmenttype: Process identification"
	case types.PRD:
		ret += "Segmenttype: Product identification"
	case types.PRE:
		ret += "Segmenttype: Price details"
	case types.PRI:
		ret += "Segmenttype: Price details"
	case types.PRO:
		ret += "Segmenttype: Promotions"
	case types.PRT:
		ret += "Segmenttype: Party information"
	case types.PRV:
		ret += "Segmenttype: Proviso details"
	case types.PSD:
		ret += "Segmenttype: Physical sample description"
	case types.PSI:
		ret += "Segmenttype: Service information"
	case types.PTY:
		ret += "Segmenttype: Priority"
	case types.PYT:
		ret += "Segmenttype: Payment terms"
	case types.QRS:
		ret += "Segmenttype: Query and response"
	case types.QTI:
		ret += "Segmenttype: Quantity"
	case types.QTY:
		ret += "Segmenttype: Quantity"
	case types.QUA:
		ret += "Segmenttype: Qualification"
	case types.QVR:
		ret += "Segmenttype: Quantity variances"
	case types.RCI:
		ret += "Segmenttype: Reservation control information"
	case types.RCS:
		ret += "Segmenttype: Requirements and conditions"
	case types.REL:
		ret += "Segmenttype: Relationship"
	case types.RFF:
		ret += "Segmenttype: Reference"
	case types.RFR:
		ret += "Segmenttype: Reference"
	case types.RJL:
		ret += "Segmenttype: Accounting journal identification"
	case types.RLS:
		ret += "Segmenttype: Relationship"
	case types.RNG:
		ret += "Segmenttype: Range details"
	case types.ROD:
		ret += "Segmenttype: Risk object type"
	case types.RPI:
		ret += "Segmenttype: Quantity and action details"
	case types.RSL:
		ret += "Segmenttype: Result"
	case types.RTC:
		ret += "Segmenttype: Rate types"
	case types.RTE:
		ret += "Segmenttype: Rate details"
	case types.RTI:
		ret += "Segmenttype: Rate details"
	case types.RUL:
		ret += "Segmenttype: Rule information"
	case types.SAL:
		ret += "Segmenttype: Remuneration type identification"
	case types.SCC:
		ret += "Segmenttype: Scheduling conditions"
	case types.SCD:
		ret += "Segmenttype: Structure component definition"
	case types.SDT:
		ret += "Segmenttype: Selection details"
	case types.SEG:
		ret += "Segmenttype: Segment identification"
	case types.SEL:
		ret += "Segmenttype: Seal number"
	case types.SEQ:
		ret += "Segmenttype: Sequence details"
	case types.SER:
		ret += "Segmenttype: Facility information"
	case types.SFI:
		ret += "Segmenttype: Safety information"
	case types.SGP:
		ret += "Segmenttype: Split goods placement"
	case types.SGU:
		ret += "Segmenttype: Segment usage details"
	case types.SPR:
		ret += "Segmenttype: Organisation classification details"
	case types.SPS:
		ret += "Segmenttype: Sampling parameters for summary statistics"
	case types.SSR:
		ret += "Segmenttype: Special requirement details"
	case types.STA:
		ret += "Segmenttype: Statistics"
	case types.STC:
		ret += "Segmenttype: Statistical concept"
	case types.STG:
		ret += "Segmenttype: Stages"
	case types.STS:
		ret += "Segmenttype: Status"
	case types.TAX:
		ret += "Segmenttype: Duty/tax/fee details"
	case types.TCC:
		ret += "Segmenttype: Charge/rate calculations"
	case types.TCE:
		ret += "Segmenttype: Time and certainty"
	case types.TDI:
		ret += "Segmenttype: Traveller document information"
	case types.TDT:
		ret += "Segmenttype: Transport information"
	case types.TEM:
		ret += "Segmenttype: Test method"
	case types.TFF:
		ret += "Segmenttype: Tariff information"
	case types.TIF:
		ret += "Segmenttype: Traveller information"
	case types.TIZ:
		ret += "Segmenttype: Time zone information"
	case types.TMD:
		ret += "Segmenttype: Transport movement details"
	case types.TMP:
		ret += "Segmenttype: Temperature"
	case types.TOD:
		ret += "Segmenttype: Terms of delivery or transport"
	case types.TPL:
		ret += "Segmenttype: Transport placement"
	case types.TRF:
		ret += "Segmenttype: Traffic restriction details"
	case types.TRU:
		ret += "Segmenttype: Technical rules"
	case types.TSR:
		ret += "Segmenttype: Transport service requirements"
	case types.TVL:
		ret += "Segmenttype: Travel product information"
	case types.TXS:
		ret += "Segmenttype: Taxes"
	case types.UNA:
		ret += "Segmenttype: Service String Advice"
	case types.UNB:
		ret += "Segmenttype: Interchange Header"
	case types.UNG:
		ret += "Segmenttype: Functional Group Header"
	case types.UNH:
		ret += "Segmenttype: Message Header"
	case types.UNT:
		ret += "Segmenttype: Message Trailer"
	case types.UNE:
		ret += "Segmenttype: Functional Group Trailer"
	case types.UNZ:
		ret += "Segmenttype: Interchange Trailer"
	case types.UNS:
		ret += "Segmenttype: Section control"
	case types.VEH:
		ret += "Segmenttype: Vehicle"
	case types.VLI:
		ret += "Segmenttype: Value list identification"
	}
	ret += " \tTag :" + s.Tag + "\tData: " + string(s.Data)
	return ret
}

// Segment.String
func (s Segment) String() string {
	return s.PrintSegment()
}
