/* Easy way to load html 
TODO - change physical address and web address to chrome-extension://[runtime id]
Generated based on template at https://termsandconditionstemplate.com/generate/
*/
const tosText = `<h1 id="login-faster-tos" class="title_size">Terms of Service</h1>` + 
`<h2 id="welcome-to-login-faster">Welcome to Login Faster</h2>
<p>These terms and conditions outline the rules and regulations for the use of Login Faster’s Chrome Extension.</p>
<p>Login Faster is located at:</p>
<p>Placeholder Address City<br />
ZIP code - CA , United States</p>
<p>By accessing this chrome extension we assume you accept these terms and conditions in full. Do not continue to use Login Faster’s chrome extension if you do not accept all of the terms and conditions stated on this page.</p>
<p>The following terminology applies to these Terms and Conditions, Privacy Statement and Disclaimer Notice and any or all Agreements: “Client”, “You” and “Your” refers to you, the person accessing this chrome extension and accepting the Company’s terms and conditions. “The Company”, “Ourselves”, “We”, “Our” and “Us”, refers to our Company. “Party”, “Parties”, or “Us”, refers to both the Client and ourselves, or either the Client or ourselves. All terms refer to the offer, acceptance and consideration of payment necessary to undertake the process of our assistance to the Client in the most appropriate manner, whether by formal meetings of a fixed duration, or any other means, for the express purpose of meeting the Client’s needs in respect of provision of the Company’s stated services/products, in accordance with and subject to, prevailing law of United States. Any use of the above terminology or other words in the singular, plural, capitalisation and/or he/she or they, are taken as interchangeable and therefore as referring to same.</p>
<h2 id="your-use-of-the-services">Your Use of the Services</h2>
<p>Login Faster grants you a personal, non-transferable, non-exclusive, revocable, limited license to use and access the Services solely as permitted by these Terms. We reserve all rights not expressly granted to you by these Terms.</p>
<p>Except as permitted through the Services or as otherwise permitted by us in writing, your license does not include the right to:</p>
<ul>
<li>license, sell, transfer, assign, distribute, host, or otherwise commercially exploit the Services or Content;</li>
<li>modify, prepare derivative works of, disassemble, decompile, or reverse engineer any part of the Services or Content; or</li>
<li>access the Services or Content in order to build a similar or competitive website, product, or service.</li>
</ul>
<p>We reserve the right to modify, suspend, or discontinue the Services (in whole or in part) at any time, with or without notice to you. Any future release, update, or other addition to functionality of the Services will be subject to these Terms, which may be updated from time to time. You agree that we will not be liable to you or to any third party for any modification, suspension, or discontinuation of the Services.</p>
<h2 id="cookies">Cookies</h2>
<p>As a chrome extension, Login Faster does not use cookies to store information about you or track you.</p>`
document.getElementById('tos_tab').innerHTML = tosText;