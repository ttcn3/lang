package documents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type Deliverable struct {
	DocID       string    `json:"doc_id"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"published_at"`
	WorkItemID  int       `json:"wki_id"`
	Files       []string  `json:"files,omitempty"`
}

type result struct {
	RowNum          string `json:"RowNum"`           // Example: "46"
	TotalCount      string `json:"total_count"`      // Example: "920"
	WkiID           string `json:"wki_id"`           // Example: "14072"
	Title           string `json:"TITLE"`            // Example: "Methods for Testing and Specification (MTS); The Tree and Tabular Combined Notation version 3; Part 1: TTCN-3 Core Language"
	WkiReference    string `json:"WKI_REFERENCE"`    // Example: "RES/MTS-00063-1r1"
	EdsPathname     string `json:"EDSpathname"`      // Example: "etsi_es/201800_201899/20187301/01.01.02_60/"
	EdsPDFFilename  string `json:"EDSPDFfilename"`   // Example: "es_20187301v010102p.pdf"
	EdsARCFilename  string `json:"EDSARCfilename"`   // Example: "es_20187301v010102p0.zip"
	EtsiDeliverable string `json:"ETSI_DELIVERABLE"` // Example: "ETSI ES 201 873-1 V1.1.2 (2001-06)"
	StatusCode      string `json:"STATUS_CODE"`      // Example: "12"
	ActionType      string `json:"ACTION_TYPE"`      // Example: "PU"
	IsCurrent       string `json:"IsCurrent"`        // Example: "0"
	Superseded      string `json:"superseded"`       // Example: "0"
	ReviewDate      string `json:"ReviewDate"`       // Example: null
	NewVersions     string `json:"new_versions"`     // Example: "68552,"
	Scope           string `json:"Scope"`            // Example: "TTCN-3 textual description, BNF syntax and operational semantics.\r\n\r\n"
	TB              string `json:"TB"`               // Example: "Methods for Testing & Specification"
	Keywords        string `json:"Keywords"`         // Example: "ASN.1,METHODOLOGY,MTS,TESTING,TTCN"
}

func (r *result) ShortTitle() string {
	var shortTitles = map[string]string{
		"EG 202 103":    "Guide for the use of the second edition of TTCN",
		"ES 201 873-1":  "Core Language",
		"ES 201 873-3":  "Graphical presentation Format (GFT)",
		"ES 201 873-4":  "Operational Semantics",
		"ES 201 873-5":  "Runtime Interface (TRI)",
		"ES 201 873-6":  "Control Interface (TCI)",
		"ES 201 873-10": "Documentation Comment Specification",
		"ES 201 873-7":  "Using ASN.1",
		"ES 201 873-8":  "Using IDL",
		"ES 201 873-9":  "Using XML schema",
		"ES 201 873-11": "Using JSON",

		"ES 202 781": "Configuration and Deployment Support",
		"ES 202 782": "Performance and Real Time Testing",
		"ES 202 784": "Advanced Parameterization",
		"ES 202 785": "Behaviour Types",
		"ES 202 786": "Support of interfaces with continuous signals",
		"ES 202 789": "Extended TRI",
		"ES 203 022": "Advanced Matching",
		"ES 203 790": "Object-Oriented Features",

		"TS 102 950-1": "Conformance Tests: Core Language: Implementation Conformance Statement (ICS)",
		"TS 102 950-2": "Conformance Tests: Core Language: Test Suite Structure and Test Purposes (TSS&TP)",
		"TS 102 950-3": "Conformance Tests: Core Language: Abstract Test Suite (ATS) and Implementation eXtra Information for Testing (IXIT)",
		"TS 103 253":   "Conformance Tests: XML and JSON schema; Implementation Conformance Statement (ICS)",
		"TS 103 254":   "Conformance Tests: XML and JSON schema; Test Suite Structure and Test Purposes (TSS&TP)",
		"TS 103 255":   "Conformance Tests: XML and JSON schema; Abstract Test Suite & IXIT",
		"TS 103 663-1": "Conformance Tests: Object-Oriented Features;  Implementation Conformance Statement (ICS)",
		"TS 103 663-2": "Conformance Tests: Object-Oriented Features;  Test Suite Structure and Test Purposes (TSS&TP)",
		"TS 103 663-3": "Conformance Tests: Object-Oriented Features;  Abstract Test Suite (ATS) and Implementation eXtra Information for Testing (IXIT)",

		"ES 203 119-6": "TDL: Mapping to TTCN-3",
		"ETR 141 ed.1": "TTCN style guide",
		"TR 101 101":   "TTCN interim version including ASN.1 1994 support",
		"TR 101 114":   "Analysis of the use of ASN.1 94 with TTCN and SDL in ETSI deliverables",
		"TR 101 666":   "TTCN (Ed. 2++)",
		"TR 101 680":   "A harmonized integration of ASN.1, TTCN and SDL",
		"TR 101 873-3": "Graphical presentation Format (GFT)",
		"TR 101 874":   "TTCN-2 to TTCN-3 Mapping",
		"TR 102 043":   "Concepts and definition of the TRI",
		"TR 102 560":   "URI name space for TTCN-3",
		"TR 102 976":   "Mobile Reference tests for TTCN-3 tools",
		"TS 101 875":   "Library of Additional Predefined Functions",
		"TS 102 219":   "IDL to TTCN-3 Mapping",
		"TS 102 351":   "IPv6 Test Specification Toolkit",
		"TS 102 995":   "Proforma for TTCN-3 reference test suite",
	}

	if t, ok := shortTitles[r.DocID()]; ok {
		return t
	}
	t := r.Title
	t = strings.ReplaceAll(t, "Methods for Testing and Specification (MTS);", "")
	t = strings.ReplaceAll(t, "Tree and Tabular Combined Notation version 3", "TTCN-3")
	t = strings.ReplaceAll(t, "Testing and Test Control Notation version 3", "TTCN-3")
	t = strings.ReplaceAll(t, "The ", "")
	t = strings.ReplaceAll(t, "TTCN-3;", "")
	t = strings.ReplaceAll(t, "Part 1:", "")
	t = strings.ReplaceAll(t, "Part 2:", "")
	t = strings.ReplaceAll(t, "Part 3:", "")
	t = strings.ReplaceAll(t, "Part 4:", "")
	t = strings.ReplaceAll(t, "Part 5:", "")
	t = strings.ReplaceAll(t, "Part 6:", "")
	t = strings.ReplaceAll(t, "Part 7:", "")
	t = strings.ReplaceAll(t, "Part 8:", "")
	t = strings.ReplaceAll(t, "Part 9:", "")
	t = strings.ReplaceAll(t, "Part 10:", "")
	t = strings.ReplaceAll(t, "Part 11:", "")
	t = strings.ReplaceAll(t, "Part 12:", "")
	t = strings.ReplaceAll(t, "Part 13:", "")
	t = strings.Replace(t, "TTCN-3", "", 1)
	t = strings.TrimSpace(t)
	return t
}

func (r *result) BaseURL() string {
	return fmt.Sprintf("http://www.etsi.org/deliver/%s", r.EdsPathname)
}

func (r *result) PDFFilename() string {
	return r.EdsPDFFilename
}

func (r *result) ArchiveFilename() string {
	return r.EdsARCFilename
}

func (r *result) WorkItemID() int {
	i, _ := strconv.Atoi(r.WkiID)
	return i
}

func (r *result) DocID() string {
	id := strings.Split(r.EtsiDeliverable, " ")
	return strings.Join(id[1:4], " ")
}

func (r *result) PublishDate() time.Time {
	id := strings.Split(r.EtsiDeliverable, " ")
	t, _ := time.Parse("(2006-01)", id[len(id)-1])
	return t
}

// Deliverables returns a list of all ETSI MTS deliverables related to TTCN-3.
func Deliverables() ([]*Deliverable, error) {
	params := []string{
		"includeScope=1",
		"option=com_standardssearch",
		"view=data",
		"format=json",
		"search=TTCN-3",
		"etsiNumber=0",
		"content=0",
		"version=0",
		"historical=0",
		"superseded=0",
		"startDate=1988-01-15",
		"harmonized=0",
		"TB=97",
	}

	deliverables := []*Deliverable{}

	nPages := 1
	for page := 1; page <= nPages; page++ {
		resp, err := http.Get(fmt.Sprintf("https://www.etsi.org/?page=%d&%s", page, strings.Join(params, "&")))
		if err != nil {
			return nil, fmt.Errorf("failed to get ETSI search results: %w", err)
		}
		defer resp.Body.Close()

		partial := []*result{}
		if err := json.NewDecoder(resp.Body).Decode(&partial); err != nil {
			return nil, fmt.Errorf("failed to decode ETSI search results: %w", err)
		}
		for _, r := range partial {
			if total, err := strconv.Atoi(r.TotalCount); err == nil && total > len(partial) {
				nPages = (total + len(partial) - 1) / len(partial)
			}

			d := Deliverable{
				DocID:       r.DocID(),
				PublishedAt: r.PublishDate(),
				WorkItemID:  r.WorkItemID(),
				Title:       r.ShortTitle(),
			}
			if r.EdsPDFFilename != "" {
				d.Files = append(d.Files, path.Join("http://www.etsi.org/deliver/", r.EdsPathname, r.EdsPDFFilename))
			}
			if r.EdsARCFilename != "" {
				d.Files = append(d.Files, path.Join("http://www.etsi.org/deliver/", r.EdsPathname, r.EdsARCFilename))
			}

			deliverables = append(deliverables, &d)
		}
	}

	// Remove superseded deliverables
	m := make(map[string]*Deliverable)
	for _, d := range deliverables {
		n, ok := m[d.DocID]
		if !ok {
			m[d.DocID] = d
			continue
		}

		if d.PublishedAt.After(n.PublishedAt) {
			m[d.DocID] = d
		}
	}
	deliverables = make([]*Deliverable, 0, len(m))
	for _, d := range m {
		deliverables = append(deliverables, d)
	}

	return deliverables, nil
}
