package fakes

import "github.com/cloudfoundry/bosh-bootloader/aws/cloudformation/templates"

type TemplateBuilder struct {
	BuildCall struct {
		Receives struct {
			KeyPairName      string
			NumberOfAZs      int
			LBType           string
			LBCertificateARN string
			IAMUserName      string
			EnvID            string
		}
		Returns struct {
			Template templates.Template
		}
	}
}

func (b *TemplateBuilder) Build(keyPairName string, numberOfAvailabilityZones int, lbType string, lbCertificateARN string, iamUserName string, envID string) templates.Template {
	b.BuildCall.Receives.KeyPairName = keyPairName
	b.BuildCall.Receives.NumberOfAZs = numberOfAvailabilityZones
	b.BuildCall.Receives.LBType = lbType
	b.BuildCall.Receives.LBCertificateARN = lbCertificateARN
	b.BuildCall.Receives.IAMUserName = iamUserName
	b.BuildCall.Receives.EnvID = envID

	return b.BuildCall.Returns.Template
}
