package crawler

import (
	"fmt"
	"venture-data-service/pkg/log"
	"venture-data-service/service/venture/dao"
)

var (
	contact_via_linkedIn = `To approach %v or any other investors, using LinkedIn is a very effective way. LinkedIn is a professional social network and there are many investors, entrepreneurs, and people in the cryptocurrency and blockchain industry who use it.
	To use LinkedIn to approach %v, you can follow these steps:
	1. Create a professional LinkedIn profile: Create a complete LinkedIn profile, ensuring that your profile contains your full personal information and work experience.
	2. Search for relevant investors and entrepreneurs: Search for investors and entrepreneurs relevant to the cryptocurrency and blockchain industry on LinkedIn by using keywords such as “blockchain investor” or “cryptocurrency venture capitalist”.
	3. Connect with these investors and entrepreneurs: Connect with relevant investors and entrepreneurs in the cryptocurrency and blockchain industry on LinkedIn by sending connection requests or through LinkedIn groups related to this industry.
	4. Send investment proposals: If you have successfully connected with investors and entrepreneurs, you can send your investment proposal through LinkedIn, including your presentation and other relevant documents.
	Note that LinkedIn is a useful tool to approach investors and entrepreneurs, but to attract their attention, you need a convincing pitch and clear benefits of your product or service.`
)

func GenContact() {
	repo := &dao.VentureRepo{}
	err := repo.GetVentureCodes()
	if err != nil {
		log.Println(log.LogLevelError, "GenContact: repo.GetVentureCodes()", err)
	}

	for _, venture := range repo.Ventures {

		linkWebsite := ""
		website, exist := venture.Socials["website"]
		if exist {
			// You can visit their website at https://www.paradigm.xyz/ to learn more about their investment criteria, portfolio companies, and investment team.
			linkWebsite = fmt.Sprintf("You can visit their website at %v to learn more about their investment criteria, portfolio companies, and investment team.", website)
		}

		normal := fmt.Sprintf(`If you are interested in approaching %v for investment, you can follow these steps:
	1. Research the fund: Before reaching out, it is important to understand the investment focus and portfolio of %v. %v
	2. Prepare your pitch: Once you have done your research, prepare a clear and concise pitch that outlines your business idea, team, traction, and why you think %v would be a good fit as an investor.
	3. Make contact: You can reach out to %v by filling out the contact form on their website or by reaching out to a member of their investment team through LinkedIn or another professional networking platform.
	4. Follow up: After you have made contact, it is important to follow up with %v to ensure that they have received your pitch and to answer any questions they may have. Be sure to keep your communication professional and respectful.
	5. Be patient: The investment process can take time, so it is important to be patient and allow %v time to evaluate your pitch and make a decision. It may also be helpful to continue building relationships with other investors and stakeholders in your industry in the meantime.
	Alternatively, you can also contact them via %v
	It is important to note that %v receives a large number of pitches and investment opportunities, so it may take some time to receive a response or move forward with the investment process`, venture.VentureName, venture.VentureName, linkWebsite, venture.VentureName, venture.VentureName, venture.VentureName, venture.VentureName, venture.VentureName, venture.VentureName)

		contact := dao.Contact{
			Normal:   normal,
			LinkedIn: fmt.Sprintf(contact_via_linkedIn, venture.VentureName, venture.VentureName),
		}
		venture.Contact = contact
		err := venture.UpdateContact()
		if err != nil {
			log.Println(log.LogLevelError, "", err.Error())
		}
	}
	fmt.Println("done")

}
