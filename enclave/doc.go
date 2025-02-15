// Copyright (c) Edgeless Systems GmbH.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

/*
Package enclave provides functionality for Go enclaves like remote attestation and sealing.

Using remote reports

Remote reports are generated by an enclave platform to attest the integrity and
confidentiality of an enclaved app instance. A remote report also attests that an app was
indeed established on a secure enclave platform, and is targeted to a remote third party
which is not running on an (or not on the same) enclave platform.

A remote report can contain additional data held by the enclave, e.g. data that was created
by the enclaved application or data the enclaved app received. This data can be included as
reportData.

GetRemoteReport creates a remote report which includes additional reportData. The following
code can be run by an enclaved app:

	// Create a report that includes the hash of an enclave generated certificate cert.
	hash := sha256.Sum256(cert)
	report, err := enclave.GetRemoteReport(hash)
	if err != nil {
		return err
	}

VerifyRemoteReport can be used by a third party to verify the previous generated remote
report. While VerifyRemoteReport verifies the reports integrity and signature, the third
party must additionally verify the content of the remote report:

	report, err := enclave.VerifyRemoteReport(report)
	if err != nil {
		return err
	}
	if report.SecurityVersion < 2 {
		return errors.New("invalid security version")
	}
	if binary.LittleEndian.Uint16(report.ProductID) != 1234 {
		return errors.New("invalid product")
	}
	if !bytes.Equal(report.SignerID, signer) {
		return errors.New("invalid signer")
	}
	// certBytes and report were sent over insecure channel
	hash := sha256.Sum256(certBytes)
	if !bytes.Equal(report.Data[:len(hash)], hash[:]) {
		return errors.New("report data does not match the certificate's hash")
	}
	// we ensured the cert was generated by the enclave

*/
package enclave
