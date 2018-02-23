package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// ReportRentalAgreementPetToText returns a string representation of the chart of accts
func ReportRentalAgreementPetToText(p *rlib.RentalAgreementPet) string {
	end := ""
	if p.DtStop.Year() < rlib.YEARFOREVER {
		end = p.DtStop.Format(rlib.RRDATEINPFMT)
	}
	return fmt.Sprintf("PET%08d  RA%08d  %-25s  %-15s  %-15s  %-15s  %6.2f lb  %-10s  %-10s\n",
		p.PETID, p.RAID, p.Name, p.Type, p.Breed, p.Color, p.Weight, p.DtStart.Format(rlib.RRDATEINPFMT), end)
}

// RRreportRentalAgreementPets generates a report of all rlib.GLAccount accounts
func RRreportRentalAgreementPets(ctx context.Context, ri *ReporterInfo) string {
	m, err := rlib.GetAllRentalAgreementPets(ctx, ri.Raid)
	if err != nil {
		return err.Error()
	}

	s := fmt.Sprintf("%-11s  %-10s  %-25s  %-15s  %-15s  %-15s  %-9s  %-10s  %-10s\n", "PETID", "RAID", "Name", "Type", "Breed", "Color", "Weight", "DtStart", "DtStop")
	for i := 0; i < len(m); i++ {

		// just before printing out, modify end date for this struct if applicable
		rlib.HandleInterfaceEDI(&m[i], ri.Bid)

		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += ReportRentalAgreementPetToText(&m[i])
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportRentalAgreementPets: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}
