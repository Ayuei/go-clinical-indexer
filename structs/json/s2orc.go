package json

type S2ORCMetadata struct {
	Abstract string `json:"abstract"`
	AclID    string `json:"acl_id"`
	ArxivID  string `json:"arxiv_id"`
	Authors  []struct {
		First  string        `json:"first"`
		Last   string        `json:"last"`
		Middle []interface{} `json:"middle"`
		Suffix string        `json:"suffix"`
	} `json:"authors"`
	Doi                  string   `json:"doi"`
	HasInboundCitations  bool     `json:"has_inbound_citations"`
	HasOutboundCitations bool     `json:"has_outbound_citations"`
	HasPdfParse          bool     `json:"has_pdf_parse"`
	HasPdfAbstractText   bool     `json:"has_pdf_parsed_abstract,omitempty"`
	HasPdfBodyText       bool     `json:"has_pdf_parsed_body_text,omitempty"`
	HasPdfBibEntries     bool     `json:"has_pdf_parsed_bib_entries,omitempty"`
	HasPdfRefEntries     bool     `json:"has_pdf_parsed_ref_entries,omitempty"`
	InboundCitations     []string `json:"inbound_citations"`
	Journal              string   `json:"journal"`
	MagFieldOfStudy      []string `json:"mag_field_of_study"`
	MagID                string   `json:"mag_id"`
	OutboundCitations    []string `json:"outbound_citations"`
	PaperID              string   `json:"paper_id"`
	PmcID                string   `json:"pmc_id"`
	PubmedID             string   `json:"pubmed_id"`
	S2URL                string   `json:"s2_url"`
	Title                string   `json:"title"`
	Venue                string   `json:"venue"`
	Year                 int64    `json:"year"`
}

type S2ORC struct {
	PaperID  string `json:"paper_id"`
	PdfHash  string `json:"_pdf_hash"`
	Abstract []struct {
		CiteSpans []struct {
			End   int64  `json:"end"`
			RefID string `json:"ref_id"`
			Start int64  `json:"start"`
			Text  string `json:"text"`
		} `json:"cite_spans"`
		RefSpans []interface{} `json:"ref_spans"`
		Section  string        `json:"section"`
		Text     string        `json:"text"`
	} `json:"abstract"`
	BodyText []struct {
		CiteSpans []struct {
			End   int64  `json:"end"`
			RefID string `json:"ref_id"`
			Start int64  `json:"start"`
			Text  string `json:"text"`
		} `json:"cite_spans"`
		RefSpans []interface{} `json:"ref_spans"`
		Section  string        `json:"section"`
		Text     string        `json:"text"`
	} `json:"body_text"`
	BibEntries map[string]struct {
		Title   string        `json:"title"`
		Authors []interface{} `json:"authors"`
		Year    interface{}   `json:"year"`
		Venue   string        `json:"venue"`
		Link    interface{}   `json:"link"`
	} `json:"bib_entries"`
	RefEntries map[string]struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"ref_entries"`
}

type Bib struct {
	Ref     string
	Title   string
	Authors []interface{}
	Year    interface{}
	Venue   string
	Link    interface{}
}

type Ref struct {
	Ref  string
	Text string
	Type string
}

type S2ORCFlat struct {
	PaperID    string
	PdfHash    string
	Abstract   interface{}
	BodyText   interface{}
	BibEntries []Bib
	RefEntries []Ref
}

func (s S2ORC) Flatten() S2ORCFlat {
	bibEntries := make([]Bib, 0, len(s.BibEntries))
	refEntries := make([]Ref, 0, len(s.RefEntries))

	for ref, value := range s.BibEntries {
		bibEntries = append(bibEntries, Bib{
			ref,
			value.Title,
			value.Authors,
			value.Year,
			value.Venue,
			value.Link,
		})
	}

	for ref, value := range s.RefEntries {
		refEntries = append(refEntries, Ref{
			ref,
			value.Text,
			value.Type,
		})
	}

	return S2ORCFlat{
		PaperID:    s.PaperID,
		PdfHash:    s.PdfHash,
		Abstract:   s.Abstract,
		BodyText:   s.BodyText,
		BibEntries: bibEntries,
		RefEntries: refEntries,
	}
}
