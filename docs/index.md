---
layout: default
title: Home
nav_order: 1
description: "Go GestPay Official Docs"
permalink: /
last_modified_date: 2022-10-21T21:54:08+0000
---

# Go GestPay Official Documentation

This is library is used for gestpay web service for the SOAP api


### About the project


WSCryptDecrypt web service is available on production and test servers and does not require any installation on the merchant’s server.

The merchant must implement – in the page(s) of the virtual store configured to handle payments – a call to the webservice which handles requests to use the Axerve Ecommerce Solutions encryption service.

To request the encryption service you can call the Encrypt method.

If the encryption operation is concluded correctly (TransactionResult tag value with OK), the encrypted data string returned by Axerve Ecommerce Solutions will be available by reading the value of the CryptDecryptString tag.

If not, the values of the ErrorCode and ErrorDescription tag will make it possible to identify the reasons that prevented the encryption operation. To request the decryption service it is necessary to call the Decrypt method, passing the shopLogin and CryptedString tags with the values communicated by Axerve Ecommerce Solutions.

The information containing the transaction result will be available by reading the information in the XML response file corresponding to the result of the transaction.


Axerve documentation: [https://api.axerve.com](https://api.axerve.com/)

