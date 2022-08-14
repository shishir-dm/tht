## Compliance

In this document we will describe how to follow HIPAA standards when building our SaaS service on AWS EKS

AWS enables covered entities and their business associates subject to HIPAA to securely process, store, and transmit PHI (protected health information) but only after executing a standardized Business Associate Addendum (BAA) with AWS. Customers who execute an AWS BAA may use any AWS service in an account designated as a HIPAA Account, but they may only process, store and transmit PHI using the HIPAA-eligible services defined in the AWS BAA.

For more details see [Architecting for HIPAA Amazon EKS](https://docs.aws.amazon.com/whitepapers/latest/architecting-hipaa-security-and-compliance-on-amazon-eks/architecting-hipaa-security-and-compliance-on-amazon-eks.html)

### Shared responsibility
AWS is responsible for the security and compliance **of** the Cloud, which refers to the infrastructure that runs all of the services offered in the AWS Cloud. This infrastructure consists of the hardware, software, networking, and facilities that run AWS Cloud services.

Customers are responsible for the security and data safeguards **in** the Cloud, which consists of customer configured systems and services provisioned on AWS. For HIPAA-related safeguards, the customer has responsibility for all system components, which includes AWS resources included in, or connected to their containerized workload which process, store, or transmit PHI.

For more details see [AWS Shared Responsibility Model](https://docs.aws.amazon.com/whitepapers/latest/architecting-hipaa-security-and-compliance-on-amazon-eks/aws-shared-responsibility-model.html)

### Our setup

By default AWS EKS, VPCs, Load Balancers are already created to be HIPAA compliant for AWS so our job is mostly to ensure that we as developers strictly adhere to the [best practices](https://docs.aws.amazon.com/whitepapers/latest/architecting-hipaa-security-and-compliance-on-amazon-eks/securing-an-amazon-eks-deployment.html) defined by AWS for EKS workloads.

Each of individual section of these best practices such as IAM, Pod Security, Network Segmentation, Host hardenning, Logging, Auditting etc can be discussed and met. Individual solutions have been currently excluded from the scope of this exercise.