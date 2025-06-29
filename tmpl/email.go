package emailtemplate

const PinVerificationEmail = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" 
  "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html dir="ltr" lang="en">
  <head>
    <link rel="preload" as="image" href="{{ .LogoURL }}" />
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
    <meta name="x-apple-disable-message-reformatting" />
  </head>
  <body style="background-color:#ffffff;font-family:HelveticaNeue,Helvetica,Arial,sans-serif">
    <table align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation"
      style="max-width:360px;background-color:#ffffff;border:1px solid #eee;border-radius:5px;box-shadow:0 5px 10px rgba(20,50,70,.2);margin-top:20px;margin:0 auto;padding:68px 0 130px">
      <tbody>
        <tr style="width:100%">
          <td>
            <img alt="Plaid" height="88" src="{{ .LogoURL }}" width="212"
              style="display:block;outline:none;border:none;text-decoration:none;margin:0 auto" />
            <p style="font-size:11px;line-height:16px;color:#0a85ea;font-weight:700;font-family:HelveticaNeue,Helvetica,Arial,sans-serif;height:16px;letter-spacing:0;margin:16px 8px 8px 8px;text-transform:uppercase;text-align:center">
              Verify Your Identity
            </p>
            <h1 style="color:#000;display:inline-block;font-family:HelveticaNeue-Medium,Helvetica,Arial,sans-serif;font-size:20px;font-weight:500;line-height:24px;margin-bottom:0;margin-top:0;text-align:center">
              Enter the following code to your email verification.
            </h1>
            <table align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation"
              style="background:rgba(0,0,0,.05);border-radius:4px;margin:16px auto 14px;vertical-align:middle;width:280px">
              <tbody>
                <tr>
                  <td>
                    <p style="font-size:32px;line-height:40px;color:#000;display:inline-block;font-family:HelveticaNeue-Bold;font-weight:700;letter-spacing:6px;padding-bottom:8px;padding-top:8px;margin:0 auto;width:100%;text-align:center">
                      {{ .Code }}
                    </p>
                  </td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
    <p style="font-size:12px;line-height:23px;color:#000;font-weight:800;letter-spacing:0;margin:0;text-align:center;text-transform:uppercase">
      Securely powered by Pentaverses Auth.
    </p>
  </body>
</html>
`
