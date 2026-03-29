---
name: Matt Burdan
left-column:
  - '+1 (415) 542 6132'
  - 'burdz@burdz.net'
  - '@burdzwastaken'
right-column:
  - 'https://www.linkedin.com/in/burdz/'
  - 'https://github.com/burdzwastaken'
  - 'https://void.burdz.net'
  - 'Last Updated: \today'
---

# Summary

Principal Platform Engineer building cloud infrastructure, security tooling and developer platforms across GCP, AWS and on-prem. Currently own the SaaS platform at Ping Identity; multi-region GKE, FedRAMP environments and tenant lifecycle automation. Background in security engineering (PKI, DFIR, vulnerability management) turned platform builder. Implements consensus algorithms in Zig for fun and runs NixOS because apparently I enjoy troubleshooting my troubleshooting tools.

# Skills

**Languages**: Go · Zig · Python · Bash · HCL · Rego · Nix

**Infrastructure**: Kubernetes · Terraform · GCP · AWS · Helm · Docker/OCI · NixOS

**Security**: OPA/Rego · Gatekeeper · Cosign · FedRAMP · osquery · PKI · DFIR

**Observability**: Prometheus · OpenTelemetry · Datadog · PagerDuty · Grafana

**CI/CD**: GitHub Actions · Codefresh · Spinnaker · Concourse

**Systems**: eBPF/XDP · WASM · Linux Kernel Modules · Distributed Systems


# Open Source

**[raftz](https://github.com/burdzwastaken/raftz)**: Raft consensus algorithm implementation in Zig

**[zig-opa-wasm](https://github.com/burdzwastaken/zig-opa-wasm)**: Open Policy Agent WASM runtime in Zig

**[regolint](https://github.com/burdzwastaken/regolint)**: Linter for OPA Rego policies, written in Go


# Experience

## Ping Identity (formerly ForgeRock)
Principal Platform Engineer
Singapore (remote)
2021 - Present

* Own the multi-region SaaS platform infrastructure powering Identity Cloud across US, EU and APAC; thousands of tenants, high availability, on-call for all of it
* Designed and built FedRAMP-compliant infrastructure from scratch, enabling Ping to serve US federal customers with assured workloads and isolated environments
* Built the multi-region tooling for active/passive state management across GKE clusters; leader election, CRD reconciliation and an API that tenant services depend on for failover decisions
* Built the Gatekeeper policy framework: Rego policies that auto-generate ConstraintTemplates, Constraints and documentation; 70+ policies enforcing security and compliance controls across the fleet
* Built self-service tenant lifecycle tooling in Go, replacing manual provisioning processes with automated workflows and reducing environment creation from days to minutes
* Established supply chain security practices across CI/CD: image signing with Cosign, policy-as-code guardrails with OPA/Rego and automated compliance gates
* Designed GCP IAM architecture, workload identity federation and disaster recovery strategies with validated RPO/RTO targets across regions
* Drive platform architecture through design docs and ADRs; mentor engineers on coding, IaC patterns, cloud-native architecture and not paging me at 3am

## MuleSoft (Salesforce)
Senior Platform Engineer
San Francisco, CA
2017 - 2021

* Architected and operated multi-region Kubernetes deployments managing 60,000+ runtimes across US, EU and AWS GovCloud using Spinnaker and Jenkins
* Built Kubernetes CRDs and service catalog extensions for IAM credential injection, infrastructure provisioning and service mesh integration
* Owned the internal OCI registry (Harbor) end-to-end: Helm deployment, Terraform config, ECR replication and vulnerability scanning with Clair
* Developed zero-downtime cluster upgrade automation using multi-cluster DNS-based cutover and automated incident response workflows with PagerDuty and Slack

## Lookout, Inc
Senior Security Engineer
San Francisco, CA
2015 - 2017

* Built and managed the entire PKI infrastructure: offline CAs, OCSP responders, validation authorities and public-facing certificate lifecycle
* Developed osquery table extensions, S3 bootstrapping tools and IAM policy review pipelines in Go for fleet security and configuration distribution
* Architected deployment of security monitoring across the fleet: osquery, OSSEC, auditd, GRR and built automated alert response processes
* Ran automated vulnerability management with Nexpose and Nessus, hardened AMI pipelines with Packer and managed the HackerOne bug bounty program
* Provided DFIR analysis for compromised hosts and ran company-wide phishing campaigns with GoPhish to keep everyone honest
* Migrated all security microservices to Kubernetes on CoreOS, operating dedicated security clusters with Terraform and Chef

## Newedge (Societe Generale)
Onboarding Analyst
Singapore
2014

* Client onboarding and account lifecycle management for a global investment bank; automated communication workflows to reduce manual processing


# Education

*2012*-*2014*
: Bachelor of Science; Cyber Forensics, Information Security Management and Business Information Systems; Murdoch University

*2011*-*2012*
: Diploma; Information Technology; Kaplan Singapore

*2005*-*2009*
: Guildford Grammar School


*References available upon request.*
