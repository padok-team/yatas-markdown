<p align="center">
<img src="docs/auditory.png" alt="yatas-logo" width="30%">
<p align="center">

# YATAS Markdown
## Usage

Generates a report in the current directory in `report.md`

## Installation

Add to your config file:

```yaml
- name: "markdown"
  enabled: true
  type: "report"
  source: "github.com/StanGirard/yatas-markdown"
  version: "latest" # or a specific version
  description: "Genereates a markdown report"
```

## Example Report

| Id | Name | Status | 
| ---- | ---- | ------ | 
| AWS_ACM_001| ACM certificates are valid | ✅ | 
| AWS_ACM_002| ACM certificate expires in more than 90 days | ✅ | 
| AWS_ACM_003| ACM certificates are used | ✅ | 
| AWS_APG_001| ApiGateways logs are sent to Cloudwatch | ✅ | 
| AWS_APG_002| ApiGateways are protected by an ACL | ✅ | 
| AWS_APG_003| ApiGateways have tracing enabled | ✅ | 
| AWS_ASG_001| Autoscaling maximum capacity is below 80% | ✅ | 
| AWS_ASG_002| Autoscaling group are in two availability zones | ✅ | 
| AWS_BAK_001| EC2's Snapshots are encrypted | ❌ | 
| AWS_BAK_002| EC2's snapshots are younger than a day old | ❌ | 
| AWS_CFT_001| Cloudfronts enforce TLS 1.2 at least | ✅ | 
| AWS_CFT_002| Cloudfronts only allow HTTPS or redirect to HTTPS | ✅ | 
| AWS_CFT_003| Cloudfronts queries are logged | ✅ | 
| AWS_CFT_004| Cloudfronts are logging Cookies | ✅ | 
| AWS_CFT_005| Cloudfronts are protected by an ACL | ✅ | 
| AWS_CLD_001| Cloudtrails are encrypted | ✅ | 
