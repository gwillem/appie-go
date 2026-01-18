# Albert Heijn GraphQL Schema

Query Type: Query
Mutation Type: Mutation

## OBJECT Types (1124)

### AccessibleImage
An image with a11y attributes.

- `images`: 
- `altText`: String

### ActivatePersonalPromotionResponse
The response of the activate personal promotion mutation

- `message`: ActivatePersonalPromotionMessage
- `status`: String

### AddMemberCard
The returned value if the checked member-card is valid

- `addingPossible`: Boolean
- `errorCode`: String

### AddSMSReminderResult
Result type for Add SMSReminder

- `errorMessage`: String
- `status`: MutationResultStatus
- `reminderStatus`: AddSMSReminderStatus

### AdditionalReceiptItem
Items that are deducted or applied on the totalPrice of OrderCalculationReceipt

- `title`: String
- `type`: AdditionalReceiptItemType
- `price`: Money
- `amount`: Int
- `items`: 

### AddressComplete
AddressComplete

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `postalCode`: PostalCode
- `city`: String
- `countryCode`: String

### AddressOptions
Complete address

- `city`: 
- `street`: 

### AddressPartial
AddressPartial

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `postalCode`: PostalCode
- `city`: String
- `countryCode`: String

### AddressSearch
Address search result

- `address`: AddressComplete
- `partialAddress`: AddressPartial
- `options`: AddressOptions
- `isComplete`: Boolean
- `isValid`: Boolean
- `isDeliveryAvailable`: Boolean

### Advertisement
An advertisement containing several attributes needed to render it.

- `slotName`: String
- `id`: String
- `advertiserId`: String
- `orderId`: String
- `creativeId`: String
- `lineItemId`: String
- `type`: String
- `colors`: AdvertisementColors
- `title`: String
- `images`: 
- `cta`: AdvertisementCta
- `trackedImpressionCounter`: String
- `viewableImpressionCounter`: String
- `clickTag`: String
- `width`: Int
- `height`: Int

### AdvertisementColors
An Advertisement's colors.

- `background`: String
- `text`: String

### AdvertisementCta
An advertisement's CTA.

- `text`: String
- `links`: AdvertisementCtaLinks

### AdvertisementCtaLinks
An advertisement's CTA links.

- `href`: String
- `appLink`: String

### AdvertisementImage
An advertisement's image.

- `type`: String
- `title`: String
- `link`: String
- `width`: Int
- `height`: Int

### AdvertisementMetadata
Advertisement Metadata

- `slotName`: String
- `id`: String
- `advertiserId`: String
- `orderId`: String
- `lineItemId`: String
- `creativeId`: String
- `trackedImpressionCounterURL`: String
- `viewableImpressionCounterURL`: String
- `clickTagURL`: String
- `width`: Int
- `height`: Int
- `comment`: String

### AdviceMutationResult
Response on syncing Lifestyle Advice with the service

- `status`: MutationResultStatus
- `errorMessage`: String
- `errors`: 

### AllerhandeMagazine
Allerhande magazine edition

- `id`: Int
- `edition`: String
- `year`: Int
- `url`: String
- `thumbnailImageUrl`: String

### AlternativeProductForIngredient
Alternative product that was found for the scanned ingredient

- `product`: Product
- `quantityForServings`: Int

### AlternativeSection
Alternative product suggestions for the scanned ingredient

- `title`: String
- `description`: String
- `products`: 

### AnswerCheck
This property will be filled when a dropdown is a question that has a correct answer and

- `answer`: String
- `correctMessage`: String
- `incorrectMessage`: String

### AvailableLoyaltyCards
Loyalty cards available for the user

- `type`: LoyaltyCard
- `displayName`: String
- `description`: String
- `imageUrl`: String
- `placeHolderGradientColors`: 

### BancontactMobilePrepayments
Bancontact Payment Details

- `date`: String
- `shortDate`: String
- `amount`: Money

### BargainCategory

- `categoryId`: String
- `categoryName`: String
- `categoryImages`: 

### BargainItem

- `product`: Product
- `categoryTitle`: String
- `markdown`: BargainMarkdown
- `stock`: Int
- `bargainPrice`: BargainPrice

### BargainMarkdown

- `markdownType`: BargainMarkdownType
- `markdownExpirationDate`: Date
- `markdownPercentage`: Float

### BargainPrice

- `priceWas`: String
- `priceNow`: String

### BasicMutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String

### Basket
Type for basket

- `id`: String
- `summary`: BasketSummary
- `items`: 
- `itemsInOrder`: 
- `itemsInList`: 
- `canChangeDelivery`: Boolean
- `notes`: 

### BasketItem

- `id`: Int
- `product`: Product
- `quantity`: Int
- `originCode`: String
- `position`: Int
- `isStrikethrough`: Boolean
- `allocatedQuantity`: Int
- `isClosed`: Boolean

### BasketItemProduct
basket item product

- `id`: Int
- `product`: Product
- `quantity`: Int
- `originCode`: String
- `position`: Int
- `isStrikethrough`: Boolean

### BasketMutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: Basket

### BasketNote
type for vague basket item

- `description`: String
- `quantity`: Int
- `searchTerm`: String
- `originCode`: String
- `position`: Int
- `isStrikethrough`: Boolean

### BasketOrderDeliveryAddress
Order delivery address

- `type`: String
- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `zipCode`: Int
- `zipCodeNumber`: Int
- `zipCodeExtra`: String
- `city`: String
- `countryCode`: String

### BasketOrderItem
basket order item

- `id`: Int
- `quantity`: Int
- `isClosed`: Boolean
- `allocatedQuantity`: Int
- `originCode`: String
- `product`: Product
- `position`: Int

### BasketSummary
Basket summary type

- `price`: BasketTotalPrice
- `quantity`: Int
- `orderId`: Int
- `deliveryDate`: String
- `lastUserChangeDateTime`: String
- `deliveryStartTime`: String
- `deliveryEndTime`: String
- `shoppingType`: BasketShoppingType
- `isCancellable`: Boolean
- `state`: BasketOrderState
- `address`: BasketOrderDeliveryAddress

### BasketTotalPrice
Basket total price type

- `priceBeforeDiscount`: Money
- `priceAfterDiscount`: Money
- `totalPrice`: Money
- `discount`: Money

### BonusCategory
Bonus category

- `id`: String
- `title`: String
- `untranslatedTitle`: String
- `type`: BonusCategoryType
- `images`: 
- `promotions`: 

### BonusCategoryImage
Bonus category image

- `url`: String
- `height`: Int
- `width`: Int

### BonusLaneItems
Bonus Lane items that can be products or segments

- `products`: Product
- `promotions`: Promotion

### BonusPeriod
Bonus period defines date range of promotions by week

- `startDate`: String
- `endDate`: String
- `weekNumber`: Int
- `actualWeek`: Boolean
- `folder`: Folder
- `folders`: Folder
- `nextPeriodVisibleFrom`: String

### BonusRecipe
Bonus recipe type

- `id`: Int
- `recipe`: RecipeSummary

### BonusSegment
Bonus segment defines bonus

- `id`: String
- `hqId`: Int
- `title`: String
- `type`: BonusSegmentType
- `description`: String
- `images`: 
- `productCount`: Int
- `productIds`: 
- `promotionType`: BonusPromotionType
- `price`: BonusSegmentPrice
- `category`: String
- `activationStatus`: BonusSegmentActivationStatus
- `discount`: BonusSegmentDiscount
- `discountUnit`: BonusSegmentDiscountUnit
- `discountLabels`: 
- `discountShields`: 
- `spotlight`: Boolean
- `salesUnitSize`: String
- `availability`: BonusSegmentAvailability
- `subtitle`: String
- `smartLabel`: String
- `storeOnly`: Boolean
- `products`: 
- `product`: Product
- `webPath`: String
- `exceptionRule`: String
- `isTieredPromotion`: Boolean
- `singleProductPromotion`: Boolean
- `redemptionCount`: Int

### BonusSegmentAvailability
Bonus Segment availability

- `startDate`: String
- `endDate`: String
- `description`: String

### BonusSegmentDiscount
BonusSegmentDiscount

- `description`: String
- `theme`: BonusTheme
- `title`: String
- `extraDescriptions`: 

### BonusSegmentDiscountLabel
BonusSegmentDiscountLabel

- `code`: BonusSegmentDiscountLabelCodes
- `defaultDescription`: String
- `price`: Float
- `percentage`: Float
- `amount`: Float
- `count`: Int
- `freeCount`: Int
- `actualCount`: Int
- `deliveryType`: String
- `unit`: String

### BonusSegmentDiscountShield
BonusSegmentDiscountLabel

- `text`: 
- `emphasis`: BonusSegmentDiscountShieldEmphasis
- `theme`: BonusSegmentDiscountShieldTheme
- `defaultDescription`: String
- `title`: String
- `topText`: String
- `centerText`: String
- `bottomText`: String

### BonusSegmentDiscountUnit
BonusSegmentDiscountUnit

- `type`: String
- `count`: Float
- `rate`: Float
- `totalDiscount`: Float
- `totalPrice`: Float
- `unitsForFree`: Float

### BonusSegmentImage
BonusSegmentImage

- `url`: String
- `title`: String
- `height`: Int
- `width`: Int

### BonusSegmentPrice
BonusSegmentPrice

- `now`: Money
- `was`: Money
- `label`: String

### BusinessActivity
Returned member company business activities

- `sbiCode`: String
- `isMainSbi`: Boolean

### CalculatedPurchaseStamp
Calculated Purchase Stamp with the new total Price

- `bookletQuantity`: Int
- `totalPrice`: Money
- `discount`: Money

### CalculatedPurchaseStampPaymentAmount
Calculated discount for purchase stamps

- `bookletQuantity`: Int
- `discount`: Money
- `paymentAmount`: Money

### CaptchaConfiguration
To retrieve the captcha configuration

- `showCaptcha`: Boolean
- `captchaKey`: String

### CardError
Details of a consent provided in the MemberConsents

- `key`: String
- `message`: String
- `path`: String

### CardMutationResult
Result of member cards can also be requested

- `status`: MutationResultStatus
- `validationErrors`: 
- `result`: MemberLoyaltyCards

### CategoryImage

- `width`: Int
- `height`: Int
- `url`: String

### ChallengeAction
Possible action return by polling payment status. Probably not used but for typing completion's sake

- `transactionId`: String
- `authorizationId`: String
- `challengeType`: ChallengeType
- `data`: String
- `sdk`: ChallengeActionSdk

### ChallengeActionSdk
Challenge Action SDK

- `id`: String
- `minVersion`: String

### ChannelAvailabilityDateTimeRange
Customer care channel availability time range

- `startDateTime`: String
- `endDateTime`: String

### ChannelAvailabilityForDay
Customer care channel availability

- `status`: ChannelAvailabilityStatus
- `startDateTime`: String
- `endDateTime`: String
- `nextOpeningHours`: ChannelAvailabilityDateTimeRange

### ChapterMutationResult
Response on syncing Chapter Answers with the service

- `status`: MutationResultStatus
- `errorMessage`: String
- `errors`: 

### ChatChannel
Customer care chat channel

- `waitingTimeMinutes`: Int
- `employeeAvailability`: ChannelAvailabilityForDay
- `id`: String
- `type`: ChannelType
- `visibleOnPlatforms`: 
- `availability`: ChannelAvailabilityForDay
- `label`: String
- `waitTimeText`: String
- `hasIncident`: Boolean

### CheckoutATPError
When Order Limit or Stock Limits are hit an ATP Error is used to solve the basket

- `stockLimits`: 
- `orderLimits`: 

### CheckoutErrorData
Checkout Error Data

- `errorType`: CheckoutErrors
- `categoryName`: String
- `orderLines`: 
- `totalAmount`: Money
- `minimalAmount`: Money
- `mov`: Money

### CheckoutErrorResponse
Checkout error Response

- `code`: CheckoutErrors
- `message`: String
- `data`: 

### CheckoutFallbackResponse
Order

- `order`: Order
- `errorMessage`: String
- `status`: MutationResultStatus

### CheckoutOrderLine
Checkout Validation Order Line

- `product`: Product
- `count`: Int
- `available`: Int
- `limitType`: CheckoutATPLimitType

### CheckoutOrderSubmissionPayment
Order Submission Payment Information

- `paymentId`: String
- `externalReference`: String
- `action`: CheckoutPaymentAction
- `transactions`: 
- `mutation`: PaymentMutation
- `paymentStatus`: PaymentStatus

### CheckoutOrderSubmissionPaymentV4
Order Submission Payment Information

- `paymentId`: String
- `transactions`: 
- `mutation`: PaymentMutationV5
- `paymentStatus`: PaymentStatus
- `externalReference`: String
- `action`: CheckoutPaymentAction

### CheckoutOrderSubmissionResponse
Order submission response data

- `order`: Order
- `payments`: 

### CheckoutOrderSubmissionResponseV4
Order submission response data

- `order`: Order
- `payments`: 

### CheckoutOrderSubmitResult
Result type for submitting order

- `errorMessage`: String
- `status`: MutationResultStatus
- `data`: CheckoutOrderSubmissionResponse
- `errors`: 
- `atpError`: CheckoutATPError

### CheckoutOrderSubmitResultV4
Result type for submitting order

- `errorMessage`: String
- `status`: MutationResultStatus
- `data`: CheckoutOrderSubmissionResponseV4
- `errors`: 
- `atpError`: CheckoutATPError

### CheckoutOrderValidation
Checkout Order Validation Response

- `errors`: 
- `atpError`: CheckoutATPError

### CheckoutOverview
Calculation Receipt of an Order

- `preselectedPaymentMethod`: PaymentMethod
- `totalPrice`: Money
- `basketPrice`: Money
- `netPrice`: Money
- `totalBonusDiscount`: Money
- `paymentAmount`: Money
- `groceriesQuantity`: Int
- `additionalItems`: 
- `totalAdditionalItems`: Money
- `additionalInfo`: 
- `walletOptions`: CheckoutOverviewWalletOptions
- `prepayments`: 
- `personalAssets`: 
- `metadata`: CheckoutOverviewMetadata

### CheckoutOverviewMetadata
Metadata about the checkout overview

- `dctTokenId`: String
- `sddTokenId`: String
- `giftCardTokenId`: String
- `legacyDebitType`: DebitType

### CheckoutOverviewPrepayment
Checkout Overview Prepayments made

- `amount`: Money
- `description`: String
- `method`: PaymentMethod

### CheckoutOverviewPurchaseStampsBookletOption
Purchase Stamps Booklet Options

- `quantity`: Int
- `totalPrice`: Money
- `amount`: Money

### CheckoutOverviewWalletOptions
Wallet Options to apply to the order (settlements, purchase stamps)

- `settlementAmount`: Money
- `bookletOptions`: 
- `appliedBooklets`: Int
- `availableBooklets`: Int
- `appliedGiftCardAmount`: Float

### CheckoutPaymentTransaction
An transaction that has ocurred during a payment

- `id`: String
- `status`: PaymentTransactionStatus
- `paymentMethod`: PaymentMethod

### CheckoutPollAction
User needs to poll for the authorization status

- `mutationId`: String
- `pollIntervalInMs`: Int

### CheckoutPrepayments
Prepayments with the leftover amount of OrderCalculationReceipt

- `idealPayments`: 
- `bancontactMobilePayments`: 

### CheckoutRedirectAction
User is redirected to external location

- `uri`: String

### CheckoutWallet
Checkout Wallet Items that are applied

- `settlements`: CheckoutWalletSettlement

### CheckoutWalletOptions
Calculated Wallet Options

- `purchaseStampsPaymentAmount`: 

### CheckoutWalletSettlement
Previous Settlements that are applied to the order

- `amount`: Money

### CiamAvailableMfaSetting
The code and description for available MFA Settings

- `code`: String
- `description`: String

### CiamDeleteAccountMutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `errorCode`: CiamDeleteAccountMutationResultErrorCode

### CiamPasskey
Passkey details.
A passkey can be used by a member to log in.

@link https://developer.mozilla.org/en-US/docs/Web/API/Web_Authentication_API

- `id`: String
- `name`: String
- `createdAt`: DateTime
- `lastUsedAt`: DateTime
- `authenticatorName`: String

### CiamPasskeyRegisterFinishMutationhResult
Mutation result for passkey register finish

- `status`: MutationResultStatus
- `result`: String
- `errorMessage`: String
- `passkeySettings`: CiamPasskeySettings

### CiamPasskeyRegisterStartMutationResult
Get a registration ID and `publicKeyCredentialCreationOptions`,
which are needed to create a new passkey client-side.

- `status`: MutationResultStatus
- `errorMessage`: String
- `registrationId`: String
- `publicKeyCredentialCreationOptions`: String

### CiamPasskeySettings
Passkey settings

- `passkeyAllowed`: Boolean
- `passKeys`: 

### CiamPasskeyUpdateMutationResult
Result type for mutations that modify passkey settings (register, update, delete).

- `status`: MutationResultStatus
- `errorMessage`: String
- `passkeySettings`: CiamPasskeySettings

### CiamPasswordChangeMutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `errorCode`: CiamPasswordChangeMutationResultErrorCode

### CiamPhoneNumberSendCodeMutationResult
Returned result after ciam-phone-number-send-code mutation

- `status`: MutationResultStatus
- `sessionId`: String
- `errorMessage`: String

### CiamPhoneNumberVerifyCodeMutationResult
Returned result after ciam-phone-number-verify-code mutation

- `status`: MutationResultStatus
- `phoneNumber`: String
- `errorMessage`: String
- `errorCode`: CiamPhoneNumberVerifyCodeErrorCode

### CiamUpdateMfaSettingMutationResult
The mutation result from updating the MFA setting for the current member

- `status`: MutationResultStatus
- `errorMessage`: String
- `member`: Member

### CiamUserNameChangeMutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `errorCode`: CiamUserNameChangeMutationResultErrorCode

### CommunicationCategories
Available communication topics categories

- `importantInformation`: CommunicationCategory
- `additionalInformation`: CommunicationCategory

### CommunicationCategory
Communication topic category

- `topics`: 
- `groups`: 

### CommunicationConsent
This type represents a communication consent. E.i. a subscription to a specific topic via a specific channel.

- `memberId`: String
- `application`: String
- `channel`: CommunicationConsentChannel
- `topic`: String

### CommunicationConsentsResponse
This type represents the response all communication consents for the current user.

- `consents`: 

### CommunicationTopicDescription
Communication topic description that contains data that should be used for displaying particular topic

- `name`: String
- `description`: CommunicationTopicsDisplayProperties
- `loyaltyDescription`: CommunicationTopicsDisplayProperties
- `userInteractionEnabled`: Boolean
- `isMemberSubscribed`: Boolean

### CommunicationTopicGroup
Entity that represents communication topics that are grouped and displayed in a separated block

- `name`: String
- `displayName`: String
- `topics`: 

### CommunicationTopicsDescriptionResponse
The response for the communicationTopicsDescription query

- `categories`: CommunicationCategories
- `channel`: CommunicationConsentChannel
- `application`: String

### CommunicationTopicsDisplayProperties
Communication topic display properties, such as name and description

- `displayName`: String
- `displayText`: String

### CompanyAccount
Company account

- `memberId`: Int
- `companyName`: String
- `address`: CompanyAccountAddress
- `memberStatusCode`: CompanyMemberStatusCode

### CompanyAccountAddress
Address

- `city`: String
- `street`: String
- `houseNumber`: String

### Consents
Details of a consent provided in the MemberConsents

- `code`: String
- `version`: Int
- `granted`: Boolean
- `timestamp`: String

### ContentAccordion
Accordion component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAccordionData

### ContentAccordionButton
Content Link

- `color`: String
- `theme`: String
- `href`: String
- `target`: String
- `title`: String
- `mobileLink`: String
- `isExternalLink`: Boolean
- `isExternalMobileLink`: Boolean

### ContentAccordionButtons
Article component buttons

- `type`: ContentAccordionBlockType
- `primaryButton`: ContentAccordionButton
- `secondaryButton`: ContentAccordionButton

### ContentAccordionDocument

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentAccordionData

### ContentAccordionItem
Accordion item

- `title`: String
- `anchor`: String
- `body`: String
- `contentBlocks`: 

### ContentAccordionMedia
Accordion component media
Contains images or videos (Youtube or BlueBillyWig)

- `type`: ContentAccordionBlockType
- `media`: 

### ContentAccordionParagraph
Accordion component paragraph

- `type`: ContentAccordionBlockType
- `title`: String
- `subtitle`: String
- `text`: String
- `mobileText`: String
- `contentTheme`: String
- `titleAnalytics`: String
- `textAnalytics`: String
- `mobileTextAnalytics`: String

### ContentAccordionQuote
Accordion component quote

- `type`: ContentAccordionBlockType
- `quote`: String
- `author`: String
- `imageAlignment`: QuoteImageAlignmentType
- `imageSet`: ContentImage

### ContentAccordionTable
Accordion component table

- `type`: ContentAccordionBlockType
- `isNumbered`: Boolean
- `items`: 
- `mobileItems`: 

### ContentAdvertorial
Advertorial component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAdvertorialData

### ContentAdvertorialData
Advertorial Component Data

- `advertorials`: 

### ContentAdvertorialItem
Advertorial Component item

- `title`: String
- `link`: ContentLink
- `textColor`: String
- `backgroundColor`: String
- `images`: 

### ContentAhrmsSpotlightList
AHRMS SpotlightList Component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAhrmsSpotlightListData
- `pageSlots`: 
- `location`: String

### ContentAhrmsSpotlightListData
AHRMS SpotlightList Component Data

- `slots`: 

### ContentAllerhandeContentListLane
Content FAQ Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeContentListLaneData

### ContentAllerhandeContentListLaneData
Content Content List Lane Data

- `title`: String
- `groups`: 
- `image`: ContentImage
- `position`: String
- `imagePosition`: String
- `anchorLabel`: String

### ContentAllerhandeContentListLaneLinkGroup
Content Content List Lane Link Group

- `title`: String
- `links`: 

### ContentAllerhandeFAQLane
Content FAQ Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeFAQLaneData

### ContentAllerhandeFAQLaneData
Content FAQ Lane Data

- `title`: String
- `questions`: 
- `anchorLabel`: String

### ContentAllerhandeFAQLaneQuestion
Content FAQ Lane Question

- `question`: String
- `answer`: String

### ContentAllerhandeFlexPage
Allerhande Flex Page

- `id`: String
- `gridLanes`: 
- `groupHeader`: ContentAllerhandeRecipeGroupHeaderLaneData
- `meta`: ContentMobileComponentMetaObject

### ContentAllerhandeFpContentList
Content Allerhande Flex Page Content List

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentAllerhandeContentListLaneData

### ContentAllerhandeFpPageEntries
Super shops (part of page entries) component for allerhande flex page

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentPageEntriesData

### ContentAllerhandeFpProductLane
Product lane for allerhande flex page

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentProductLaneData

### ContentAllerhandeFpQueryableRecipeLane
Query recipe lane for allerhande flex page

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentQueryableRecipeLaneData

### ContentAllerhandeFpRecipeDetailCompound
Content Allerhande Flex Page Recipe Detail Compound

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentAllerhandeRecipeDetailLaneData

### ContentAllerhandeFpRecipeGrid
Content Allerhande Flex Page Recipe Grid Component

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentAllerhandeRecipeGridLaneData

### ContentAllerhandeFpSmartPromotions
ContentSmartLane component

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentBonusLaneData

### ContentAllerhandeFpTextImageItems
Content Allerhande Flex Page Text Image Items

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentAllerhandeTextImageItemsContentLaneData

### ContentAllerhandeFpVideoAllerhande
Content Allerhande Flex Page Video Allerhande Component

- `type`: ContentAllerhandeGridlaneComponentType
- `data`: ContentAllerhandeVideoLaneVideo

### ContentAllerhandeHighlightedThemesLane
Content Highlighted Themes Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeHighlightedThemesLaneData

### ContentAllerhandeHighlightedThemesLaneData
Content Highlighted Themes Lane Data

- `title`: String
- `subtitle`: String
- `links`: 
- `cta`: ContentAllerhandeLink

### ContentAllerhandeHighlightedThemesLaneLinks
Content Highlighted Themes Lane Links Data

- `title`: String
- `query`: RecipeThemeSearchParams
- `recipes`: 
- `count`: Int
- `link`: ContentAllerhandeLink

### ContentAllerhandeHighlightedThemesLaneLinksRecipe
Content Highlighted Themes Lane Links Recipe Data

- `id`: Int
- `title`: String
- `images`: 

### ContentAllerhandeImageCollectionLane
Content Image Collection Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeImageCollectionLaneData

### ContentAllerhandeImageCollectionLaneData
Content Image Collection Lane Data

- `bannerImage`: ContentAllerhandeImageCollectionLaneItem
- `images`: 

### ContentAllerhandeImageCollectionLaneImage
Content Image Collection Lane Image

- `image`: ContentImage
- `caption`: ContentAllerhandeImageCollectionLaneImageCaption

### ContentAllerhandeImageCollectionLaneImageCaption
Content Image Collection Lane Image Caption

- `label`: String
- `link`: ContentAllerhandeImageCollectionLaneImageCaptionLink

### ContentAllerhandeImageCollectionLaneImageCaptionLink
Content Image Collection Lane Image Caption Link

- `label`: String
- `url`: String
- `target`: String

### ContentAllerhandeImageCollectionLaneRecipeImage
Content Image Collection Lane Recipe Image

- `recipeId`: Int
- `recipe`: Recipe

### ContentAllerhandeLink
Content Link

- `url`: String
- `label`: String
- `target`: String

### ContentAllerhandeRecipeDetailLane
Content Recipe Detail Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeRecipeDetailLaneData

### ContentAllerhandeRecipeDetailLaneData
Content Recipe Detail Lane Data

- `title`: String
- `recipeId`: Int
- `recipe`: Recipe

### ContentAllerhandeRecipeGridLane
Content Recipe Grid Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeRecipeGridLaneData

### ContentAllerhandeRecipeGridLaneData
Content Recipe Grid Lane Data

- `title`: String
- `subtitle`: String
- `anchorLabel`: String
- `recipeQuery`: RecipeThemeSearchParams

### ContentAllerhandeRecipeGroupHeaderLane
Content Recipe Group Header Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeRecipeGroupHeaderLaneData

### ContentAllerhandeRecipeGroupHeaderLaneAnchor
Content Recipe Group Header Lane Anchor

- `id`: String
- `label`: String

### ContentAllerhandeRecipeGroupHeaderLaneData
Content Recipe Group Header Lane Data

- `title`: String
- `description`: String
- `anchors`: 
- `recipeId`: Int
- `recipe`: Recipe
- `image`: ContentImage

### ContentAllerhandeRecipeSearchBanner
Content Recipe Search Banner

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeRecipeSearchBannerData

### ContentAllerhandeRecipeSearchBannerData
Content Recipe Search Banner Data

- `searchHint`: String
- `searchHintShort`: String
- `quickSearchEntries`: 
- `recipeId`: Int
- `recipe`: Recipe

### ContentAllerhandeSpotlight
Content Spotlight

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeSpotlightData

### ContentAllerhandeSpotlightData
Content Spotlight Data

- `title`: String
- `subtitle`: String
- `spotlightItems`: 
- `cta`: ContentAllerhandeLink
- `anchorLabel`: String

### ContentAllerhandeSpotlightItemRecipe
Content Spotlight Item Recipe

- `type`: ContentAllerhandeSpotlightType
- `recipeId`: Int
- `item`: Recipe

### ContentAllerhandeSpotlightItemVideoBlueBillywig
Content Spotlight Item Video Blue Billywig

- `type`: ContentAllerhandeSpotlightType
- `blueBillywigId`: Int
- `item`: RecipeVideo

### ContentAllerhandeTextImageItemsContentLane
Content Text Image Items Content Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeTextImageItemsContentLaneData

### ContentAllerhandeTextImageItemsContentLaneData
Content Text Image Items Content Lane Data

- `contentLaneItems`: 
- `anchorLabel`: String

### ContentAllerhandeTextImageItemsContentLaneItem
Content Text Image Items Content Lane Item Data

- `title`: String
- `description`: String
- `recipeId`: Int
- `recipe`: Recipe
- `image`: ContentImage

### ContentAllerhandeThemeListCardCollectionLane
Content Theme List Card Collection

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeThemeListCardCollectionLaneData

### ContentAllerhandeThemeListCardCollectionLaneCard
Content Theme List Card Collection Card

- `title`: String
- `image`: ContentImage
- `links`: 
- `anchor`: ContentAllerhandeLink

### ContentAllerhandeThemeListCardCollectionLaneData
Content Theme List Card Collection Data

- `cards`: 

### ContentAllerhandeVideoLane
Content Video Lane

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentAllerhandeVideoLaneData

### ContentAllerhandeVideoLaneData
Content Video Lane Data

- `videos`: 

### ContentAllerhandeVideoLaneVideo
Content Video Lane Video

- `blueBillywigId`: Int
- `anchorLabel`: String
- `video`: RecipeVideo

### ContentArticle
Article component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentArticleData

### ContentArticleButton
Content Link

- `color`: String
- `theme`: String
- `href`: String
- `target`: String
- `title`: String
- `mobileLink`: String
- `isExternalLink`: Boolean
- `isExternalMobileLink`: Boolean

### ContentArticleButtons
Article component buttons

- `type`: ContentArticleBlockType
- `primaryButton`: ContentArticleButton
- `secondaryButton`: ContentArticleButton

### ContentArticleData
Article Component Data

- `documentUUID`: String
- `blocks`: 

### ContentArticleMedia
Article component media
Contains images or videos (Youtube or BlueBillyWig)

- `type`: ContentArticleBlockType
- `media`: 

### ContentArticleParagraph
Article component paragraph

- `type`: ContentArticleBlockType
- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `text`: String
- `textAnalytics`: String
- `mobileText`: String
- `mobileTextAnalytics`: String
- `contentTheme`: String

### ContentArticleQuote
Article component quote

- `type`: ContentArticleBlockType
- `quote`: String
- `author`: String
- `imageAlignment`: QuoteImageAlignmentType
- `imageSet`: ContentImage

### ContentArticleTable
Article component table

- `type`: ContentArticleBlockType
- `isNumbered`: Boolean
- `items`: 
- `itemsAnalytics`: 
- `mobileItems`: 
- `mobileItemsAnalytics`: 

### ContentBonusLane
Bonus Lane component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentBonusLaneData

### ContentBonusLaneData
Bonus Lane Component Data

- `title`: String
- `subtitle`: String
- `bonusLaneItems`: BonusLaneItems
- `bonusLaneType`: BonusLaneType
- `link`: ContentLink
- `supplierBoosted`: Boolean
- `serviceError`: Boolean

### ContentBrandHeader
BrandHeader component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentBrandHeaderData

### ContentBrandHeaderData
BrandHeader component data

- `imageSet`: ContentImage

### ContentButton
Button component

- `icon`: String
- `iconBackgroundTheme`: String
- `iconImage`: ContentImage
- `title`: String
- `subtitle`: String
- `link`: ContentButtonLink

### ContentButtonLink
Button component link

- `href`: String
- `target`: String

### ContentButtonlane
ButtonLane component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentButtonlaneData

### ContentButtonlaneData
ButtonLane component data

- `documentUUID`: String
- `id`: String
- `buttons`: 
- `theme`: String

### ContentContentHalfWidthMediaMonetizationLaneData
Half with media monetization lane component data

- `promotions`: 

### ContentCuratedList
Curated List Component Data

- `documentUUID`: String
- `title`: String
- `subtitle`: String
- `description`: String
- `personQuantity`: Int
- `image`: ContentImage
- `stickerImage`: ContentImage
- `items`: 

### ContentCuratedListItem
Curated List Item Component Data

- `vagueTerm`: String
- `quantity`: Int
- `unit`: String
- `productId`: String
- `productIdQuantity`: Int

### ContentCuratedListLane
Curated List Lane Component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentCuratedListLaneData

### ContentCuratedListLaneData
Curated List Lane Component Data

- `documentUUID`: String
- `title`: String
- `curatedLists`: 

### ContentCustomerAttributes
Customer attributes will be used to target customers
based on certain tags

- `tags`: 

### ContentCustomerCareLinkBox
Customer Care Link Box

- `title`: String
- `links`: ContentLink
- `seeMoreLink`: ContentLink

### ContentCustomerCareLinkBoxLane
Customer Care ListItem Component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentCustomerCareLinkBoxLaneData

### ContentCustomerCareLinkBoxLaneData
Customer Care Link Box Lane Data

- `linkBoxes`: 

### ContentCustomerTag
Customer tag

- `key`: String
- `description`: String
- `type`: String
- `attribute`: String
- `attributeValue`: String
- `domain`: String

### ContentDeliveryAddress
Delivery Address (Default is member's address)

- `title`: String
- `text`: String

### ContentDeliveryGrid
CMS Delivery Grid: User can select a delivery slot or pickup slot
It's used in ah.nl/kies-een-moment page.

- `weeklyOffers`: 
- `dailyOffers`: 
- `communications`: ContentDeliveryGridCommunications
- `configuration`: ContentDeliveryGridConfiguration
- `type`: String

### ContentDeliveryGridCommunications
Delivery Grid Communications

- `isUrgent`: Boolean
- `showPointer`: Boolean
- `delivery`: String
- `pickup`: String
- `pickupStore`: String
- `deliveryPickup`: String
- `deliveryAddress`: ContentDeliveryAddress

### ContentDeliveryGridConfiguration
Delivery Grid Configuration

- `obtainmentMethods`: 

### ContentDeliveryGridDailyOffer
Delivery Grid Daily Offer

- `day`: String
- `offer`: String

### ContentDeliveryGridObtainmentMethods
Delivery Grid Obtainment Obtainment Methods.
These methods are selectable by a user (if more than one are enabled for the user)

- `text`: String
- `title`: String
- `icon`: String
- `isVisible`: Boolean
- `type`: String
- `sustainableSlotsTitle`: String
- `sustainableSlotsDescription`: String

### ContentDeliveryGridWeeklyOffer
Delivery Grid Weekly Offer

- `offer`: String
- `startDate`: String
- `endDate`: String

### ContentError
Error Status for content module to be able to handle errors

- `code`: Int
- `message`: String
- `type`: ContentErrorTypes

### ContentFPAccordion
Accordion for flex page

- `type`: ContentFlexPageComponentType
- `data`: ContentAccordionData

### ContentFPAhrmsSpotlightList
AHRMS SpotlightList Component for Flex Page

- `type`: ContentFlexPageComponentType
- `data`: ContentAhrmsSpotlightListData

### ContentFPArticle
Article component for flex page

- `type`: ContentFlexPageComponentType
- `data`: ContentArticleData

### ContentFPBrandHeader
Brand header document for flex page

- `type`: ContentFlexPageHeaderType
- `data`: ContentBrandHeaderData

### ContentFPCuratedLists
Curated Lists for flex page

- `type`: ContentFlexPageComponentType
- `data`: ContentCuratedListLaneData

### ContentFPForm
Form component for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentFormData

### ContentFPHardcodedComponent
Hardcoded component

- `type`: ContentFlexPageComponentType
- `hardcodedComponentId`: String
- `hardcodedComponentType`: String

### ContentFPHighlightList
Highlight list component for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentHighlightListData

### ContentFPHolidayPromotion
Holiday promotion products for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentHolidayPromotionData

### ContentFPMarketingHeader
Marketing header document for flex page

- `type`: ContentFlexPageHeaderType
- `data`: ContentMarketingHeaderData

### ContentFPProductLane
Product lane for flex page

- `type`: ContentFlexPageComponentType
- `data`: ContentProductLaneData

### ContentFPQueryableRecipeLane
Queryable recipe lane for flex page

- `id`: String
- `type`: ContentFlexPageComponentType
- `data`: ContentQueryableRecipeLaneData

### ContentFPQuickEntry
Quick entry lane for flexpage

- `type`: ContentFlexPageComponentType
- `data`: ContentButtonlaneData

### ContentFPRichContent
RichContent component for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentRichContentData

### ContentFPShopHeader
Shop header document for flex page

- `type`: ContentFlexPageHeaderType
- `data`: ContentShopHeaderData

### ContentFPShorts
Video list component for flex page

- `type`: ContentFlexPageComponentType
- `title`: String
- `url`: String

### ContentFPSpotlightList
SpotlightList Component for Flex Page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentFPSpotlightListData

### ContentFPSpotlightListData

- `documentUUID`: String
- `title`: String
- `titleAnalytics`: String
- `spotlightCards`: 

### ContentFPSuperShops
Super shops (part of page entries) component for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentPageEntriesData

### ContentFPVideoList
Video list component for flex page

- `type`: ContentFlexPageComponentType
- `id`: String
- `data`: ContentVideoListData

### ContentFederatedDomains
Federated domain links

- `destinations`: String

### ContentFlexPage
Content Flex Page
(The page is used for both and mobile to display certain documents)

- `id`: String
- `submenu`: ContentSharedSubmenuData
- `header`: ContentFlexPageHeaderBase
- `components`: 
- `meta`: ContentMobileComponentMetaObject
- `seoText`: String
- `share`: ContentFlexPageShareData

### ContentFlexPageShareData
Flexpage share data

- `text`: String
- `url`: String

### ContentFooterLinks
CMS Footer Links: navigation in the footer

- `linkGroups`: 
- `id`: String
- `type`: String

### ContentFooterLinksGroup
Group of link items with a subject

- `links`: 
- `subject`: String

### ContentFooterLinksItem
Link item with title and anchor target

- `link`: String
- `title`: String
- `target`: String

### ContentForm
Form component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentFormData

### ContentFormCheckbox
Checkbox field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `media`: ContentMediaBlockBase

### ContentFormData
Form component data

- `formId`: String
- `title`: String
- `description`: String
- `mobileDescription`: String
- `action`: String
- `submitLabel`: String
- `successConfirmation`: String
- `successThankYou`: String
- `mobileSuccessThankYou`: String
- `expiredNotification`: ExpiredNotification
- `fields`: 

### ContentFormDropdown
Dropdown field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `renderType`: ContentFormDropdownRenderType
- `options`: 
- `media`: ContentMediaBlockBase
- `answerCheck`: AnswerCheck

### ContentFormEmail
Email field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `media`: ContentMediaBlockBase
- `pattern`: String

### ContentFormInformationBlock
Information field

- `type`: ContentFormFieldType
- `title`: String
- `description`: String
- `media`: ContentMediaBlockBase

### ContentFormTelephone
Telephone field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `media`: ContentMediaBlockBase
- `pattern`: String

### ContentFormTermsBlock
Terms block field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `required`: Boolean
- `termsLink`: ContentLink
- `media`: ContentMediaBlockBase

### ContentFormTextArea
Textarea field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `maxLength`: Int
- `media`: ContentMediaBlockBase

### ContentFormTextField
Text field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `placeholder`: String
- `description`: String
- `required`: Boolean
- `maxLength`: Int
- `pattern`: String
- `media`: ContentMediaBlockBase

### ContentFormUpload
Upload field

- `type`: ContentFormFieldType
- `name`: String
- `label`: String
- `description`: String
- `required`: Boolean
- `maxItems`: Int
- `media`: ContentMediaBlockBase

### ContentGridLane
ContentGridLane component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentGridLaneData

### ContentGridLaneData
Content grid lane Component Data

- `documentUUID`: String
- `title`: String
- `items`: 

### ContentGridLaneItem
ContentGridLane Component item

- `title`: String
- `body`: String
- `link`: ContentLink
- `images`: ContentImage

### ContentHalfWidthMediaMonetizationLane
Half with media monetization lane component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentContentHalfWidthMediaMonetizationLaneData

### ContentHeroBanner
HeroBanner component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentHeroBannerData

### ContentHeroBannerData
HeroBanner Component Data

- `title`: String
- `subtitle`: String
- `textColor`: String
- `images`: 

### ContentHeroCombinationLane
Content Hero Combination Lane component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentHeroCombinationLaneData

### ContentHeroCombinationLaneData
Hero Combination data

- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `image`: ContentImage
- `stickerImage`: ContentImage
- `link`: String
- `linkType`: LinkType
- `isExternalLink`: Boolean
- `lane`: ContentQueryableRecipeLaneDocument
- `theme`: String
- `documentUUID`: String

### ContentHighlight
Highlight item

- `title`: String
- `text`: String
- `documentUUID`: String
- `keyVisual`: ContentImage
- `theme`: String
- `sticker`: ContentImage
- `link`: ContentLink

### ContentHighlightList
ContentHighlightList component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentHighlightListData

### ContentHighlightListData
Highlight List Component Data

- `id`: String
- `documentUUID`: String
- `title`: String
- `highlights`: 
- `seeMoreLink`: ContentLink
- `renderType`: ContentHighlightLaneRenderType

### ContentHighlightListWeb
Highlight gallery component (A part of CMS Highlight List). Shows all highlights in a gallery view.

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentHighlightListData

### ContentHolidayPromotion
HolidayPromotion component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentHolidayPromotionData

### ContentHolidayPromotionData
HolidayPromotionData component data

- `title`: String
- `subtitle`: String
- `themeOptions`: String
- `products`: 

### ContentImage
Content image list of variants and image description

- `variants`: 
- `description`: String

### ContentImageSet
ImageSet type for Content module

- `url`: String
- `width`: Int
- `height`: Int
- `variant`: ContentImageVariantType

### ContentLink
Content Link

- `theme`: String
- `href`: String
- `target`: String
- `title`: String
- `mobileLink`: String
- `isExternalMobileLink`: Boolean
- `isExternalLink`: Boolean

### ContentLinkGroup
Link group component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentLinkGroupData

### ContentLinkGroupData
Link group data

- `documentUUID`: String
- `linkgroups`: 

### ContentLinkGroupItem
Link group item

- `title`: String
- `links`: 

### ContentLoyaltyCampaign

- `documentId`: String
- `faq`: ContentLoyaltyCampaignFAQ
- `onboardingImageSet`: ContentImage
- `entrypointImageSet`: ContentImage
- `entrypointTitle`: String

### ContentLoyaltyCampaignDocument
Content Loyalty Campaign document

- `loyaltyCampaign`: ContentLoyaltyCampaign

### ContentLoyaltyCampaignFAQ

- `questions`: 
- `seeMoreLink`: String
- `hasMoreQuestions`: Boolean

### ContentLoyaltyCampaignFAQItems

- `title`: String
- `text`: String

### ContentManualAccordion
Manual accordion returns items

- `title`: String
- `items`: 
- `type`: ContentAccordionType
- `documentUUID`: String

### ContentMarketingCarousel
Marketing Carousel component (A part of CMS PageEntries)

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentPageEntriesData

### ContentMarketingHeader
MarketingHeader component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentMarketingHeaderData

### ContentMarketingHeaderData
Marketing header data

- `content`: String
- `contentAnalytics`: String
- `contentTheme`: String
- `title`: String
- `titleAnalytics`: String
- `theme`: String
- `imageSet`: ContentImage
- `stickerImageSet`: ContentImage
- `primaryButton`: ContentArticleButton
- `secondaryButton`: ContentArticleButton
- `stickyButton`: Boolean
- `nonTransparentImage`: Boolean

### ContentMediaImage
Media image

- `type`: ContentMediaBlockType
- `imageSet`: ContentImage

### ContentMediaMonetizationSlot
Content media monetization slot

- `id`: String

### ContentMediaMonetizationTargeting
Content media monetization targeting

- `orderMode`: String
- `memberStateCode`: String
- `memberType`: String
- `googlePublishTags`: String
- `position`: String
- `category`: String

### ContentMediaVideoBlueBillyWig
Content media bbw video

- `type`: ContentMediaBlockType
- `imageSet`: ContentImage
- `videoId`: String
- `videoData`: VideoBbw

### ContentMediaVideoYoutube
Media youtube video

- `type`: ContentMediaBlockType
- `imageSet`: ContentImage
- `videoId`: String

### ContentMegaMenuLinks
CMS MegaMenu Links: Header MegaMenu links groups

- `linkGroups`: 
- `id`: String
- `type`: String

### ContentMegaMenuLinksGroup
Group of link items with a subject

- `links`: 
- `subject`: String

### ContentMegaMenuLinksItem
Link item with title and anchor target

- `link`: String
- `title`: String
- `target`: String

### ContentMobileArticleDocument
ContentArticleDocument

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseMobileCMSComponentType
- `title`: String
- `body`: String

### ContentMobileButtonLane
ButtonLane component

- `id`: String
- `meta`: ContentMobileComponentMetaObject
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentButtonlaneData

### ContentMobileComponentMetaObject
Meta object for mobile component

- `targetedHeaders`: String
- `promoModel`: String
- `promoOrchestrator`: String
- `promoIsManual`: Boolean

### ContentMobileCuratedLists
ContentMobileCuratedLists Component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentCuratedListLaneData

### ContentMobileImageSet
ImageSet type for Content mobile module TODO: change this into the main content image set

- `link`: String
- `width`: Int
- `height`: Int
- `mimeType`: String
- `variant`: ContentImageVariantType

### ContentMobileNextBestAction
Next Best Action

- `link`: String
- `isExternalLink`: Boolean
- `title`: String
- `titleAnalytics`: String
- `theme`: String
- `images`: ContentImage
- `stickerImages`: ContentImage
- `documentUUID`: String
- `campaignTag`: String
- `propositionTag`: String

### ContentMobileNextBestActionCard
ContentMobileNextBestActionCard component

- `meta`: ContentMobileComponentMetaObject
- `nextBestAction`: ContentMobileNextBestAction

### ContentMobileOverblijversBoxesDocument
Content Overblijvers Boxes Document

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseMobileCMSComponentType
- `overblijversBoxes`: 

### ContentMobilePageEntries
ContentMobilePageEntries component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseMobileCMSComponentType
- `renderType`: RenderType
- `view`: String
- `seeMoreLink`: ContentLink
- `title`: String
- `pageEntries`: 

### ContentMobilePageEntry
Page entry

- `link`: String
- `linkType`: LinkType
- `isExternalLink`: Boolean
- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `theme`: String
- `imageSet`: ContentImage
- `stickerImageSet`: ContentImage
- `documentUUID`: String
- `campaignTag`: String

### ContentMobilePageTemplate
ContentMobilePageTemplate component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `documentUUID`: String
- `lanes`: 

### ContentMobilePageTemplateLane
Composable home lane

- `laneType`: String
- `laneId`: String

### ContentMobilePropositionHeader
ContentMobilePropositionHeader component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseMobileCMSComponentType
- `title`: String
- `subtitle`: String
- `landscapeImageSet`: ContentImage
- `portraitImageSet`: ContentImage
- `stickerImageSet`: ContentImage

### ContentMobileSmartPromotions
ContentSmartLane component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentBonusLaneData

### ContentMobileSpotlightList
ContentMobileSpotlightList Component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentSpotlightListData

### ContentOptIn
CMS Opt In

- `title`: String
- `introduction`: String
- `mobileIntroduction`: String
- `closingText`: String
- `mobileClosingText`: String
- `optIns`: 
- `type`: String

### ContentOptInCompound
Content Opt In Compound

- `id`: String
- `text`: String
- `mobileText`: String

### ContentOverblijversBox
Content Overblijvers Boxes

- `image`: ContentMobileImageSet
- `avatar`: ContentMobileImageSet
- `boxType`: String
- `title`: String
- `description`: String
- `allergies`: String
- `backgroundColor`: String

### ContentPage
CMS Page

- `submenu`: ContentSubmenu
- `sharedSubmenu`: ContentSharedSubmenu
- `components`: 
- `parents`: 
- `previewData`: ContentPreviewData
- `pageMeta`: ContentPageMeta
- `error`: ContentError

### ContentPageEntriesData
PageEntries Component Data

- `title`: String
- `seeMoreLink`: ContentLink
- `entries`: 
- `type`: String
- `meta`: ContentPromoComponentMetaObject

### ContentPageEntriesDocument
ContentPageEntriesDocument component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `renderType`: RenderType
- `view`: String
- `seeMoreLink`: ContentLink
- `title`: String
- `pageEntries`: 

### ContentPageEntry
Page entry item

- `href`: String
- `mobileLink`: String
- `isExternalLink`: Boolean
- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `theme`: String
- `images`: ContentImage
- `stickerImages`: ContentImage
- `documentUUID`: String
- `campaignTag`: String
- `linkType`: LinkType

### ContentPageEntryDocument
Page entry

- `link`: String
- `linkType`: LinkType
- `isExternalLink`: Boolean
- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `theme`: String
- `imageSet`: ContentImage
- `stickerImageSet`: ContentImage
- `documentUUID`: String
- `campaignTag`: String

### ContentPageMeta
Page meta data (Used in frontend for SEO Purposes)

- `title`: String
- `description`: String
- `index`: Boolean
- `follow`: Boolean
- `jsonLD`: String
- `canonical`: String
- `theme`: String

### ContentParent
Page parent (used for breadcrumbs)

- `name`: String
- `link`: String

### ContentPlaceholder
Content placeholder

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentPlaceholderData

### ContentPlaceholderData
Placeholder data

- `title`: String

### ContentPreviewData
Preview data is needed for CMS experience manager

- `begin`: String
- `end`: String
- `contentLink`: String
- `visible`: Boolean

### ContentProductCategoryHeader
SEO text component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentProductCategoryHeaderData

### ContentProductCategoryHeaderData
Shop header data

- `headerImage`: ContentImage

### ContentProductLane
ProductLane Component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentProductLaneData

### ContentProductLaneData
Productlane Component Data

- `products`: 
- `promotions`: 
- `renderType`: ContentProductLaneRenderType
- `subtitle`: String
- `title`: String
- `anchorLabel`: String
- `link`: ContentLink
- `serviceError`: Boolean

### ContentProductListDocument
ContentProductList document

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `documentUUID`: String
- `data`: ContentProductLaneData

### ContentProductSegmentLaneData
ProductSegmentLane Component Data

- `segmentId`: String
- `showInStockProducts`: Boolean
- `products`: 
- `scrollable`: Boolean
- `subtitle`: String
- `title`: String
- `link`: ContentLink
- `serviceError`: Boolean

### ContentProductSuggestion
Product suggestions for curated list components

- `name`: String
- `description`: String
- `label`: String
- `selectedProduct`: ContentQuantifiedProduct
- `alternativeProducts`: ContentProductSuggestionAlternative

### ContentProductSuggestionAlternative
The alternative product suggestions with ContentQuantifiedProduct

- `title`: String
- `description`: String
- `products`: ContentQuantifiedProduct

### ContentPromoComponentMetaObject
Meta object for promo components (supershops & page entries)

- `promoModel`: String
- `promoOrchestrator`: String
- `promoIsManual`: Boolean

### ContentPromotion
Content promotion object

- `slot`: ContentMediaMonetizationSlot
- `targeting`: ContentMediaMonetizationTargeting
- `type`: String

### ContentPropositionHeader
ContentPropositionHeader component

- `id`: String
- `type`: ContentCMSComponentType
- `anchorId`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `previewData`: ContentPreviewData
- `data`: ContentPropositionHeaderData

### ContentPropositionHeaderData
ContentPropositionHeader component data

- `title`: String
- `subtitle`: String
- `landscapeImage`: ContentImage
- `portraitImage`: ContentImage
- `sticker`: ContentImage

### ContentPropositionHeaderDocument
ContentPropositionHeaderDocument component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `title`: String
- `subtitle`: String
- `theme`: String
- `landscapeImageSet`: ContentImage
- `portraitImageSet`: ContentImage
- `stickerImageSet`: ContentImage

### ContentQuantifiedProduct
The product and its quantity
If the product exist then it should always have a quantity

- `quantity`: Int
- `product`: Product

### ContentQueryableRecipeLane
Queryable recipe lane component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentQueryableRecipeLaneData

### ContentQueryableRecipeLaneData
Smartlane Queryable Recipe lane data

- `title`: String
- `link`: ContentLink
- `recipeQuery`: RecipeThemeSearchParams
- `items`: 
- `serviceError`: Boolean

### ContentQueryableRecipeLaneDocument
ContentQueryableRecipeLane document

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentQueryableRecipeLaneData

### ContentResourceBundle
CMS Resource bundle which contains a list of key/values as labels for translations

- `id`: String
- `labels`: 
- `locale`: String

### ContentResourceBundleLabel
Resource bundle label

- `key`: String
- `value`: String

### ContentRichContent
Rich content component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentRichContentData

### ContentRichContentData
Rich content data

- `documentUUID`: String
- `title`: String
- `products`: 
- `itemsTitle`: String
- `body`: String
- `searchTerm`: String
- `link`: ContentLink
- `cardImage`: ContentImage
- `media`: ContentMediaBlockBase

### ContentRootComponent
Root Component (e.g. Submenu, Component)

- `id`: String
- `type`: ContentComponentType
- `name`: String
- `items`: 
- `previewData`: ContentPreviewData

### ContentSearchAccordion
Search accordion returns items and hasMoreItems

- `title`: String
- `searchTerm`: String
- `items`: 
- `hasMoreItems`: Boolean
- `seeMoreLink`: ContentLink
- `type`: ContentAccordionType
- `documentUUID`: String

### ContentSeoText
SEO text component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentSeoTextData

### ContentSeoTextData
Shop header data

- `seoText`: String

### ContentSharedSubmenu
Shared Submenu component

- `name`: String
- `id`: String
- `previewData`: ContentPreviewData
- `data`: ContentSharedSubmenuData

### ContentSharedSubmenuData
Submenu Component Data

- `parent`: ContentSharedSubmenuItem
- `items`: 

### ContentSharedSubmenuItem
Submenu navigation item

- `label`: String
- `selected`: Boolean
- `href`: String
- `mobileLink`: String

### ContentShopHeader
ShopHeader component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentShopHeaderData

### ContentShopHeaderData
Shop header data

- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `theme`: String
- `imageSet`: ContentImage
- `stickerImageSet`: ContentImage
- `primaryButton`: ContentArticleButton
- `nonTransparentImage`: Boolean

### ContentShopInShop
Shop in Shop component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `type`: ContentCMSComponentType
- `componentStatus`: ContentCMSComponentStatus
- `previewData`: ContentPreviewData
- `data`: ContentShopInShopData

### ContentShopInShopData
Shop in Shop Component Data

- `documentUUID`: String
- `title`: String
- `items`: ContentShopInShopItem

### ContentShopInShopItem
Shop in Shop item

- `title`: String
- `link`: ContentLink
- `bgColor`: String
- `images`: ContentImage
- `sticker`: String

### ContentSpotlightCard
Spotlight Item

- `documentUUID`: String
- `title`: String
- `titleAnalytics`: String
- `subtitle`: String
- `subtitleAnalytics`: String
- `link`: ContentLink
- `renderType`: ContentSpotlightVisualRenderType
- `visualV2`: ContentSpotlightVisualV2
- `sticker`: ContentImage

### ContentSpotlightCoverVisual
Spotlight Cover Visual

- `image`: ContentImage
- `shortImageV2`: ContentImage

### ContentSpotlightEmptyVisual
Spotlight Empty Visual

- `image`: ContentImage

### ContentSpotlightList
SpotlightList Component

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentSpotlightListData

### ContentSpotlightListData
SpotlightList Component Data

- `documentUUID`: String
- `title`: String
- `titleAnalytics`: String
- `spotlightCards`: 

### ContentSpotlightPackshotVisual
Spotlight Packshot Visual

- `image`: ContentImage
- `theme`: String

### ContentSpotlightVideoVisual
Spotlight Video Visual

- `image`: ContentImage
- `loopVideo`: Boolean
- `asset`: ContentImageSet
- `vastXml`: String
- `bbwUrl`: String
- `metadata`: ContentVideoMetadata

### ContentSubmenu
Submenu component

- `name`: String
- `id`: String
- `previewData`: ContentPreviewData
- `data`: ContentSubmenuData

### ContentSubmenuData
Submenu Component Data

- `items`: 
- `header`: ContentSubmenuHeader

### ContentSubmenuHeader
Submenu header

- `title`: String
- `focusPoint`: ContentSubmenuFocusPoint
- `image`: ContentImage

### ContentSubmenuItem
Submenu navigation item

- `label`: String
- `expanded`: Boolean
- `selected`: Boolean
- `href`: String
- `items`: 

### ContentSuperShops
Supershops component (A part of CMS PageEntries)

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentPageEntriesData

### ContentThemeConfiguration
Themes configuration

- `theme`: String
- `primaryColor`: String
- `secondaryColor`: String
- `textColor`: String
- `backgroundImageSet`: ContentImage
- `pantryTheme`: PantryTheme

### ContentThemesConfiguration
Themes configurations

- `themesConfiguration`: 

### ContentTitle
Content Title

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentTitleData

### ContentTitleData
Content Title Data

- `title`: String
- `hasOffset`: Boolean

### ContentUsp
Content usp

- `name`: String
- `link`: String

### ContentUspGroup
Content usp group

- `meta`: ContentMobileComponentMetaObject
- `type`: ContentBaseTargetedContentCMSDocumentType
- `id`: String
- `title`: String
- `themeColor`: String
- `link`: String
- `uspList`: 

### ContentVideo
Video item

- `videoTitle`: String
- `videoSubtitle`: String
- `videoType`: String
- `videoId`: String
- `videoStillImageSet`: ContentImage
- `videoData`: VideoBbw

### ContentVideoCarousel
Video carousel component (A part of CMS Video List). Shows max 7 videos in a carousel view.

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentVideoListData

### ContentVideoGrid
Video grid component (A part of CMS Video List). Shows all videos in a grid view.

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData
- `data`: ContentVideoListData

### ContentVideoList
ContentVideoList component

- `meta`: ContentMobileComponentMetaObject
- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `data`: ContentVideoListData

### ContentVideoListData
Video List Component Data

- `id`: String
- `documentUUID`: String
- `title`: String
- `items`: 
- `seeMoreLink`: ContentLink

### ContentVideoMetadata
VideoMetadata type for Content module

- `error`: String
- `start`: String
- `firstQuartile`: String
- `secondQuartile`: String
- `thirdQuartile`: String
- `complete`: String

### Conversation
A conversation

- `id`: String
- `memberContext`: ConversationMemberContext
- `handlerContext`: ConversationHandlerContext
- `status`: ConversationStatus
- `events`: 

### ConversationClosedEvent
Conversation event corresponding to a closed event

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType

### ConversationFileUrl
Short living url for files

- `id`: String
- `url`: String

### ConversationFiles
File with urls

- `files`: 

### ConversationHandedOverEvent
Conversation event corresponding to when a conversation has been handed over

- `id`: String
- `createdAt`: String
- `handler`: ConversationHandlerType
- `type`: ConversationEventType

### ConversationHandlerCloseRequestedEvent
Conversation event corresponding to a requested close event from the handler

- `id`: String
- `createdAt`: String
- `handler`: ConversationHandlerType
- `type`: ConversationEventType

### ConversationHandlerContext
Handler context

- `isTyping`: Boolean

### ConversationHandlerFileSentEvent
File event from handler

- `id`: String
- `createdAt`: String
- `fileId`: String
- `type`: ConversationEventType
- `handler`: ConversationHandlerType
- `handlerName`: String
- `contentType`: String

### ConversationHandlerLinkListSentEvent
A link list send by the handler of the conversation

- `id`: String
- `type`: ConversationEventType
- `createdAt`: String
- `handler`: ConversationHandlerType
- `handlerName`: String
- `links`: 

### ConversationHandlerMessage
Handler message message

- `text`: String

### ConversationHandlerMessageSentEvent
Conversation event corresponding to a message sent by the message handler

- `id`: String
- `createdAt`: String
- `handler`: ConversationHandlerType
- `handlerName`: String
- `message`: ConversationHandlerMessage
- `type`: ConversationEventType
- `automated`: Boolean
- `viaLiveChat`: Boolean

### ConversationHandlerReopenedEvent
Conversation event corresponding to a reopened event from the handler

- `id`: String
- `createdAt`: String
- `handler`: ConversationHandlerType
- `handlerName`: String
- `type`: ConversationEventType

### ConversationHandlerTableSentEvent
A table message send by the handler of the conversation

- `id`: String
- `type`: ConversationEventType
- `createdAt`: String
- `handler`: ConversationHandlerType
- `handlerName`: String
- `header`: 
- `rows`: 
- `footer`: 

### ConversationLinkListHyperLink
Hyper link send by the Digital Assistant

- `href`: String
- `text`: String
- `type`: ConversationHyperLinkType

### ConversationMemberCloseRequestedEvent
Conversation event corresponding to a requested close event from the member

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType

### ConversationMemberContext
Member context

- `memberId`: Int
- `name`: String
- `anonymous`: Boolean
- `isTyping`: Boolean
- `isActive`: Boolean
- `lastHeartbeatAt`: String

### ConversationMemberFileSentEvent
File event from member

- `id`: String
- `createdAt`: String
- `fileId`: String
- `type`: ConversationEventType
- `contentType`: String

### ConversationMemberMessageSentEvent
Conversation event corresponding to a message sent by a member

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType
- `text`: String

### ConversationMemberReopenRequestedEvent
Conversation event corresponding to a requested reopen event from the member

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType

### ConversationReopenRequestedEvent
Conversation event corresponding to a requesting a reopen

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType

### ConversationReopenedEvent
Conversation event corresponding to a reopened event

- `id`: String
- `createdAt`: String
- `type`: ConversationEventType

### ConversationSummary
Type for conversation summary

- `summary`: String

### CookBook
Short summary of a complete cook book from an user

- `id`: Int
- `intro`: String
- `moderated`: CookBookModerated
- `isShared`: Boolean
- `date`: CookBookDate
- `user`: CookBookUser
- `stats`: CookBookStats

### CookBookAddRecipeResult
Cookbook recipe report offensive result

- `status`: MutationResultStatus
- `result`: CookBookMemberRecipe
- `errorMessage`: String

### CookBookCollectedRecipe
Collected information

- `id`: Int
- `isFavourite`: Boolean
- `note`: String

### CookBookCollectedRecipeRemoveResult
Cookbook recipe report offensive result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### CookBookCollectedRecipeResult
Cookbook recipe report offensive result

- `status`: MutationResultStatus
- `result`: CookBookCollectedRecipe
- `errorMessage`: String

### CookBookDate
Date information on the cookbook

- `changed`: String
- `inserted`: String

### CookBookDeleteRecipeResult
Cookbook delete recipe result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### CookBookEditRecipeResult
Cookbook mutation results

- `status`: MutationResultStatus
- `result`: CookBookMemberRecipe
- `errorMessage`: String

### CookBookMemberMessage
Message from another cookbook member to a cookbook member

- `id`: Int
- `sender`: CookBookMemberMessageSender
- `sent`: String
- `subject`: String
- `message`: String
- `isRead`: Boolean

### CookBookMemberMessageOverview
Overview of a message

- `id`: Int
- `sender`: CookBookMemberMessageSender
- `sent`: String
- `subject`: String
- `isRead`: Boolean

### CookBookMemberMessageSender
Details of the user who sent the message, people can block and unblock users so make sure to cache/update correctly

- `id`: Int
- `name`: String
- `isBlocked`: Boolean

### CookBookMemberMessagesResult
Result for paginated messages

- `result`: 
- `page`: PageInfo

### CookBookMemberProfile
Cookbook profile information

- `intro`: String
- `avatar`: Int
- `isShared`: Boolean
- `hasMessaging`: Boolean
- `hasNotifications`: Boolean

### CookBookMemberRecipe
CookBook Member Recipe

- `id`: Int
- `public`: CookBookRecipe
- `private`: CookBookRecipePrivate

### CookBookMessagingBlockedMember
Blocked person for contacting a specific cookbook user

- `userId`: Int
- `name`: String

### CookBookRecipe
One specific public recipe

- `id`: Int
- `title`: String
- `meta`: CookBookRecipeMeta
- `ingredients`: 
- `times`: CookBookRecipeTimes
- `isSlim`: Boolean
- `isFavourite`: Boolean
- `tools`: String
- `preparation`: String
- `date`: CookBookDate
- `collected`: Int
- `source`: String
- `userId`: Int
- `user`: CookBookUser

### CookBookRecipeCountResult
Cookbook recipe count result, each recipe has its own counter for visits by members and is used for creating the popularity of the cookbook author

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### CookBookRecipeIngredient
Recipe ingredient

- `quantity`: String
- `unit`: String
- `name`: String

### CookBookRecipeListItem
Small details about top list recipe (cached for 1 hour)

- `id`: Int
- `title`: String
- `userId`: Int
- `user`: CookBookRecipeListItemUser

### CookBookRecipeListItemUser
Recipe owner details for top list (cached for 10 minutes)

- `name`: String

### CookBookRecipeMeta
Meta information about a recipe

- `servings`: Int
- `course`: String
- `cuisine`: String
- `protein`: String
- `dish`: String

### CookBookRecipePrivate
Private data a user can retrieve about its recipe

- `note`: String
- `isShared`: Boolean
- `state`: CookBookModerated

### CookBookRecipeReportOffensiveResult
Cookbook recipe report offensive result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### CookBookRecipeTimes
Times

- `preparation`: Int
- `waiting`: Int
- `cooking`: Int

### CookBookRecipes
All own recipes for a cookbook user (cached for ten minutes)

- `id`: Int
- `result`: 
- `page`: PageInfo
- `user`: CookBookUser

### CookBookRecipesCollected
Collected recipes for a cookbook user

- `id`: Int
- `result`: 
- `page`: PageInfo
- `user`: CookBookUser

### CookBookRecipesCollectedItem
Collected recipe

- `id`: Int
- `userId`: Int
- `title`: String
- `meta`: CookBookRecipeMeta
- `isFavourite`: Boolean
- `owner`: CookBookRecipesCollectedOwner

### CookBookRecipesCollectedMember
Collected recipes for the cookbook member

- `id`: Int
- `result`: 
- `page`: PageInfo
- `user`: CookBookUser

### CookBookRecipesCollectedOwner
Owner information

- `name`: String

### CookBookRecipesItem
Recipe summary

- `id`: Int
- `userId`: Int
- `title`: String
- `meta`: CookBookRecipeMeta
- `isFavourite`: Boolean
- `isShared`: Boolean
- `state`: CookBookModerated

### CookBookRecipesMember
All own recipes for current cookbook member

- `id`: Int
- `result`: 
- `page`: PageInfo
- `user`: CookBookUser

### CookBookSearchRecipeItem
Recipe overview for search results

- `id`: Int
- `title`: String
- `meta`: CookBookRecipeMeta
- `user`: CookBookSearchRecipeUser

### CookBookSearchRecipeResult
Found recipes based upon your search

- `result`: 
- `page`: PageInfo

### CookBookSearchRecipeUser
User info

- `id`: Int
- `name`: String

### CookBookSearchResult
Found books based upon your search

- `result`: 
- `page`: PageInfo

### CookBookStats
Statistics about the cookbook

- `recipes`: Int
- `collected`: Int
- `visited`: Int

### CookBookUser
Minimal user details

- `id`: Int
- `avatar`: Int
- `name`: String
- `hasMessaging`: Boolean

### CostOverview
CostOverview Of an Order

- `invoiceId`: String
- `primaryPaymentMethod`: PaymentMethod
- `groceryNetPrice`: Money
- `aggregatedTotalPriceBeforeDelivery`: Money
- `aggregatedTotalPriceAfterDelivery`: Money
- `aggregatedBasketPrice`: Money
- `aggregatedTotalPriceSettlement`: Money
- `groceriesQuantity`: Int
- `aggregatedTotalBonusDiscount`: Money
- `aggregatedTotalAdditionalItems`: Money
- `additionalItems`: 
- `settlements`: SettlementItems
- `payments`: CostOverviewPayments
- `refundAmount`: Money

### CostOverviewPaymentItem
Payment Item

- `title`: String
- `method`: PaymentMethod
- `amount`: Money

### CostOverviewPayments
CostOverview Payments

- `alreadyPaid`: 
- `toBePaid`: CostOverviewPaymentItem
- `aggregatedAlreadyPaidTotal`: Money

### CoverImageUrl
CoverImageUrl

- `original`: String
- `sizeAt800`: String
- `sizeAt600`: String
- `sizeAt200`: String

### CreateConversationResult

- `id`: String

### CreateDepositComplaintResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String

### CreateFeedbackResult
Create feedback result

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String
- `storeId`: Int

### CreateGeneralFeedbackResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String

### CreateProductComplaintResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String

### CreateProductReturnResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String

### CurrentBudget
The customer's current V2 budget.

- `id`: String
- `memberId`: Int
- `proposition`: String
- `type`: BudgetType
- `amount`: Money
- `carriedOver`: Money
- `realizedCosts`: Money
- `unRealizedCosts`: Money
- `availableToSpend`: Money
- `createdAt`: String
- `startAt`: String
- `startAtOfFirstBudgetInCurrentYear`: String

### Customer
The customer

- `id`: Int
- `defaultPickupLocationId`: Int
- `deliveryInstructions`: String
- `debitType`: DebitType
- `memberType`: CustomerMemberType

### CustomerCareEmployeeLoad
Current load of the Albert Heijn Customer Service Employee's

- `day`: String
- `times`: 
- `currentTime`: String

### CustomerCareSettlement
Settlment for the 'mijn verrekeningen' page

- `id`: String
- `order`: Order
- `product`: Product
- `title`: String
- `subtitle`: String
- `amount`: Money
- `cancelable`: Boolean
- `dateLabel`: String
- `feedbackLabel`: String
- `comment`: String
- `imageIds`: 

### CustomerCareSettlementConclusion
Settlement bucket conclusion

- `message`: String
- `amount`: Money

### CustomerCareSettlementsBank
Settlements with refund method bank

- `title`: String
- `refundMethod`: SettlementRefundMethod
- `settlements`: 
- `conclusion`: CustomerCareSettlementConclusion

### CustomerCareSettlementsDeliverer
Settlements with refund method deliverer

- `title`: String
- `refundMethod`: SettlementRefundMethod
- `settlements`: 
- `conclusion`: CustomerCareSettlementConclusion

### CustomerCareSettlementsInvoice
Settlements with refund method invoice

- `title`: String
- `refundMethod`: SettlementRefundMethod
- `settlements`: 
- `conclusion`: CustomerCareSettlementConclusion

### CustomerCareSettlementsPage
The customer care mijn verrekeningen page

- `settlements`: 

### CustomerCareSpecialDay
Special day for customer care operations

- `name`: String
- `dates`: 

### CustomerCareStore
Store

- `id`: Int
- `address`: CustomerCareStoreAddress

### CustomerCareStoreAddress
Store address

- `street`: String
- `houseNumber`: String
- `postalCode`: String
- `city`: String

### CustomerCareTime
Start and end time

- `start`: String
- `end`: String

### CustomerMutationResult
Result of performing mutation with customer data

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: Customer

### CustomerServiceButtonHit
Customer service button hit

- `title`: String
- `subTitle`: String
- `target`: String
- `href`: String
- `icon`: String
- `theme`: String

### CustomerServiceSearchAccordionResponse
Search with accordion response

- `hasNextPage`: Boolean
- `totalHits`: Int
- `accordion`: ContentSearchAccordion
- `pages`: 

### CustomerServiceSearchHit
Customer service search hit

- `documentType`: String
- `title`: String
- `url`: String
- `description`: String
- `text`: String

### CustomerServiceSearchResponse
Search response

- `buttonHits`: 
- `searchHits`: 
- `hasNextPage`: Boolean
- `totalHits`: Int

### CustomerServiceSuggestion
Search suggestion

- `title`: String
- `url`: String

### CustomerServiceSuggestionResponse
Search suggestion response

- `suggestions`: 

### DCTOnboardingResult
Onboarding result

- `responseCode`: PaymentsTokenOnboardingResponseCode
- `externalTokenReference`: String
- `pushAccountReceipt`: String

### DeleteMemberRecipeMutationResult
Delete member recipe result

- `errorMessage`: String
- `status`: MutationResultStatus

### Delivery
Delivery belonging to order

- `weekday`: String
- `address`: DeliveryAddress
- `addressSingleLine`: String
- `date`: String
- `dateDisplay`: String
- `dateDisplayShort`: String
- `delivererMessage`: String
- `deliveryLocationId`: Int
- `endTime`: String
- `id`: Int
- `method`: DeliveryMethod
- `offset`: String
- `pickup`: DeliveryPickup
- `pickupLocationId`: Int
- `shiftCode`: String
- `slot`: String
- `startTime`: String
- `status`: DeliveryStatus
- `trackAndTraceV2`: TrackAndTraceV2

### DeliveryAddress
Address delivered to

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `zipCode`: String
- `city`: String
- `countryCode`: String

### DeliveryBundle
Description for a Delivery Bundle subscription including terms and pricing

- `id`: ID
- `code`: String
- `description`: String
- `duration`: DeliveryBundleDuration
- `price`: DeliveryBundlePrice
- `priceBeforeDiscount`: DeliveryBundlePrice
- `deliveryDays`: DeliveryBundleDeliveryDays

### DeliveryBundleCancellation
Delivery Bundle subscription cancellation specification

- `earliestDate`: Date
- `latestDate`: Date
- `date`: Date

### DeliveryBundleNextSubscription
Description of next delivery bundle subscription

- `id`: ID
- `code`: String
- `description`: String
- `startDate`: String
- `price`: Money

### DeliveryBundlePaymentIDealIssuer
Available iDeal issuers

- `id`: Int
- `name`: String
- `brandUrl`: String

### DeliveryBundlePrice
Delivery Bundle pricing specification

- `duration`: Money
- `monthly`: Money

### DeliveryBundlePriceToChange
Pricing calculation to upgrade or downgrade Delivery Bundle

- `description`: String
- `discountDescription`: String
- `price`: Money
- `priceBeforeDiscount`: Money

### DeliveryBundleRenewal
Delivery Bundle next renewal pricing specification

- `date`: String
- `price`: Money

### DeliveryBundleSlot
When an order for a member with a delivery bundle subscription will be delivered

- `date`: String
- `from`: String
- `until`: String
- `isPreferred`: Boolean
- `isFavourite`: Boolean

### DeliveryBundleSubscription
A Delivery Bundle customer subscription

- `id`: ID
- `code`: String
- `deliveryDays`: DeliveryBundleDeliveryDays
- `price`: Money
- `renewal`: DeliveryBundleRenewal
- `endDate`: String
- `cancellation`: DeliveryBundleCancellation
- `isActive`: Boolean
- `isChangeable`: Boolean
- `nextSubscription`: DeliveryBundleNextSubscription

### DeliveryNotification
Delivery Notification

- `id`: String
- `type`: NotificationType
- `link`: NotificationLink
- `dateInfo`: String
- `icon`: NotificationIcon

### DeliveryPickup
Order pickup information

- `id`: String

### DeliverySlot
When an order will be delivered

- `date`: String
- `from`: String
- `until`: String
- `isPreferred`: Boolean

### DeliveryTrackTraceInformation
Delivery Information contains information about the delivery of an order

- `orderId`: Int
- `orderType`: DeliveryTrackTraceInformationOrderType
- `state`: DeliveryTrackTraceInformationState
- `pickupState`: DeliveryTrackTraceInformationState
- `deliveryState`: DeliveryTrackTraceInformationState
- `deliveryMessage`: String
- `messages`: 
- `messageContainsDeliveryEta`: Boolean
- `startEstimatedTimeArrival`: DateTime
- `endEstimatedTimeArrival`: DateTime
- `realisedDeliveryTime`: DateTime
- `stopsLeftBeforeThisOrder`: Int

### DeliveryTrackTraceInformationMessage
A message about the delivery that can be displayed

- `type`: DeliveryTrackTraceInformationMessageType
- `text`: String

### DepositComplaintDepositItem
The deposit item that the complaint is about

- `type`: DepositComplaintDepositType
- `label`: String
- `quantities`: DepositComplaintDepositItemQuantities
- `explanations`: 

### DepositComplaintDepositItemQuantities
Quantity properties for the deposit item

- `claimed`: Int
- `claimable`: Int
- `settledByDeliverer`: Int

### DepositComplaintExplanation
Explanations

- `type`: DepositComplaintExplanationType
- `label`: String

### DepositComplaintOrder
Order eligible to deposits

- `id`: String
- `transactionDateTime`: String
- `deliveryStartTime`: String
- `deliveryEndTime`: String
- `transactionAddress`: DepositComplaintOrderAddress
- `shoppingType`: DepositComplaintShoppingType
- `claimStatus`: DepositComplaintStatus

### DepositComplaintOrderAddress
The adress of the delivery

- `street`: String
- `houseNumber`: String
- `postalCode`: String
- `city`: String

### DeviceChallengeAction
Device Challenge that is performed by verifying the device fingerprint

- `transactionId`: String
- `authorizationId`: String
- `data`: String
- `sdk`: PaymentSDK
- `issuerId`: String

### EntryPoint
A type containing all information on an entry point, which is part of a component.

- `name`: String
- `contentVariant`: EntryPointContentVariant
- `properties`: 
- `metadata`: EntryPointMetadata
- `altText`: String

### EntryPointCardContent
A type containing the content of a card entry point. Contains elements such as a title,
a subtitle, an action label, an image, background colors and the target to route to on
interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `body`: String
- `label`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointCardExtendedContent
A type containing the content of a card extended entry point. Contains elements such as
a title, a subtitle, an action label, an image, background colors and the target to route
to on interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `body`: String
- `label`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointColorConfiguration
A type containing the configuration of color

- `font`: String
- `background`: 
- `labelFont`: String
- `labelBackground`: 

### EntryPointColors
A type containing colors used for the an entry point.

- `light`: EntryPointColorConfiguration
- `highContrast`: EntryPointColorConfiguration

### EntryPointComponent
A type containing a full entry point component. This is a renderable component by a
frontend consumer and acts as a lane or page with entries to other parts of the app or
website, supporting both deeplinks and weblinks.

- `name`: String
- `entryPoints`: 
- `content`: EntryPointComponentContent

### EntryPointComponentContent
Additional Component Content.

- `foregroundImageURL`: String
- `themeConfiguration`: ContentThemeConfiguration

### EntryPointDismissalMutationResult
A type representing the mutation result of an entry point dismissal. Either the dismissal was successful, or an error
occurred whilst trying to dismiss the entry point. In case of an error, an error message might be present.

- `status`: MutationResultStatus
- `errorMessage`: String

### EntryPointLargeHorizontalContent
A type containing the content of a large horizontal entry point. Contains elements
such as a title, a subtitle, an image, background colors and the target to route to on
interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointLargeVerticalContent
A type containing the content of a large vertical entry point. Contains elements
such as a title, a subtitle, an image, background colors and the target to route to on
interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointMetadata
A type containing metadata of an entry point.

- `group`: String
- `dismissible`: Boolean

### EntryPointProperty
A type containing a property of an entry point, with the value serialized as a string.

- `key`: String
- `type`: EntryPointPropertyType
- `value`: String

### EntryPointRegularHorizontalContent
A type containing the content of a regular horizontal entry point. Contains elements
such as a title, a subtitle, an image, background colors and the target to route to on
interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointRegularHorizontalExtendedContent
A type containing the content of a regular horizontal extended entry point. Contains
elements such as a title, a subtitle, an action label, an image, background colors and
the target to route to on interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `action`: String
- `label`: String
- `footer`: String
- `footerTarget`: EntryPointTarget
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointRegularVerticalContent
A type containing the content of a regular vertical entry point. Contains elements
such as a title, a subtitle, an image, background colors and the target to route to on
interaction.

- `variantType`: EntryPointContentVariantType
- `title`: String
- `subtitle`: String
- `imageUrls`: 
- `colors`: EntryPointColors
- `target`: EntryPointTarget

### EntryPointTarget
A type containing the target to route to on interaction with the entry point.

- `uri`: String
- `type`: EntryPointTargetLinkType

### EtaBlock
Track and Trace Eta Block

- `type`: String
- `range`: EtaRange

### EtaRange
Track and Trace Eta Block Range

- `start`: String
- `end`: String

### EvalueAccountBalance
The balance of the evalue account

- `value`: EvalueAccountBalanceAmount
- `reserved`: EvalueAccountBalanceAmount
- `available`: EvalueAccountBalanceAmount

### EvalueAccountBalanceAmount
Amount of the balance of the evalue account

- `amount`: Float
- `formatted`: String

### EvalueTransaction
Evalue Transaction

- `order`: Order
- `settlement`: SettlementV2
- `id`: String
- `createdDate`: DateTime
- `amount`: TransactionAmountCurrency
- `type`: EvalueTransactionType
- `status`: EvalueTransactionStatus
- `category`: String
- `description`: String
- `externalReference`: String
- `completionDate`: DateTime

### EvalueTransactionHistory
Evalue Transaction History

- `transactions`: 

### ExpandOrder
Expand Order

- `orderId`: Int
- `externalOrderId`: String
- `memberId`: Int
- `date`: String
- `deliveryStatus`: String
- `delivery`: String
- `payment`: String
- `orderLines`: ExpandOrderLine
- `billingAddress`: ExpandOrderAddress
- `shippingAddress`: ExpandOrderAddress
- `salesChannelId`: String
- `salesChannelName`: String

### ExpandOrderAddress
Address of the expand order

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `postalCode`: String
- `city`: String
- `countryCode`: String
- `type`: String
- `email`: String

### ExpandOrderLine
Expand Order lines

- `productId`: Int
- `productTitle`: String
- `quantity`: Int
- `price`: ExpandOrderLinePrice
- `resendable`: Boolean

### ExpandOrderLinePrice
Expand Order line price

- `price`: Money
- `discount`: Money
- `priceAfterDiscount`: Money

### ExperimentalBonusSegmentAvailability

- `startDate`: String
- `endDate`: String
- `description`: String

### ExperimentalMoney

- `amount`: Float
- `currency`: String

### ExperimentalProductAvailability

- `productId`: Int
- `isOrderable`: Boolean
- `isVisible`: Boolean
- `online`: ExperimentalProductAvailabilityIndication
- `offline`: ExperimentalProductAvailabilityIndication
- `unavailableForOrder`: ExperimentalProductUnavailableForOrderIndication
- `availabilityLabel`: String
- `maxUnits`: Int

### ExperimentalProductAvailabilityIndication

- `status`: ExperimentalProductAvailabilityStatus
- `availableFrom`: String

### ExperimentalProductCard

- `id`: ID
- `hqId`: Int
- `title`: String
- `brand`: String
- `category`: String
- `webPath`: String
- `minBestBeforeDays`: Int
- `salesUnitSize`: String
- `availabilityLabel`: String
- `interactionLabel`: String
- `icons`: ExperimentalProductIcon
- `isSample`: Boolean
- `highlight`: String
- `highlights`: String
- `shopType`: ExperimentalProductShopType
- `smartLabel`: String
- `maxUnits`: Int
- `privateLabel`: Boolean
- `hasListPrice`: Boolean
- `imagePack`: 
- `availability`: ExperimentalProductAvailability
- `tradeItem`: ExperimentalProductTradeItem
- `price`: ExperimentalProductPrice
- `discount`: ExperimentalProductDiscount
- `priceV2`: ExperimentalProductPriceV2
- `virtualBundleProducts`: ExperimentalVirtualBundleProduct
- `variant`: ExperimentalProductVariant
- `variants`: 
- `properties`: 

### ExperimentalProductDiscount

- `includeScratchCards`: Boolean
- `description`: String

### ExperimentalProductImage

- `angle`: ExperimentalProductImageAngle
- `small`: ExperimentalProductImageRendition
- `medium`: ExperimentalProductImageRendition
- `large`: ExperimentalProductImageRendition

### ExperimentalProductImageRendition

- `width`: Int
- `height`: Int
- `url`: String

### ExperimentalProductPrice

- `now`: ExperimentalMoney
- `was`: ExperimentalMoney
- `unitInfo`: ExperimentalProductUnitInfo
- `discount`: ExperimentalProductPriceDiscount

### ExperimentalProductPriceDiscount

- `segmentId`: Int
- `description`: String
- `type`: String

### ExperimentalProductPriceDiscountV2

- `segmentId`: Int
- `description`: String
- `type`: String
- `promotionType`: ExperimentalBonusPromotionType
- `segmentType`: ExperimentalBonusSegmentType
- `subtitle`: String
- `theme`: ExperimentalBonusTheme
- `tieredOffer`: 
- `availability`: ExperimentalBonusSegmentAvailability
- `wasPriceVisible`: Boolean
- `productCount`: Int
- `multipleItemPromotion`: Boolean
- `smartLabel`: String

### ExperimentalProductPriceV2

- `now`: ExperimentalMoney
- `was`: ExperimentalMoney
- `unitInfo`: ExperimentalProductUnitInfo
- `discount`: ExperimentalProductPriceDiscountV2

### ExperimentalProductProperty

- `code`: String
- `values`: 
- `icon`: ExperimentalProductPropertyIcon

### ExperimentalProductTradeItem

- `gtin`: String
- `gtinRevisions`: 

### ExperimentalProductUnavailableForOrderIndication

- `status`: ExperimentalProductUnavailableReason
- `availableFrom`: String

### ExperimentalProductUnitInfo

- `price`: ExperimentalMoney
- `description`: String

### ExperimentalProductVariant

- `type`: ExperimentalProductVariantType
- `label`: String
- `product`: ExperimentalProductCard

### ExperimentalSearchFacetConfig
Available Facet Configuration

- `label`: String
- `name`: String
- `type`: ExperimentalSearchFacetType

### ExperimentalSearchFacetOption
A facet option type

- `value`: PrimitiveTypeScalar
- `name`: String
- `matches`: Int
- `suboptions`: 
- `selected`: Boolean

### ExperimentalSearchFacetPayload
Facet payload type

- `label`: String
- `name`: String
- `type`: ExperimentalSearchFacetType
- `rangeSelectedFacetValues`: ExperimentalSearchRangeSelectedFacetValues
- `options`: 

### ExperimentalSearchFacetsPayload
Search facets payload contains all the available facets

- `facets`: 
- `quickFilters`: 

### ExperimentalSearchPayload
Search payload type

- `products`: 
- `facets`: ExperimentalSearchFacetsPayload
- `sponsoring`: ExperimentalSearchSponsoringPayload
- `totalFound`: Int
- `searchId`: String

### ExperimentalSearchRangeSelectedFacetValues
The min and/or max values for a range selected facet type

- `min`: PrimitiveTypeScalar
- `max`: PrimitiveTypeScalar

### ExperimentalSearchSponsoringPayload
SearchSponsoringPayload stores sponsoring related data (eg: 'auctionId')

- `auctionId`: String

### ExperimentalVirtualBundleProduct

- `quantity`: Int
- `product`: ExperimentalProductCard

### ExpiredNotification
This property will be filled when the form end date is in the past

- `title`: String
- `description`: String

### ExplanationWithLabel
Explanation with label

- `type`: ExplanationType
- `label`: String
- `byFeedbackTypes`: 

### FavoriteList
List containing a member's favorite products.

- `id`: Int
- `referenceId`: String
- `description`: String
- `totalSize`: Int
- `items`: 
- `products`: 
- `sharedBy`: FavoriteListSharedBy
- `imageUrl`: String
- `updatedAt`: String

### FavoriteListItem
Items on a favorite list.

- `id`: Int
- `productId`: Int
- `quantity`: Int

### FavoriteListItemV2
Items on a favorite list.

- `id`: String
- `productId`: Int
- `quantity`: Int

### FavoriteListMutationResultV2
Result of a mutation done to a favorite list.

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: FavoriteListV2

### FavoriteListV2
List containing a member's favorite products.

- `id`: String
- `description`: String
- `totalSize`: Int
- `items`: 
- `products`: 
- `sharedBy`: FavoriteListSharedBy
- `imageUrl`: String
- `updatedAt`: String

### Feedback
A feedback object

- `id`: String
- `type`: FeedbackType
- `label`: String
- `rating`: Int
- `shoppingType`: FeedbackShoppingType
- `loyaltyProgram`: FeedbackLoyaltyProgram
- `product`: Product
- `productDescription`: String
- `depositLabel`: String
- `quantity`: Int
- `settlementFraction`: Float
- `settlementsV2`: 

### FeedbackImageWithUrl
The feedback image with url

- `imageId`: String
- `url`: String

### FeedbackOption
FeedbackOption is a type that represents an option for feedback

- `label`: String
- `type`: FeedbackType
- `shoppingType`: FeedbackShoppingType
- `channel`: FeedbackChannel
- `loyaltyProgram`: FeedbackLoyaltyProgram

### FeedbackOptionList
FeedbackOption is a type that represents an option for feedback

- `label`: String
- `filter`: String
- `defaultValue`: FeedbackOption
- `children`: 

### FeedbackOrder
Order available for feedback

- `id`: String
- `shoppingType`: FeedbackOrderShoppingType
- `transactionDateTime`: String
- `deliveryStartTime`: String
- `deliveryEndTime`: String
- `transactionAddress`: FeedbackOrderAddress

### FeedbackOrderAddress
The order address can be a store location or a home address of a customer

- `city`: String
- `countryCode`: String
- `houseNumber`: String
- `postalCode`: String
- `street`: String

### FeedbackSettlement
The settlement linked to feedback

- `id`: Int
- `type`: SettlementType
- `quantity`: Int
- `label`: String
- `amount`: Money

### FeedbackSettlementV2
The settlement linked to feedback

- `id`: String
- `type`: SettlementType
- `quantity`: Int
- `label`: String
- `amount`: Money

### FeedbackTypeWithLabel
Feedback type with label

- `value`: FeedbackType
- `label`: String

### Folder
Folder

- `id`: Int
- `type`: FolderType
- `url`: String
- `title`: String
- `slug`: String
- `coverUrl`: String
- `coverImageUrl`: CoverImageUrl
- `pageCount`: Int
- `state`: FolderState
- `onlineAt`: String
- `offlineAt`: String
- `scheduleOnlineAt`: String
- `scheduleOfflineAt`: String
- `publicUrl`: String
- `metatags`: Metatag

### FreeDeliveryBonusSegments
FreeDelivery BonusSegments with the subscription

- `bonusSegments`: 

### Fulfillment
A fulfillment is an order with additional information.

- `costOverview`: CostOverview
- `orderId`: Int
- `totalPrice`: FulfillmentPrice
- `transactionCompleted`: Boolean
- `delivery`: FulfillmentDelivery
- `cancellable`: Boolean
- `cancellableByCustomerService`: Boolean
- `reopenable`: Boolean
- `modifiable`: Boolean
- `orderMethod`: Int
- `isOciOrder`: Boolean
- `isSubscriptionOrder`: Boolean
- `cancellationReason`: FulfillmentCancellationReason
- `approved`: Boolean
- `isAfterCutOff`: Boolean
- `closingDateTime`: String
- `shoppingType`: String
- `statusCode`: Int
- `statusDescription`: String

### FulfillmentCancellationReason
Reason why an order was cancelled. Only available for the fulfillment detail query.

- `type`: FulfillmentCancellationReasonType
- `messageNL`: String
- `messageBE`: String

### FulfillmentDelivery
Delivery information of a fulfillment.

- `deliveryMessage`: String
- `method`: FulfillmentDeliveryMethod
- `slot`: FulfillmentDeliveryTime
- `address`: FulfillmentDeliveryAddress
- `addressSingleLine`: String
- `status`: OrderDeliveryStatus
- `shiftCode`: String
- `homeShopCenterId`: Int
- `ride`: FulfillmentDeliveryRide
- `eta`: FulfillmentDeliveryEta

### FulfillmentDeliveryAddress
Address the orders should be delivered on.

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `city`: String
- `countryCode`: String
- `postalCode`: String

### FulfillmentDeliveryEta
Delivery ETA information of a fulfillment.

- `status`: String
- `estimated`: String
- `lower`: String
- `upper`: String

### FulfillmentDeliveryRide
Delivery ride information of a fulfillment.

- `number`: Int
- `sequenceNumber`: Int
- `homeShopCenterId`: Int

### FulfillmentDeliveryTime
Delivery date and time estimation of a fulfillment.

- `date`: String
- `dateDisplay`: String
- `dateDisplayShort`: String
- `timeDisplay`: String
- `dayDisplay`: String
- `startTime`: String
- `endTime`: String

### FulfillmentPrice
Price of a fulfillment.

- `beforeDiscount`: Money
- `afterDiscount`: Money
- `discount`: Money
- `totalPrice`: Money

### FulfillmentsResult
Result of the fulfillment query.

- `result`: 
- `page`: PageInfo

### GeoLocation
Geolocation

- `latitude`: Float
- `longitude`: Float

### GetNewsletterStatus
Contains newsletter subscription status

- `contactSubscription`: String
- `error`: String

### GradientPantryColorSet
Gradient color set (has primary and secondary colors, can be linear or radial)

- `fillStyle`: ThemeFillStyle
- `primaryColor`: String
- `secondaryColor`: String
- `textColor`: String

### GroceryItem
Item on a grocery list

- `id`: String
- `product`: Product
- `quantity`: Int
- `position`: Int
- `crossedOff`: Boolean

### GroceryList
The complete grocery list

- `id`: String
- `title`: String
- `type`: GroceryListType
- `groceryItems`: 
- `notes`: 

### GroceryListGetResult
GroceryList result we added the status code to indicate when a list is not found later on we can add more status codes for no persmission etc.

- `statusCode`: Int
- `groceryList`: GroceryList

### GroceryListResult
Result of creating a grocery list

- `status`: MutationResultStatus
- `errorMessage`: String
- `groceryList`: GroceryList

### IdealPayment
iDeal payment details

- `date`: String
- `shortDate`: String
- `amount`: Money

### ImageSet
Image Set

- `url`: String
- `width`: Int
- `height`: Int

### IntakeMutationResult
Response on syncing Intake Answers with the service, includes Lifestylecheck Chapter metas

- `status`: MutationResultStatus
- `errorMessage`: String
- `errors`: 
- `result`: 

### Invoice
Invoice specification

- `price`: Money
- `date`: String
- `id`: String
- `url`: String
- `orderId`: String
- `openAmount`: Money

### InvoicePagination
Pagination

- `totalSize`: Int
- `offset`: Int
- `pageSize`: Int

### InvoiceReport

- `dateFrom`: String
- `dateTo`: String
- `downloadUrl`: String

### InvoiceV2
Invoice type

- `orderId`: String
- `invoiceNumber`: String
- `invoiceDate`: String
- `amount`: Money
- `openAmount`: Money
- `companyName`: String
- `downloadUrl`: String

### Invoices
Result of found invoices

- `results`: 
- `pagination`: Pagination

### InvoicesV2
Invoices search results

- `results`: 
- `pagination`: InvoicePagination

### KeyValue
A key-value pair

- `key`: String
- `value`: KeyValueValue

### LifestyleAdvice
Lifestyle Advice

- `goalId`: String
- `goalName`: String
- `goalImageUrl`: String
- `title`: String
- `description`: String
- `minimumTopics`: Int
- `maximumTopics`: Int
- `chapters`: 
- `askConfirmation`: Boolean

### LifestyleAdviceChapter
Lifestyle Advice Chapter

- `id`: String
- `header`: String
- `description`: String
- `infoText`: String
- `topics`: 

### LifestyleAdviceTopicResponse
Lifestyle Advice Topic

- `title`: String
- `id`: String
- `iconType`: LifestyleAdviceIconType
- `label`: LifestyleAdviceTopicLabel
- `isTicked`: Boolean
- `imageUrl`: String
- `description`: String
- `deeplink`: String

### LifestyleCheckChapter
Lifestylecheck Chapter, full version

- `id`: String
- `status`: LifestyleCheckChapterStatus
- `isLastChapter`: Boolean
- `questions`: 

### LifestyleCheckChapterBasic
Lifestylecheck Chapter, meta version

- `id`: String
- `status`: LifestyleCheckChapterStatus
- `header`: String
- `description`: String
- `iconType`: LifestyleCheckChapterIconType

### LifestyleCheckChapterScore
Lifestylecheck Chapter Score

- `chapterId`: String
- `name`: String
- `iconType`: LifestyleCheckChapterIconType
- `averageScore`: Int

### LifestyleCheckCounterAnswer
Lifestylecheck Counter Answer

- `id`: String
- `answer`: String
- `imageUrl`: String

### LifestyleCheckCounterQuestion
Lifestylecheck Counter Question

- `id`: String
- `question`: String
- `description`: String
- `possibleAnswers`: 
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean

### LifestyleCheckIntake
Lifestylecheck Intake

- `questions`: 
- `status`: LifestyleCheckIntakeStatus

### LifestyleCheckIntakeAnswer
Lifestylecheck Intake Answer

- `id`: String
- `answer`: String
- `iconType`: LifestyleCheckIntakeAnswerIconType
- `imageUrl`: String

### LifestyleCheckIntakeQuestion
Lifestylecheck Intake Question

- `id`: String
- `question`: String
- `description`: String
- `possibleAnswers`: 
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean

### LifestyleCheckMultiSelectAnswer
Lifestylecheck Multi Select Answer

- `id`: String
- `answer`: String
- `isExclusive`: Boolean

### LifestyleCheckMultiSelectQuestion
Lifestylecheck Multi Select Question

- `id`: String
- `question`: String
- `description`: String
- `possibleAnswers`: 
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean
- `headerImageUrl`: String
- `defaultAnswerId`: String

### LifestyleCheckScore
Lifestylecheck Score

- `chosenGoal`: String
- `goalImageUrl`: String
- `totalAverageScore`: Int
- `chapterScores`: 

### LifestyleCheckSelectedAnswer
Lifestylecheck Single Select Answer

- `answerId`: String
- `count`: Int

### LifestyleCheckSingleSelectAnswer
Lifestylecheck Single Select Answer

- `id`: String
- `answer`: String

### LifestyleCheckSingleSelectQuestion
Lifestylecheck Single Select Question

- `id`: String
- `question`: String
- `description`: String
- `possibleAnswers`: 
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean
- `headerImageUrl`: String
- `defaultAnswerId`: String

### LifestyleCheckSliderSelectAnswer
Lifestylecheck Slider Select Answer

- `id`: String
- `answer`: String
- `imageUrl`: String

### LifestyleCheckSliderSelectQuestion
Lifestylecheck Slider Select Question

- `id`: String
- `question`: String
- `description`: String
- `possibleAnswers`: 
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean
- `defaultAnswerId`: String

### LifestyleErrorResponse
Lifestyle error Response

- `code`: LifestyleErrors
- `message`: String

### LifestyleGoalInfo
Lifestyle Hub Info

- `id`: String
- `name`: String
- `goalImageUrl`: String

### LifestyleHubInfo
Lifestyle Hub Info

- `goalInfo`: LifestyleGoalInfo
- `topics`: 

### LifestyleIntroduction
Lifestyle Introduction Info containing welcome and legal fields

- `welcome`: LifestyleWelcome
- `legal`: LifestyleLegal

### LifestyleLegal
Lifestyle legal info

- `title`: String
- `text`: String

### LifestyleWelcome
Lifestyle welcome info

- `headerImageUrl`: String
- `header`: String
- `body`: String
- `legalIntro`: String
- `uniqueSellingPoints`: 

### ListPrice
product list price

- `amount`: Float
- `currency`: String

### Load
Load

- `value`: Int
- `label`: String

### LoyaltyCards
Loyalty cards which are added to member

- `type`: LoyaltyCard
- `cardNumber`: String
- `barcodeNumber`: String

### LoyaltyChallenge
A type containing a loyalty challenge that a member can participate in. Contains information
on the challenge configuration and content. Does not contain any member-specific data, thus
is cacheable.

- `id`: Int
- `status`: LoyaltyChallengeStatus
- `name`: String
- `statusMessage`: String
- `theme`: ContentThemeConfiguration
- `info`: LoyaltyChallengeInfo
- `reward`: LoyaltyChallengeReward
- `steps`: 
- `overviewInfoLink`: LoyaltyChallengeInfoLink
- `accessibleImage`: AccessibleImage

### LoyaltyChallengeBanner
A type containing information for a banner to show in the challenge.

- `title`: String
- `body`: String
- `image`: 

### LoyaltyChallengeDuration
A type containing a duration of a loyalty challenge. Could represent the start and end date
of a step, or of the entire challenge, depending on where it is used.

- `start`: Date
- `end`: Date

### LoyaltyChallengeInfo
A type containing the information of a loyalty challenge. Contains various bits of content to
display to a member.

- `title`: String
- `description`: String
- `termsAndConditions`: LoyaltyChallengeTermsAndConditions
- `terms`: 
- `moreInfoLink`: LoyaltyChallengeInfoLink

### LoyaltyChallengeInfoLink
A type containing a link to more information about the challenge. Contains text, an image,
and a target to route to.

- `title`: String
- `description`: String
- `target`: LoyaltyChallengeInfoLinkTarget
- `accessibleImage`: AccessibleImage
- `theme`: ContentThemeConfiguration

### LoyaltyChallengeInfoLinkTarget
A type containing the target of a link to more information about the challenge.

- `type`: LoyaltyChallengeInfoLinkTargetType
- `link`: String

### LoyaltyChallengeOptInMutationResult
Result after opting a member in to the challenge.

- `status`: MutationResultStatus
- `challenge`: LoyaltyChallenge
- `errorMessage`: String

### LoyaltyChallengeOptOutMutationResult
Result after opting a member out of the challenge.

- `status`: MutationResultStatus
- `errorMessage`: String

### LoyaltyChallengeReward
A type containing the reward of a challenge. Contains various bits of content to display to a
member.

- `duration`: LoyaltyChallengeDuration
- `accessibleImage`: AccessibleImage
- `label`: String
- `subtitle`: String
- `body`: String

### LoyaltyChallengeStep
A type containing a step of a challenge.

- `id`: Int
- `status`: LoyaltyChallengeStepStatus
- `duration`: LoyaltyChallengeDuration
- `label`: String
- `subtitle`: String
- `body`: String
- `progress`: Int
- `maxProgress`: Int
- `progressDetails`: LoyaltyChallengeStepProgressDetails

### LoyaltyChallengeStepProgressDetails
Step formatted progress details.

- `ineligibleProgress`: String
- `ineligibleProgressDescription`: String
- `progress`: String
- `remainingProgress`: String
- `total`: String

### LoyaltyChallengeTerm
A type containing information on a term of a challenge.

- `title`: String
- `description`: String
- `accessibleImage`: AccessibleImage

### LoyaltyChallengeTermsAndConditions
A type wrapping the terms and conditions target

- `target`: LoyaltyChallengeInfoLinkTarget

### LoyaltyPointsBalance
A single loyalty points balance. This balance can be of a campaign of any type.

- `programId`: Int
- `balance`: Int

### LoyaltyPointsTransaction
A single loyalty points transaction. These transactions indicate additions or removals
of loyalty points from a member's loyalty points balance.

- `dateTime`: DateTime
- `reason`: LoyaltyPointsTransactionReason
- `type`: LoyaltyPointsTransactionType
- `amount`: Int

### LoyaltyPointsTransactionPage
A page containing transactions. This type also has a pagination object attached to it
with information on the pagination object that was passed along.

- `pagination`: LoyaltyPointsTransactionPagination
- `transactions`: 

### LoyaltyPointsTransactionPagination
A pagination object containing information on the page of transactions that was
returned.

- `offset`: Int
- `limit`: Int
- `key`: String

### LoyaltyProgram
Represents a loyalty program, including basic program information and content.

- `id`: Int
- `name`: String
- `type`: LoyaltyProgramType
- `savingPeriod`: LoyaltyProgramPeriod
- `redeemPeriod`: LoyaltyProgramPeriod
- `status`: LoyaltyProgramState
- `content`: LoyaltyProgramContent
- `products`: 

### LoyaltyProgramAboutContent
Basic information about the program that can be displayed.

- `title`: String
- `body`: String
- `image`: 

### LoyaltyProgramBannerContent
Information on the reusable banner of a program.

- `body`: String
- `title`: String
- `themeConfiguration`: ContentThemeConfiguration
- `image`: 
- `cta`: LoyaltyProgramCtaContent

### LoyaltyProgramContent
The main type containing all displayable program content.

- `title`: String
- `subtitle`: String
- `themeConfiguration`: ContentThemeConfiguration
- `about`: LoyaltyProgramAboutContent
- `stamps`: LoyaltyProgramStampContent
- `banner`: LoyaltyProgramBannerContent
- `heroBanner`: LoyaltyProgramHeroBannerContent
- `links`: LoyaltyProgramLinks
- `cta`: LoyaltyProgramCtaContent
- `faq`: LoyaltyProgramQuestionsContent
- `easterEgg`: String

### LoyaltyProgramCtaContent
Information on a CTA of a program.

- `text`: String
- `link`: String
- `linkType`: LoyaltyProgramCtaLinkType

### LoyaltyProgramHeroBannerContent
Information on the reusable hero banner of a program. Differs from a regular banner as
a hero banner does not contain any CTA.

- `title`: String
- `themeConfiguration`: ContentThemeConfiguration
- `image`: 
- `body`: String

### LoyaltyProgramLink
A url with a defined type

- `url`: String
- `type`: LoyaltyProgramCtaLinkType

### LoyaltyProgramLinks
The program's external links.

- `termsAndConditionsLink`: LoyaltyProgramLink
- `detailsLink`: LoyaltyProgramLink

### LoyaltyProgramPeriod
A type containing a period between two dates.

- `start`: Date
- `end`: Date

### LoyaltyProgramProduct
A redeemable product as part of a program. Contains the resolved product from the
product-resolver, as well as the overrides defined for the product. These overrides
should be applied on the resolved products' data. For example, if the overrides
contains a different product title, then it should override the NASA product's title.

- `points`: Int
- `overrides`: LoyaltyProgramProductOverrides
- `product`: Product

### LoyaltyProgramProductOverrides
Defines the content overrides for a LoyaltyProgramProduct.

- `basePrice`: Money
- `discountedPrice`: Money

### LoyaltyProgramQuestionContent
A Frequently Asked Question (FAQ) with an answer attached to it.

- `title`: String
- `link`: String
- `text`: String

### LoyaltyProgramQuestionsContent
Content related to the Frequently Asked Questions (FAQs) of a program. These are derived
from the global search on the website, e.g. when searching for "oven", you'll find some
customer service pages with FAQ questions on them.

- `searchTerm`: String
- `questions`: 
- `hasMoreQuestions`: Boolean
- `seeMoreLink`: String

### LoyaltyProgramStampContent
Information on the stamps involved in the program.

- `stampsPerCard`: Int
- `stampImage`: 
- `emptyImage`: 

### MaximumOrderValue

- `amount`: Float

### Member
Basic member information

- `isB2B`: Boolean
- `company`: MemberCompany
- `permissions`: 
- `contactSubscriptions`: 
- `analytics`: MemberAnalytics
- `consents`: MemberConsents
- `consentsToShow`: 
- `loyaltyCards`: MemberLoyaltyCards
- `isPhoneNumberVerified`: Boolean
- `isPhoneNumberMobile`: Boolean
- `mfaSetting`: String
- `availableMfaSettings`: 
- `id`: Int
- `name`: MemberName
- `gender`: Gender
- `dateOfBirth`: String
- `address`: MemberAddress
- `phoneNumber`: PhoneNumber
- `emailAddress`: EmailAddress
- `memberships`: 
- `cards`: MemberCards
- `customerProfileAudiences`: 
- `customerProfileProperties`: 

### MemberAddress
Member address

- `street`: String
- `houseNumber`: Int
- `houseNumberExtra`: String
- `postalCode`: PostalCode
- `city`: String
- `countryCode`: String

### MemberAnalytics
Member analytics data

- `idga`: String
- `idmon`: String
- `digimon`: String
- `idsas`: String
- `batch`: String
- `firebase`: String
- `sitespect`: String
- `optins`: 

### MemberBusinessActivity
Member business activities

- `businessActivityId`: String
- `isMainBusinessActivity`: Boolean

### MemberCards
All membercards with the card code for example bonus: '1234'

- `bonus`: String
- `gall`: String
- `airmiles`: String
- `xl`: String
- `wijndomein`: String
- `etos`: String

### MemberCompany
Member company details

- `id`: String
- `branchId`: String
- `name`: String
- `customOffersAllowed`: Boolean
- `addressInvoice`: MemberAddress
- `businessActivities`: 
- `vatId`: String

### MemberCompanyAddress

- `street`: String
- `houseNumber`: String
- `houseNumberExtra`: String
- `postalCode`: String
- `city`: String
- `country`: String

### MemberCompanyContactName

- `firstName`: String
- `lastName`: String
- `lastNamePrefix`: String

### MemberCompanyDetails

- `id`: Int
- `name`: MemberCompanyContactName
- `address`: MemberCompanyAddress
- `addressInvoice`: MemberCompanyAddress
- `emailAddress`: String
- `username`: String
- `phoneNumber`: String
- `dateOfBirth`: String
- `companyName`: String
- `chamberOfCommerceNumber`: String
- `branchNumber`: String
- `vatNumber`: String

### MemberCompanyRegistration
Search for a specific company within the dutch Chamber of Commerce

- `id`: String
- `name`: String
- `branchId`: String
- `countryCode`: String
- `businessActivities`: 
- `addressSearch`: AddressSearch

### MemberCompanyRegistrationResult
The returned object with the company registration and error state

- `company`: MemberCompanyRegistration
- `isKvkServiceDown`: Boolean

### MemberCompanyRegistrationSearch
Contains the rules and result against user password

- `id`: String
- `name`: String
- `branchId`: String
- `countryCode`: String
- `addressText`: String

### MemberCompanyRegistrationSearchResult
The result of the company search

- `companyRegistrations`: 
- `pageNumber`: Int
- `totalPages`: Int
- `isKvkServiceDown`: Boolean

### MemberConsents
Show all consents a member has given or declined

- `items`: 

### MemberDietPreferenceOption
Returns member diet preference option

- `value`: MemberDietPreferencesOptionType
- `label`: String
- `image`: ImageSet
- `selected`: Boolean

### MemberError
Returned error for invalid property

- `path`: String
- `error`: String
- `key`: String

### MemberFavouriteDishOption
Returns member meat preference option

- `value`: MemberDishPreferenceOptionType
- `label`: String
- `image`: ImageSet
- `selected`: Boolean

### MemberFoodPreferences
All member preferences pertaining food

- `protein`: 
- `favouriteDishTypes`: 
- `unwantedIngredients`: 

### MemberHasAccount
Check if an email-address is found on multiple domains

- `email`: String
- `domains`: 

### MemberLoyaltyCards
Member Loyalty Cards object

- `loyaltyCards`: 
- `availableLoyaltyCards`: 

### MemberMeatFrequencyPreferenceOption
Returns member meat preference option

- `value`: MemberMeatFrequencyPreferenceOptionType
- `label`: String
- `timesPerWeek`: Int
- `selected`: Boolean

### MemberMeatPreferenceOption
Returns member meat preference option

- `value`: MemberMeatPreferencesOptionType
- `label`: String
- `image`: ImageSet
- `href`: String
- `selected`: Boolean

### MemberName
Name of member containing first and lastname with prefix

- `first`: String
- `last`: String

### MemberNutritionPreferenceOption
Returns member meat preference option

- `value`: MemberNutritionPreferenceOptionType
- `label`: String
- `selected`: Boolean

### MemberPhoneNumberIsValid
Return type for memberPhoneNumberIsValid query

- `isValid`: Boolean
- `isMobile`: Boolean

### MemberPreferenceDishOptions
Returns member dish preference

- `selected`: Boolean
- `label`: String

### MemberPreferenceProteinOptions
Returns member protein preference

- `selected`: Boolean
- `label`: MemberProteinPreferencesOptionType

### MemberPreferences
Returns member food preference

- `food`: MemberFoodPreferences

### MemberRecipe
Member recipe is type of recipe that is created by members or scraped by members from third party sources.
Unlike regular Recipe, Member Recipe can be created, updated or delete. Regular Recipe is a read-only recipe
that's been created by dedicated team and is managed by them

- `id`: Int
- `title`: String
- `description`: String
- `ingredients`: 
- `preparation`: RecipePreparation
- `servings`: MemberRecipeServing
- `time`: RecipeTime
- `images`: 
- `author`: RecipeAuthor

### MemberRecipeIngredient
Member Recipe ingredient type

- `text`: String

### MemberRecipePreferencePreviewOption
Preference option that's going to be displayed if member haven't specified the preferences
and as the result doesn't have any recipes recommended

- `href`: String
- `image`: ImageSet
- `label`: String
- `value`: String

### MemberRecipePreferenceSelectedOption
Preference option that's been selected

- `label`: String
- `value`: String

### MemberRecipePreferences
All member recipe related preferences

- `meatPreferences`: 
- `meatFrequencyPreferences`: 
- `favouriteDishTypes`: 
- `unwantedIngredients`: 
- `dietPreferences`: 
- `nutritionPreferences`: 
- `familySize`: Int

### MemberRecipeServing
Member Recipe Serving Details

- `number`: Int
- `servingType`: RecipeServingType

### MemberUpdateMutationResult
Returned result after member update

- `status`: MutationResultStatus
- `result`: Member
- `errorMessage`: String
- `validationErrors`: 

### MembershipError
Details of a consent provided in the MemberConsents

- `path`: String
- `key`: String
- `message`: String

### MembershipResult
Result of trying to add membership to user

- `status`: MutationResultStatus
- `errors`: 

### MembershipsResult
Result of trying to add memberships to user

- `status`: MutationResultStatus
- `result`: 
- `errors`: 

### MessageCenterMessage
Represents a message center message

- `id`: String
- `isRead`: Boolean
- `isUrgent`: Boolean
- `content`: MessageCenterMessageContent
- `period`: MessageCenterMessageDisplayTimePeriod
- `optionalAttributes`: MessageCenterMessageOptionalAttributes

### MessageCenterMessageContent
Represents a message center message content related data

- `deeplink`: String
- `title`: String
- `category`: String
- `icons`: 

### MessageCenterMessageDisplayTimePeriod
Represents a message center message display time period data

- `displayDate`: String
- `section`: String
- `startDate`: String
- `endDate`: String

### MessageCenterMessageOptionalAttributes
Represents optional attributes of each message that should be used for analytics purposes

- `campaignId`: String
- `trackingCode`: String

### MessageCenterUnreadMessagesInfo
Represent message center unread messages info

- `unreadMessagesAmount`: Int
- `urgentUnreadMessagesAmount`: Int
- `showProfileIcon`: Boolean

### MessageCenterUpdateMessageMutationResult
Result of a Message Center message update

- `status`: MutationResultStatus
- `errorMessage`: String
- `messages`: 

### MessengerButton
Messenger button

- `label`: String
- `url`: String
- `target`: String
- `type`: ButtonType
- `disabled`: Boolean

### MessengerClosedEvent
Messenger event corresponding to a closed event sent by the message handler

- `id`: String
- `createdAt`: String
- `type`: MessengerConversationEventType

### MessengerComponentInteractionEvent
Messenger event corresponding to a message sent by a member

- `id`: String
- `createdAt`: String
- `type`: MessengerConversationEventType

### MessengerConversation
A messenger conversation with a subset of messenger events

- `events`: 
- `id`: String
- `status`: MessengerStatus
- `locale`: String
- `inputFieldEnabled`: Boolean
- `reopenable`: Boolean
- `handlerContext`: MessengerHandlerContext

### MessengerConversationStatus
Status of the conversation in the messenger

- `channelType`: MessengerChannelTypeEntrypoint
- `status`: MessengerStatus

### MessengerEventGroupsGroupedByDay
Messenger event grouped by the day of the event

- `groups`: 
- `date`: String
- `label`: String

### MessengerEventsGroupedByRules
Messenger event grouped by the grouping rules like pause period

- `events`: 
- `date`: String
- `labelV2`: String
- `footer`: String

### MessengerFileUrl
Short living url for files

- `id`: String
- `url`: String

### MessengerFiles
File with urls

- `files`: 

### MessengerGroupedConversation
A messenger conversation with a subset of messenger events

- `groupedEvents`: 
- `id`: String
- `status`: MessengerStatus
- `locale`: String
- `inputFieldEnabled`: Boolean
- `reopenable`: Boolean
- `handlerContext`: MessengerHandlerContext

### MessengerHandlerContext
Context about the assigned handler

- `isTyping`: Boolean

### MessengerHandlerFileSentEvent
Messenger file event from handler

- `id`: String
- `createdAt`: String
- `fileId`: String
- `type`: MessengerConversationEventType
- `handlerName`: String
- `handlerType`: MessengerHandlerType
- `avatar`: String
- `contentType`: String

### MessengerHandlerLinkListSentEvent
A link list send by the handler of the conversation

- `id`: String
- `type`: MessengerConversationEventType
- `createdAt`: String
- `handlerName`: String
- `handlerType`: MessengerHandlerType
- `links`: 

### MessengerHandlerMessageSentEvent
Messenger event corresponding to a message sent by the message handler

- `id`: String
- `createdAt`: String
- `handlerName`: String
- `handlerType`: MessengerHandlerType
- `message`: MessengerMessage
- `type`: MessengerConversationEventType
- `avatar`: String
- `automated`: Boolean

### MessengerHandlerTableSentEvent
A table message send by the handler of the conversation

- `id`: String
- `type`: MessengerConversationEventType
- `createdAt`: String
- `handlerName`: String
- `handlerType`: MessengerHandlerType
- `header`: 
- `rows`: 
- `footer`: 

### MessengerLinkListHyperLink
Hyper link send by the Digital Assistant

- `href`: String
- `text`: String
- `type`: MessengerHyperLinkType

### MessengerList
A list of messenger buttons

- `type`: String
- `text`: String

### MessengerMemberFileSentEvent
Messenger file event from member

- `id`: String
- `createdAt`: String
- `fileId`: String
- `type`: MessengerConversationEventType
- `contentType`: String

### MessengerMemberMessageSentEvent
Messenger event corresponding to a message sent by a member

- `id`: String
- `createdAt`: String
- `type`: MessengerConversationEventType
- `text`: String

### MessengerMessage
MessengerHandlerMessageSentEvent message

- `text`: String
- `button`: MessengerButton
- `suggestions`: 
- `list`: 

### MessengerReopenedEvent
Messenger event corresponding to a reopened event sent by the message handler

- `id`: String
- `createdAt`: String
- `type`: MessengerConversationEventType

### Metatag
Metatag

- `id`: Int
- `category`: String
- `value`: String

### MilesAccountResponse
Response after the request to create a miles account

- `status`: MilesStatus
- `message`: String

### MilesBalance
The Air Miles balance of the member. Might also contain an error if the member needs to re-auth at airmiles.nl

- `balance`: Int
- `errorState`: MilesErrorState

### MilesBalanceMutation
A type representing a mutation on an Miles balance. All values inside the mutation
represent a number of Miles.

- `old`: Int
- `new`: Int
- `mutation`: Int

### MilesCharity
A type containing a charity to where Miles can be donated to.

- `id`: String
- `content`: MilesCharityContent
- `validityPeriod`: MilesCharityPeriod

### MilesCharityContent
A type containing the displayable content of a charity where Miles can be donated
to. Contains information such as the charity's name, logo, summary, etc.

- `name`: String
- `imageUrl`: String
- `summary`: String
- `description`: String
- `gratitude`: String
- `options`: 
- `link`: String

### MilesCharityPeriod
A type containing a period in which the charity is valid (can be donated to).

- `start`: DateTime
- `end`: DateTime

### MilesDonation
A type containing a summary of a successful Miles donation to a charity.

- `requestId`: UUID
- `charityId`: String
- `timestamp`: DateTime
- `balanceMutation`: MilesBalanceMutation

### MilesDonationMutationResult
A type containing the result of making a donation using Miles.

- `status`: MutationResultStatus
- `errorMessage`: String
- `donation`: MilesDonation

### MilesTransaction
Transaction made with Air Miles

- `value`: Int
- `date`: String
- `domain`: String
- `description`: String

### MinimumOrderValue

- `amount`: Float
- `deadline`: String

### MissingBonusOrderLine
Recommendations for a bonus offer a member is missing out on in their order.
The recommendations are in the form of enriched order lines.

- `product`: Product
- `productQuantity`: Int
- `position`: Int
- `segmentId`: Int
- `title`: String
- `description`: String
- `offerQuantity`: Int

### Model
Product model

- `name`: String
- `version`: String

### Money
Represents a monetary value with currency.

- `amount`: Float
- `formattedV2`: String

### Mutation
Need to define mutation somehow, otherwise sometimes 'extend' fails

- `ahMemberOnboard`: MemberUpdateMutationResult
- `ahMemberUpdate`: MemberUpdateMutationResult
- `ahMemberUndoOnboarding`: BasicMutationResult
- `ahMemberUnsubscribeNewsletter`: BasicMutationResult
- `ahMemberConsentsUpdate`: MemberConsents
- `ahMemberAddMembership`: MembershipResult
- `ahMemberAddCard`: CardMutationResult
- `ahMemberAddCardEnum`: CardMutationResult
- `ahMemberUpdateCard`: CardMutationResult
- `ahMemberUpdateCardEnum`: CardMutationResult
- `ahMemberDeleteCard`: CardMutationResult
- `ahMemberDeleteCardEnum`: CardMutationResult
- `ahMemberAddMemberships`: MembershipsResult
- `ahMemberRemoveMembership`: BasicMutationResult
- `ahMemberRemoveSubscription`: BasicMutationResult
- `ahMemberAddSubscription`: BasicMutationResult
- `ahMemberUpdateProfile`: MemberUpdateMutationResult
- `basketItemsAdd`: BasketMutationResult
- `basketItemsUpdate`: BasketMutationResult
- `basketItemsDelete`: BasketMutationResult
- `basketItemReplace`: BasketMutationResult
- `basketMergeList`: BasketMutationResult
- `basketDeleteV2`: BasketMutationResult
- `bonusActivatePersonalPromotion`: ActivatePersonalPromotionResponse
- `businessPaymentMethodUpdate`: BasicMutationResult
- `checkoutConfirmOrder`: CheckoutOrderSubmitResult
- `checkoutConfirmFallbackSubmission`: CheckoutFallbackResponse
- `checkoutConfirmOrderV4`: CheckoutOrderSubmitResultV4
- `ciamDeleteAccount`: CiamDeleteAccountMutationResult
- `ciamPasskeyDelete`: CiamPasskeyUpdateMutationResult
- `ciamPasskeyRegisterFinish`: CiamPasskeyRegisterFinishMutationhResult
- `ciamPasskeyRegisterStart`: CiamPasskeyRegisterStartMutationResult
- `ciamPasskeyUpdate`: CiamPasskeyUpdateMutationResult
- `ciamPasswordChange`: CiamPasswordChangeMutationResult
- `ciamPasswordReset`: BasicMutationResult
- `ciamPasswordResetChange`: BasicMutationResult
- `ciamPasswordResetRequest`: BasicMutationResult
- `ciamPhoneNumberSendCode`: CiamPhoneNumberSendCodeMutationResult
- `ciamPhoneNumberVerifyCode`: CiamPhoneNumberVerifyCodeMutationResult
- `ciamSkipPasswordUpdate`: BasicMutationResult
- `ciamUpdateMfaSetting`: CiamUpdateMfaSettingMutationResult
- `ciamUserNameChange`: CiamUserNameChangeMutationResult
- `ciamUserNameChangeConfirm`: BasicMutationResult
- `memberProfileUpdate`: MemberUpdateMutationResult
- `memberRemoveCard`: BasicMutationResult
- `dummyMutation`: String
- `milesAccount`: MilesAccountResponse
- `milesDonate`: MilesDonationMutationResult
- `conversationHandlerSendMessage`: BasicMutationResult
- `conversationHandlerUpdateContext`: BasicMutationResult
- `conversationMemberHeartbeat`: BasicMutationResult
- `cookBookAddRecipeV2`: CookBookAddRecipeResult
- `cookBookCollectedRecipeRemoveV2`: CookBookCollectedRecipeRemoveResult
- `cookBookCollectedRecipeSetV2`: CookBookCollectedRecipeResult
- `cookBookDeleteRecipeV2`: CookBookDeleteRecipeResult
- `cookBookEditRecipeV2`: CookBookEditRecipeResult
- `cookBookMemberAddProfile`: CookBookMemberProfile
- `cookBookMemberDeleteMessage`: Boolean
- `cookBookMemberDeleteProfile`: CookBookMemberProfile
- `cookBookMemberSendMessage`: CookBookMemberMessage
- `cookBookMemberToggleBlocking`: Boolean
- `cookBookMemberUpdateProfile`: CookBookMemberProfile
- `cookBookRecipeCountV2`: CookBookRecipeCountResult
- `cookBookRecipeReportOffensiveV2`: CookBookRecipeReportOffensiveResult
- `customerDeliveryInstructionsUpdate`: CustomerMutationResult
- `deliveryAddSMSReminder`: AddSMSReminderResult
- `feedbackCreateFeedback`: CreateFeedbackResult
- `feedbackUpdateFeedback`: UpdateFeedbackResult
- `feedbackCreateCustomerSatisfactionScore`: CreateFeedbackResult
- `feedbackUpdateCustomerSatisfactionScore`: BasicMutationResult
- `feedbackCreateCustomerInteractionScore`: CreateFeedbackResult
- `feedbackUpdateCustomerInteractionScore`: BasicMutationResult
- `feedbackCreateCustomerContentScore`: CreateFeedbackResult
- `feedbackUpdateCustomerContentScore`: BasicMutationResult
- `feedbackCreateCustomerEffortScore`: CreateFeedbackResult
- `feedbackUpdateCustomerEffortScore`: BasicMutationResult
- `favoriteListProductsAddV2`: FavoriteListMutationResultV2
- `favoriteListProductsDeleteV2`: BasicMutationResult
- `favoriteListAddV2`: FavoriteListMutationResultV2
- `favoriteListChangeV2`: FavoriteListMutationResultV2
- `favoriteListDeleteV2`: BasicMutationResult
- `favoriteListShareWithSubAccountsV2`: BasicMutationResult
- `favoriteListRefreshShareWithSubAccountsV2`: SharedEntityMutationResult
- `favoriteListUnshareWithSubAccountsV2`: BasicMutationResult
- `loyaltyChallengeOptIn`: LoyaltyChallengeOptInMutationResult
- `loyaltyChallengeOptOut`: LoyaltyChallengeOptOutMutationResult
- `messengerGetOrCreateConversation`: CreateConversationResult
- `messengerCreateConversationV2`: CreateConversationResult
- `messengerSendMessage`: BasicMutationResult
- `messengerCloseConversation`: BasicMutationResult
- `messengerReopenConversation`: BasicMutationResult
- `messengerOrderCompletenessComplaint`: BasicMutationResult
- `messengerCreateComponentInteractionEvent`: BasicMutationResult
- `messengerSetTyping`: BasicMutationResult
- `messengerMemberHeartbeat`: BasicMutationResult
- `orderActiveDelete`: BasicMutationResult
- `orderBudgetV2Update`: OrderBudgetV2MutationResult
- `orderCancel`: BasicMutationResult
- `orderCheckin`: OrderCheckinMutationResult
- `orderDelete`: BasicMutationResult
- `orderExpandResend`: OrderResendResponse
- `orderMergeById`: Basket
- `orderReopen`: BasicMutationResult
- `orderRevert`: BasicMutationResult
- `orderSampleUpdate`: BasicMutationResult
- `orderSubmitActionCode`: BasicMutationResult
- `paymentsAuthorizeTransaction`: PaymentsAuthorizeResult
- `paymentsAddGiftCard`: PaymentsAddGiftCardResult
- `paymentsOnboardDCTToken`: PaymentsOnboardDCTTokenResult
- `paymentsUpdateCard`: PaymentsUpdateCardResult
- `paymentsPayPaymentRequest`: PaymentsCreatePaymentRequestPaymentResult
- `paymentsAuthorizePaymentRequest`: PaymentsCreatePaymentRequestPaymentResult
- `paymentsEnrollPersonalAsset`: PaymentsEnrollPersonalAssetResult
- `pickupLocationEnable`: PickupLocationEnableResult
- `purchaseStampCreateSecret`: PurchaseStampSecretMutationResult
- `purchaseStampSavingGoalDelete`: PurchaseStampSavingGoalMutationResult
- `purchaseStampSavingGoalSet`: PurchaseStampSavingGoalMutationResult
- `questionnaireFormSubmit`: BasicMutationResult
- `recipeCollectionAssignRecipeToCategories`: RecipeCollectionAssignRecipeToCategoriesResult
- `recipeCollectionCreateCategoryV2`: RecipeCollectionCreateCategoryResult
- `recipeCollectionDeleteCategoryV2`: RecipeCollectionDeleteCategoryResult
- `recipeCollectionRemoveRecipeFromCategoriesV3`: RecipeCollectionRemoveRecipeFromCategoriesResult
- `recipeCollectionUpdateCategoryV2`: RecipeCollectionUpdateCategoryResult
- `recipeDeleteFoodPreferences`: BasicMutationResult
- `recipeDeleteMemberRecipe`: DeleteMemberRecipeMutationResult
- `recipeDeleteShoppableRecipe`: BasicMutationResult
- `recipeSaveFoodPreferencesV2`: BasicMutationResult
- `recipeSaveMemberRecipe`: SaveMemberRecipeMutationResult
- `recipeSetMemberRatingV2`: RecipeSetMemberRatingResult
- `recipeShoppableSaveSelectionV2`: RecipeShoppableSaveSelectionResult
- `recipeUpdateWeekPlannerItems`: RecipeUpdateWeekPlannerItemsResult
- `settlementCancel`: BasicMutationResult
- `stampSharingGiftLinkCreate`: StampShareableMutationResult
- `stampSharingGiftLinkDelete`: StampShareableMutationResult
- `stampSharingRequestTransfer`: StampTransferMutationResult
- `stampSharingGiftTransfer`: StampTransferMutationResult
- `storesSetFavouriteStore`: BasicMutationResult
- `subscriptionBundleSubscribe`: SubscriptionBundleSubscribeMutationResult
- `subscriptionSetFixedDeliverySlot`: SubscriptionSetFixedDeliverySlotMutationResult
- `subscriptionCancel`: SubscriptionCancelMutationResult
- `subscriptionDefinitionChange`: SubscriptionDefinitionChangeMutationResult
- `subscriptionBundlePromoCodeApply`: SubscriptionBundlePromoCodeApplyResult
- `subscriptionCurrentPromoCodeApply`: SubscriptionCurrentPromoCodeApplyResult
- `subscriptionReactivate`: SubscriptionReactivateResult
- `subscriptionSavingsCarouselUpdate`: SubscriptionSavingsCarouselUpdateMutationResult
- `testSuccess`: TestMutationResult
- `testFailed`: TestMutationResult
- `testPartial`: TestMutationResult
- `testError`: TestMutationResult
- `testErrorThrown`: TestMutationResult
- `testSuppressed`: TestMutationResult
- `testWithArguments`: TestMutationResult
- `targetedOfferAllocate`: TargetedOfferAllocationMutationResult
- `productReturn`: CreateProductReturnResult
- `productComplaintOnlineOrder`: CreateProductComplaintResult
- `productComplaintStorePurchase`: CreateProductComplaintResult
- `depositComplaint`: CreateDepositComplaintResult
- `messageCenterMarkMessage`: BasicMutationResult
- `messageCenterUpdateIsDeletedByUser`: MessageCenterUpdateMessageMutationResult
- `entryPointDismiss`: EntryPointDismissalMutationResult
- `groceryListAdd`: GroceryListResult
- `pushNotificationsTokenUpdate`: BasicMutationResult
- `communicationConsentDeviceNotificationSettingsUpdate`: BasicMutationResult
- `communicationConsentAdd`: BasicMutationResult
- `communicationConsentDelete`: BasicMutationResult
- `communicationConsentAddAll`: BasicMutationResult
- `communicationConsentCookieConsentUpdate`: BasicMutationResult
- `lifestyleCheckSyncIntakeAnswers`: IntakeMutationResult
- `lifestyleCheckSyncChapterAnswers`: ChapterMutationResult
- `lifestyleCheckSyncAdvice`: AdviceMutationResult
- `lifestyleCheckSyncAdviceV2`: AdviceMutationResult
- `lifestyleAddOptIn`: Boolean
- `lifestyleRemoveOptIn`: Boolean
- `posReceiptsDelete`: PosReceiptMutationResult
- `productPurchaseHistoryHideProducts`: ProductPurchaseHistoryResult

### Note
A note on a grocery list

- `id`: String
- `description`: String
- `position`: Int
- `crossedOff`: Boolean

### NotificationCMS
CMS Notification

- `id`: String
- `type`: NotificationType
- `link`: NotificationLink
- `classification`: String
- `isGlobalNotification`: Boolean
- `priority`: NotificationPriorityType
- `priorityLabel`: String

### NotificationChannel
This type represents a communication topic. A topic is a specific type of message (e.g. track and trace) that is received
via a specific channel that a user can be subscribed to.

- `id`: String
- `name`: String
- `importance`: Int
- `allowedPushTypes`: 

### NotificationChannelsResponse
This type represents the response for the availablePushNotificationChannels

- `channels`: 

### NotificationLink
Notification link with href and title

- `href`: String
- `title`: String

### Order
Order

- `checkoutOverview`: CheckoutOverview
- `submitValidation`: CheckoutOrderValidation
- `delivery`: Delivery
- `id`: Int
- `state`: OrderState
- `price`: OrderPrice
- `isTransactionComplete`: Boolean
- `isExpired`: Boolean
- `entryEndDisplayTime`: String
- `reopened`: Int
- `orderMethod`: Int
- `receipt`: OrderReceipt
- `orderLines`: 
- `minimumValue`: Money
- `closingDateTime`: String
- `closingDateTimeStamp`: DateTime
- `lastUserChangeTime`: String
- `shoppingType`: OrderShoppingType
- `submitted`: Boolean
- `productsClosingTime`: String
- `messageForRecipient`: String
- `transactionsV2`: OrderTransactions

### OrderBudgetV2MutationResult
Result of performing a mutation.

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: CurrentBudget

### OrderBudgetV2QueryResult
Result of performing a query.

- `result`: CurrentBudget
- `errorMessage`: String

### OrderCalculationReceipt
Calculation Receipt of an Order

- `totalPrice`: Money
- `paymentAmount`: Money
- `groceriesQuantity`: Int
- `additionalItems`: 
- `totalAdditionalItems`: Money
- `additionalInfo`: 
- `walletOptions`: CheckoutWalletOptions
- `prepayments`: CheckoutPrepayments
- `wallet`: CheckoutWallet

### OrderCheckinMutationResult
Order checkin mutation result.

- `orderId`: Int
- `status`: MutationResultStatus
- `errorMessage`: String

### OrderLineV2
Order lines

- `product`: Product
- `quantity`: Int
- `originCode`: String
- `allocatedQuantity`: Int
- `closingTime`: String

### OrderPrice
Price of Order

- `priceBeforeDiscount`: Money
- `priceAfterDiscount`: Money
- `priceDiscount`: Money
- `priceTotalPayable`: Money

### OrderReceipt
Receipt of the order

- `netPrice`: Money
- `totalBonusDiscount`: Money
- `deposit`: Money
- `serviceCosts`: 
- `discounts`: 
- `subTotal`: Money
- `total`: Money
- `personalDiscount`: Money

### OrderReceiptModifier
Service cost or discount modifier

- `price`: Money
- `title`: String

### OrderReport
Summary of the total products ordered

- `products`: 

### OrderReportDiscount
Discount breakdown for order report

- `total`: Money
- `terra`: Money
- `bio`: Money
- `other`: Money
- `currency`: String

### OrderReportProduct
Product information in the order report

- `id`: Int
- `description`: String
- `brand`: String
- `salesUnitSize`: String
- `quantity`: Int
- `taxonomy`: OrderReportProductTaxonomy
- `totalCosts`: OrderReportProductCosts

### OrderReportProductCosts
Total costs for a product in the order report

- `amount`: Money
- `discount`: Money
- `deposit`: Money
- `vatRate`: Float
- `currency`: String

### OrderReportProductTaxonomy
Product taxonomy information

- `name`: String

### OrderReportTotal
Summary report of the total costs of the orders

- `quantityOfOrders`: Int
- `totalCosts`: OrderReportTotalCosts

### OrderReportTotalCosts
Total costs breakdown for order report

- `amount`: Money
- `discount`: OrderReportDiscount
- `serviceCharge`: Money

### OrderResendResponse
Order resend response

- `orderStatus`: String
- `externalOrderId`: String

### OrderSample
Sample products that belong to an order.

- `id`: Int
- `description`: String
- `salesUnitSize`: String
- `rejectAllowed`: Boolean
- `rejectSample`: Boolean
- `premium`: Boolean
- `product`: Product

### OrderTransactions
Returns transactions for a given order

- `hasSettlementTransaction`: Boolean
- `mandateInfo`: PaymentsPersonalAssetsAssets
- `primaryPaymentMethod`: PaymentMethod
- `transaction`: TransactionDetails

### OrderValueLimts
Order value limits

- `id`: Int
- `minimumOrderValue`: MinimumOrderValue
- `maximumOrderValue`: MaximumOrderValue
- `submittable`: Boolean

### OrderWalletOptions
Wallet options

- `giftCard`: 
- `purchaseStamps`: 
- `hasSettlementTransaction`: Boolean

### OriginalScannedIngredient
ScannedIngredient

- `name`: String
- `quantity`: Float
- `quantityUnit`: String
- `index`: Int

### PageInfo
We will show the total amount of items available, the offset its started in and the amount of items given back.

- `total`: Int
- `hasNextPage`: Boolean

### Pagination
Keep track of the pages of given items

- `totalSize`: Int
- `offset`: Int
- `pageSize`: Int

### PantryPromotionLabel
Promotion label props to pass to the PromotionLabel pantry component

- `emphasis`: PromotionLabelEmphasis
- `variant`: PromotionLabelVariant
- `title`: String
- `topText`: String
- `centerText`: String
- `bottomText`: String

### PantryTheme
Pantry theme

- `pantryThemeId`: String
- `lightTheme`: PantryColorSet
- `highContrastTheme`: PantryColorSet

### PasswordRulesValidated
Contains the rules and result against user password

- `isValid`: Boolean
- `rules`: 

### PayPaymentRequestResult
Result of creating a payment request

- `errorMessage`: String
- `status`: MutationResultStatus
- `request`: PaymentRequest
- `payment`: PaymentProcess

### PaymentAmount
Type that defines currency and the amount in money

- `amount`: Money
- `currency`: PaymentTransactionCurrencies

### PaymentIssuer
Available Payment Issuer

- `id`: Int
- `issuerLogoId`: String
- `name`: String
- `image`: String
- `url`: String
- `recentlyUsed`: Boolean
- `availability`: PaymentIssuerAvailability

### PaymentMutation
A change that ocurred to a payment

- `mutationId`: String
- `status`: PaymentMutationStatus
- `operations`: 

### PaymentMutationError
Payment Mutation Error

- `success`: Boolean
- `errorCode`: String
- `errorMessage`: String

### PaymentMutationErrorV5
Payment Mutation Error

- `type`: String
- `status`: Int
- `message`: String
- `title`: String
- `error`: String
- `instance`: String

### PaymentMutationOperation
Operations that were applied on an PaymentMutation

- `transactionId`: String
- `operationId`: String
- `operationStatus`: PaymentOperationStatus
- `operationType`: PaymentOperationType
- `error`: PaymentMutationError

### PaymentMutationOperationV5
Operations that were applied on an PaymentMutation

- `transactionId`: String
- `operationId`: String
- `operationStatus`: PaymentOperationStatus
- `operationType`: PaymentOperationType
- `error`: PaymentMutationErrorV5

### PaymentMutationV5
A change that ocurred to a payment

- `mutationId`: String
- `status`: PaymentMutationStatus
- `operations`: 

### PaymentOption
Payment option that the user can use

- `option`: PaymentMethod
- `issuers`: 
- `tokens`: 
- `status`: PaymentOptionStatus

### PaymentProcess
Ongoing Payment

- `id`: String
- `mutation`: PaymentMutation
- `action`: PaymentProcessingAction

### PaymentProcessingDeviceChallengeAction
Action that requires device input. Processed within the payment flow

- `transactionId`: String
- `authorizationId`: String
- `data`: String
- `sdk`: PaymentProcessingSdk

### PaymentProcessingPollAction
Action that requires the consumer to poll

- `interval`: Int
- `mutationId`: String

### PaymentProcessingRedirectAction
Action that requires leaving the payment flow.

- `redirectUri`: String

### PaymentProcessingSdk
SDK required for the Challenge

- `sdkType`: PaymentProcessingSdkType
- `minVersion`: String

### PaymentProcessingUserChallengeAction
Action that requires user input. Processed within the payment flow

- `transactionId`: String
- `authorizationId`: String
- `data`: String
- `sdk`: PaymentProcessingSdk

### PaymentRequest
Payment request

- `id`: String
- `title`: String
- `status`: PaymentRequestStatus
- `requestedAmount`: PaymentAmount
- `availablePaymentOptions`: 

### PaymentSDK
Payment SDK that needs to be consumed

- `id`: PaymentSDKType
- `minVersion`: String

### PaymentToken
Registered token

- `id`: String
- `cardAlias`: String
- `status`: PaymentTokenStatus
- `default`: Boolean
- `issuerId`: String
- `cardArtId`: String
- `psp`: PaymentTokenPspProvider
- `createdDate`: DateTime

### PaymentsAddGiftCardResult
Result of adding a GiftCard

- `errorMessage`: String
- `status`: MutationResultStatus
- `onboardingResponseCode`: PaymentsTokenOnboardingResponseCode
- `card`: PaymentsGiftCard

### PaymentsAuthorizeResult
Payments Authorize result

- `status`: MutationResultStatus
- `errorMessage`: String
- `action`: CheckoutOrderSubmissionPayment

### PaymentsBancontactMobileTransaction
Bancontact Mobile Payment Transaction

- `issuerId`: String
- `createdDate`: DateTime

### PaymentsCreatePaymentRequestPaymentResult
Result of the payment request

- `errorMessage`: String
- `status`: MutationResultStatus
- `payment`: PaymentProcess

### PaymentsDCTCard
DCT Card Type

- `cardId`: String
- `cardAlias`: String
- `default`: Boolean
- `issuerId`: String
- `cardArtId`: String
- `createdDate`: DateTime
- `status`: PaymentTokenStatus

### PaymentsEnrollPersonalAssetResult
Result of enrolling personal asset

- `errorMessage`: String
- `status`: MutationResultStatus
- `id`: String
- `asset`: PaymentsPersonalAssetsAssets

### PaymentsGiftCard
Gift Card such as NPL

- `cardId`: String
- `balance`: PaymentAmount
- `cardNumber`: String
- `status`: PaymentTokenStatus
- `createdDate`: DateTime

### PaymentsIDEALTransaction
iDEAL Payment Transaction

- `issuerId`: String
- `createdDate`: DateTime

### PaymentsOnboardDCTTokenResult
Result of onboarding DCT Token

- `errorMessage`: String
- `status`: MutationResultStatus
- `card`: PaymentsDCTCard
- `onboardingResult`: DCTOnboardingResult

### PaymentsOrderTransactions
Return type for applied transactions on a order

- `dct`: 
- `iDEAL`: PaymentsIDEALTransaction
- `giftCard`: 
- `purchaseStamps`: 
- `bancontactMobile`: PaymentsBancontactMobileTransaction
- `hasSettlementTransaction`: Boolean
- `sepaDirectDebit`: PaymentsSepaDirectDebitTransaction

### PaymentsPersonalAsset
Payment Option for a customer with included assets or available potential assets

- `method`: PaymentMethod
- `issuers`: 
- `tokens`: 
- `enroll`: 
- `assets`: 
- `accessPermissionState`: PaymentsAccessPermissionState
- `availability`: PaymentIssuerAvailability

### PaymentsPersonalAssetIssuer
Issuer for Personal Asset

- `bic`: String
- `name`: String
- `defaultIssuer`: Boolean
- `availability`: PaymentIssuerAvailability
- `enrollUrl`: String

### PaymentsPersonalAssetToken
Token for personal assets

- `id`: String
- `cardAlias`: String
- `status`: PaymentTokenStatus
- `defaultToken`: Boolean
- `bic`: String
- `createdDate`: Date

### PaymentsPersonalAssetsAssets
Mandate for personal assets

- `name`: String
- `status`: PaymentsPersonalAssetsStatus
- `externalReference`: String
- `mandateSource`: String
- `iban`: String
- `accountHolderName`: String
- `bic`: String
- `createdDate`: String
- `id`: String
- `assetType`: PaymentsPersonalAssetsAssetType
- `lastModifiedDate`: String
- `lastUsedDate`: String

### PaymentsPersonalAssetsEnroll
Enroll personal assets

- `createdDate`: String
- `referenceId`: String
- `iban`: String
- `bic`: String
- `enrollType`: PaymentsPersonalAssetsEnrollType
- `url`: String

### PaymentsPurchaseStampTransaction
Purchase stamp account

- `bookletsAmount`: Int

### PaymentsQRAction
User is shown a qr code to complete action

- `qrCodeData`: String
- `pspData`: String
- `authorizationId`: String

### PaymentsSepaDirectDebitTransaction
Sepa Direct Debit Payment Transaction

- `createdDate`: DateTime

### PaymentsUpdateCardResult
Payments Update card result

- `errorMessage`: String
- `errorCode`: UpdateCardResultCode
- `status`: MutationResultStatus
- `card`: PaymentsOrderCards

### PdfResponse
PdfResponse

- `pdfBase64`: String

### PersonalPromotionBundles
Personal promotion bundles

- `maximumActivations`: Int
- `validityPeriod`: PersonalPromotionBundlesValidityPeriod
- `error`: PersonalPromotionErrorMessage

### PersonalPromotionBundlesValidityPeriod
Personal promotion bundle validity period

- `start`: String
- `end`: String

### PhoneChannel
Customer care phone channel

- `waitingTimeMinutes`: Int
- `employeeAvailability`: ChannelAvailabilityForDay
- `id`: String
- `type`: ChannelType
- `visibleOnPlatforms`: 
- `availability`: ChannelAvailabilityForDay
- `phoneNumber`: String
- `phoneNumberDisplay`: String
- `chargePerMinute`: Int
- `label`: String
- `hasIncident`: Boolean

### PickupLocation
Pickup location

- `id`: Int
- `name`: String
- `description`: String
- `deliveryLocationId`: Int
- `type`: PickupLocationType
- `address`: StoresAddress
- `geoLocation`: GeoLocation
- `openingHours`: 

### PickupLocationEnableResult
Result of enabling a pickup location

- `status`: MutationResultStatus
- `errorMessage`: String
- `statusCode`: Int

### PickupLocationOpeningHours
Pickup location opening hours date and time ranges

- `date`: String
- `blocks`: 

### PickupLocationOpeningHoursBlock
Pickup location opening hours time range

- `from`: String
- `to`: String

### PickupLocationType
Pickup location type

- `code`: PickupLocationTypeCode
- `description`: String

### PollAction
Poll payment status may return a suggested poll policy

- `mutationId`: String
- `pollInterval`: Int

### PosReceipt
PosReceipt

- `id`: String
- `memberId`: String
- `bonusCardId`: String
- `storeInfo`: String
- `verifications`: PosReceiptVerification
- `products`: PosReceiptProduct
- `loyaltyCards`: PosReceiptLoyaltyCard
- `coupons`: String
- `subtotalProducts`: PosReceiptSubtotal
- `footnotes`: String
- `discounts`: PosReceiptDiscount
- `discountTotal`: Money
- `discountPersonal`: PosReceiptPersonalDiscount
- `subtotalDiscount`: Money
- `stamps`: PosReceiptStamps
- `total`: Money
- `payments`: PosReceiptPayment
- `change`: Money
- `giftCards`: PosReceiptGiftCards
- `eftReceipt`: String
- `vat`: PosReceiptVat
- `promotions`: PosReceiptPromotion
- `transaction`: PosReceiptTransaction
- `storeAdditionalInfo`: String
- `errors`: String
- `address`: PosReceiptStoreAddress
- `storeType`: PosReceiptStoreType
- `memberships`: String
- `returnPayment`: PosReceiptReturnPayment

### PosReceiptActivation
PosReceiptActivation

- `id`: String
- `name`: String
- `price`: Money
- `status`: String

### PosReceiptBalance
PosReceiptBalance

- `id`: String
- `balance`: Money

### PosReceiptDiscount
PosReceiptDiscount

- `type`: String
- `name`: String
- `amount`: Money

### PosReceiptGiftCards
PosReceiptGiftCards

- `balance`: PosReceiptBalance
- `activations`: PosReceiptActivation

### PosReceiptIndicator
PosReceiptIndicator

- `name`: String
- `discount`: String
- `percentage`: Float

### PosReceiptListItem
A single pos receipt.

- `id`: String
- `dateTime`: DateTime
- `totalAmount`: Money
- `storeAddress`: PosReceiptStoreAddress

### PosReceiptLoyaltyCard
PosReceiptLoyaltyCard

- `id`: String
- `description`: String

### PosReceiptMutationResult
Result after deleting Pos Receipt(s)

- `status`: MutationResultStatus
- `errorMessage`: String
- `receiptIds`: String

### PosReceiptPackageTax
PosReceiptPackageTax

- `name`: String
- `amount`: Float
- `currency`: String

### PosReceiptPage
A page containing Pos Receipts. This type also has a pagination object attached to it
with information on the pagination object that was passed along.

- `pagination`: PosReceiptPagination
- `posReceipts`: 

### PosReceiptPagination
A pagination object containing information on the page of receipts that was
returned.

- `key`: String
- `limit`: Int
- `offset`: Int
- `totalElements`: Int

### PosReceiptPayment
PosReceiptPayment

- `method`: String
- `amount`: Money
- `products`: String

### PosReceiptPersonalDiscount
PosReceiptPersonalDiscount

- `description`: String
- `amount`: Money

### PosReceiptProduct
PosReceiptProduct

- `id`: Int
- `quantity`: Float
- `weight`: PosReceiptWeight
- `name`: String
- `externalIds`: String
- `price`: Money
- `amount`: Money
- `deposit`: Money
- `indicators`: PosReceiptIndicator
- `packageTax`: PosReceiptPackageTax

### PosReceiptPromotion
PosReceiptPromotion

- `quantity`: Int
- `name`: String

### PosReceiptReturnPayment
PosReceiptReturnPayment

- `method`: String
- `amount`: Money
- `footnote`: String

### PosReceiptStamps
PosReceiptStamps

- `quantity`: Int
- `amount`: Money

### PosReceiptStoreAddress
Address of the store where the receipt is issued.

- `street`: String
- `houseNumber`: String
- `postalCode`: String
- `city`: String
- `countryCode`: String

### PosReceiptSubtotal
PosReceiptSubtotal

- `quantity`: Int
- `amount`: Money

### PosReceiptTotal
PosReceiptTotal

- `salesAmount`: Money
- `amount`: Money

### PosReceiptTransaction
PosReceiptTransaction

- `store`: Int
- `lane`: Int
- `id`: Int
- `operator`: Int
- `dateTime`: String

### PosReceiptVat
PosReceiptVat

- `levels`: 
- `total`: PosReceiptTotal

### PosReceiptVatLevel
PosReceiptVatLevel

- `percentage`: Int
- `salesAmount`: Money
- `amount`: Money

### PosReceiptVerification
PosReceiptVerification

- `id`: Int
- `quantity`: Int
- `weight`: PosReceiptWeight
- `name`: String
- `amount`: Money
- `discountPercentage`: Int
- `description`: String

### PosReceiptWeight
PosReceiptWeight

- `amount`: Float
- `unit`: String

### PrePayment
Pre payments

- `paymentMethod`: PaymentMethod
- `createdDate`: DateTime
- `tokenId`: String

### PreviouslyBoughtBargainItem

- `product`: Product
- `markdown`: BargainMarkdown
- `stock`: Int
- `bargainPrice`: BargainPrice

### PriceTotal
Contains the total price of the products along with the total discount.

- `withoutDiscount`: Money
- `discount`: Money

### Product
A product is a thing that can be ordered.

- `priceV2`: ProductPriceV2
- `id`: Int
- `additionalInformation`: String
- `ageCheck`: Boolean
- `availability`: ProductAvailability
- `brand`: String
- `category`: String
- `externalWebshopUrl`: String
- `highlight`: String
- `highlights`: String
- `hqId`: Int
- `icons`: 
- `imagePack`: 
- `interactionLabel`: String
- `isDeactivated`: Boolean
- `isMedicalDevice`: Boolean
- `isMedicine`: Boolean
- `isSample`: Boolean
- `listPrice`: ListPrice
- `minBestBeforeDays`: Int
- `nutritionalInfo`: ProductNutritionalInfo
- `otherSorts`: 
- `privateLabel`: Boolean
- `properties`: 
- `salesUnitSize`: String
- `shopType`: ProductShopType
- `summary`: String
- `taxonomies`: 
- `taxonomyPath`: 
- `title`: String
- `tradeItem`: ProductTradeItem
- `variant`: ProductVariant
- `variants`: 
- `virtualBundleProducts`: ProductVirtualBundleItem
- `virtualBundles`: Product
- `webPath`: String
- `isSponsored`: Boolean

### ProductAvailability
product availability

- `productId`: Int
- `isOrderable`: Boolean
- `isVisible`: Boolean
- `online`: ProductAvailabilityIndication
- `offline`: ProductAvailabilityIndication
- `unavailableForOrder`: ProductUnavailableForOrderIndication
- `availabilityLabel`: String
- `maxUnits`: Int
- `storeSpecificAssortmentStatus`: StoreSpecificAssortmentStatus
- `storeAssortmentStatus`: StoreAssortmentAvailabilityStatus

### ProductAvailabilityIndication
indication of the availability of a product

- `status`: ProductAvailabilityStatus
- `availableFrom`: String

### ProductBrand
A product brand

- `title`: String
- `name`: String
- `count`: Int

### ProductCardAvailability
product availability

- `productId`: Int
- `isOrderable`: Boolean
- `isVisible`: Boolean
- `online`: ProductCardAvailabilityIndication
- `offline`: ProductCardAvailabilityIndication
- `availabilityLabel`: String
- `maxUnits`: Int

### ProductCardAvailabilityIndication
indication of the availability of a product

- `status`: ProductCardAvailabilityStatus
- `availableFrom`: String

### ProductCardBonusSegmentAvailability
Bonus Segment availability

- `startDate`: String
- `endDate`: String
- `description`: String

### ProductCardImage
Each product is expected to have 1 or more image(s)/angle(s) associated with it.

- `angle`: ProductCardImageAngle
- `small`: ProductCardImageRendition
- `medium`: ProductCardImageRendition
- `large`: ProductCardImageRendition

### ProductCardImageRendition
The minimal description of a product image that includes the URL, width and height

- `width`: Int
- `height`: Int
- `url`: String

### ProductCardPriceDiscountV2
Price discount information

- `segmentId`: Int
- `description`: String
- `type`: String
- `promotionType`: ProductCardBonusPromotionType
- `segmentType`: ProductCardBonusSegmentType
- `subtitle`: String
- `theme`: ProductCardBonusTheme
- `tieredOffer`: 
- `availability`: ProductCardBonusSegmentAvailability
- `wasPriceVisible`: Boolean
- `productCount`: Int
- `multipleItemPromotion`: Boolean
- `smartLabel`: String

### ProductCardPriceV2
Product card price V2 type

- `now`: Float
- `was`: Float
- `discount`: ProductCardPriceDiscountV2
- `promotionLabel`: ProductCardPromotionLabel

### ProductCardPriceV2Original
Product card price V2 type

- `now`: Money
- `was`: Money
- `discount`: ProductCardPriceDiscountV2
- `promotionLabel`: ProductCardPromotionLabel

### ProductCardPromotionLabel
Promotion label, which might contain information about the discount or other labels (ex: Nieuw, virtual bundle without a discount)

- `type`: ProductCardPromotionLabelType
- `tiers`: ProductCardPromotionLabelTier

### ProductCardPromotionLabelTier
A promotion label tier

- `mechanism`: BonusSegmentDiscountLabelCodes
- `description`: String
- `alternateDescription`: String
- `originalDescription`: String
- `amount`: Float
- `incentiveType`: String
- `count`: Int
- `freeCount`: Int
- `actualCount`: Int
- `price`: Float
- `percentage`: Float
- `unit`: String

### ProductCardVariant
Variant (e.g. wine by the case vs wine per bottle)

- `type`: ProductCardVariantType
- `label`: String
- `product`: SearchProductCard

### ProductCardVirtualBundleItem
Product virtual bundle item

- `quantity`: Int
- `productId`: Int

### ProductComplaintOrder
Product complaint order

- `id`: String
- `shoppingType`: ProductComplaintShoppingType
- `transactionDateTime`: String
- `deliveryStartTime`: String
- `deliveryEndTime`: String
- `transactionAddress`: ProductComplaintOrderAddress
- `claimStatus`: ProductComplaintClaimStatus

### ProductComplaintOrderAddress
Product complaint order address

- `street`: String
- `houseNumber`: String
- `postalCode`: String
- `city`: String

### ProductComplaintOrderProduct
Product complaint order product

- `id`: Int
- `product`: Product
- `description`: String
- `quantities`: Quantities
- `explanations`: 
- `submittableFeedbackTypes`: 
- `isClaimable`: Boolean

### ProductComplaintStoreProduct
Product complaint store product

- `id`: Int
- `hqId`: Int
- `description`: String
- `quantity`: Int
- `product`: Product
- `submittableFeedbackTypes`: 

### ProductComplaintStorePurchase
Product complaint store purchase

- `id`: String
- `transactionDateTime`: String
- `dateTimeFormatted`: String
- `transactionAddress`: ProductComplaintOrderAddress
- `claimStatus`: ProductComplaintClaimStatus

### ProductForIngredient
Product that was found for the scanned ingredient

- `product`: Product
- `quantityForServings`: Int
- `optional`: Boolean
- `ingredient`: OriginalScannedIngredient
- `alternativeSections`: 

### ProductImage
Each product is expected to have 1 or more image(s)/angle(s) associated with it.

- `angle`: ProductImageAngle
- `small`: ProductImageRendition
- `medium`: ProductImageRendition
- `large`: ProductImageRendition

### ProductImageMulti
Each product is expected to have 1 or more images associated with it.

- `angle`: ProductImageAngle
- `small`: ProductImageRendition
- `medium`: ProductImageRendition
- `large`: ProductImageRendition

### ProductImageRendition
The minimal description of a product image that includes the URL, width and height

- `width`: Int
- `height`: Int
- `url`: String

### ProductImageSet
For each product we have some images

- `productId`: Int
- `d24x24`: String
- `d80x80`: String
- `d200x200`: String
- `d708x708`: String

### ProductNutritionalInfo
Nutritional information for a product

- `amount`: String
- `unit`: String
- `type`: ProductNutrientType

### ProductPrice
A product price

- `now`: Money
- `was`: Money
- `unitInfo`: ProductUnitInfo
- `discount`: ProductPriceDiscount

### ProductPriceDiscount
Price discount information

- `segmentId`: Int
- `description`: String
- `type`: String

### ProductPriceDiscountV2
Price discount information

- `segmentId`: Int
- `description`: String
- `type`: String
- `promotionType`: BonusPromotionType
- `segmentType`: BonusSegmentType
- `subtitle`: String
- `theme`: BonusTheme
- `tieredOffer`: 
- `availability`: BonusSegmentAvailability
- `wasPriceVisible`: Boolean
- `productCount`: Int
- `multipleItemPromotion`: Boolean
- `smartLabel`: String
- `activationStatus`: BonusSegmentActivationStatus
- `exceptionRule`: String

### ProductPriceV2
Product price with discount

- `now`: Money
- `was`: Money
- `unitInfo`: ProductUnitInfo
- `discount`: ProductPriceDiscountV2
- `promotionLabel`: PromotionLabel
- `promotionLabels`: PantryPromotionLabel

### ProductProperty
Each product can have properties

- `code`: String
- `values`: 

### ProductPurchaseHistory
Purchase history

- `total`: Int
- `products`: 

### ProductPurchaseHistoryResult
Result of removing products from the purchase history

- `status`: MutationResultStatus
- `errorMessage`: String
- `hiddenProducts`: Int

### ProductRecommendations
Product recommendations

- `model`: Model
- `products`: 

### ProductReturnExplanations
Explanations

- `type`: ProductReturnExplanationType
- `label`: String

### ProductReturnOrder
Order eligible to return products

- `id`: String
- `transactionDateTime`: String
- `deliveryStartTime`: String
- `deliveryEndTime`: String
- `transactionAddress`: ProductReturnOrderAddress
- `shoppingType`: ProductReturnShoppingType
- `returnStatus`: ProductReturnStatus

### ProductReturnOrderAddress
The adress of the delivery

- `street`: String
- `houseNumber`: String
- `postalCode`: String
- `city`: String

### ProductReturnProduct
product with return properties

- `id`: Int
- `product`: Product
- `description`: String
- `quantities`: ProductReturnProductQuantities
- `explanations`: 
- `isReturnable`: Boolean

### ProductReturnProductQuantities
Quantity properties for the return product

- `returnable`: Int
- `ordered`: Int

### ProductSearchResult
Result type for productSearch query

- `page`: ProductSearchResultPage
- `products`: 

### ProductSearchResultPage
Pagination details for productSearch

- `pageSize`: Int
- `pageNumber`: Int
- `totalPages`: Int
- `totalElements`: Int

### ProductTaxonomy
A product taxonomy

- `id`: Int
- `name`: String
- `slug`: String
- `active`: Boolean
- `parents`: 
- `children`: 
- `images`: TaxonomyImage
- `topLevelImage`: String
- `totalProductCount`: Int
- `coverProduct`: Product
- `products`: 
- `ageCheck`: Boolean

### ProductTaxonomyPath
A taxonomy path for a given product

- `id`: Int
- `name`: String
- `slugifiedName`: String
- `active`: Boolean

### ProductTradeItem
Trade item info

- `gln`: String
- `gtin`: String
- `gtinRevisions`: 
- `description`: ProductTradeItemDescription
- `contents`: ProductTradeItemContents
- `ingredients`: ProductTradeItemIngredients
- `nutritions`: 
- `feedingInstructions`: ProductTradeItemFeedingInstructions
- `usage`: ProductTradeItemUsage
- `storage`: ProductTradeItemStorage
- `origin`: ProductTradeItemOrigin
- `contact`: ProductTradeItemContact
- `additionalInfo`: ProductTradeItemAdditionalInfo
- `resources`: ProductTradeItemResources
- `marketing`: ProductTradeItemMarketing

### ProductTradeItemAdditionalInfo
Additional info

- `salesConditions`: 
- `identificationNumbers`: 
- `certificationNumbers`: 

### ProductTradeItemCommunicationChannel
Contact communication channel

- `type`: ProductTradeItemCommunicationChannelType
- `value`: String

### ProductTradeItemContact
Contact

- `name`: String
- `address`: String
- `communicationChannels`: ProductTradeItemCommunicationChannel

### ProductTradeItemContents
Contents data

- `netContents`: 
- `servingSize`: String
- `drainedWeight`: String
- `servingsPerPackage`: String
- `statement`: String
- `eMark`: Boolean

### ProductTradeItemDefinition
A generic product definition

- `value`: String
- `label`: String

### ProductTradeItemDescription
Generic product descriptions

- `descriptions`: 
- `definitions`: ProductTradeItemDescriptionDefinitions

### ProductTradeItemDescriptionDefinitions
Descriptive definitions, i.e. regarding alcohol or animal type

- `dosageForm`: String
- `percentageOfAlcohol`: String
- `sunProtectionFactor`: String
- `fishCatchInfo`: String
- `fishCatchMethod`: String
- `animalType`: String
- `animalFeedType`: String

### ProductTradeItemFeedingInstructions
Feedings instructions

- `statement`: String

### ProductTradeItemIdentificationNumber
Identification number

- `type`: String
- `label`: String
- `value`: String

### ProductTradeItemIngredientAllergens
Ingredient allergens data

- `list`: String
- `contains`: String
- `mayContain`: String
- `freeFrom`: String

### ProductTradeItemIngredients
Ingredients data

- `allergens`: ProductTradeItemIngredientAllergens
- `statement`: String
- `nonfoodIngredientStatement`: String
- `animalFeeding`: ProductTradeItemIngredientsAnimalFeeding

### ProductTradeItemIngredientsAnimalFeeding
Ingredients animal feeding data

- `statement`: String
- `analyticalConstituents`: String
- `additives`: String

### ProductTradeItemMarketing
Marketing data

- `features`: 
- `description`: String

### ProductTradeItemNutrient
Info about one specific nutrient

- `type`: String
- `name`: String
- `value`: String
- `superscript`: Int
- `dailyValue`: String

### ProductTradeItemNutrition
Nutrition data

- `dailyValueIntakeReference`: String
- `nutrients`: ProductTradeItemNutrient
- `servingSize`: String
- `servingSizeDescription`: String
- `preparationState`: String
- `additionalInfo`: ProductTradeItemDefinition
- `basisQuantity`: String
- `basisQuantityDescription`: String

### ProductTradeItemOrigin
Origin data

- `provenance`: 
- `activities`: ProductTradeItemOriginActivities

### ProductTradeItemOriginActivities
Origin activities

- `rearing`: 
- `birth`: 
- `slaughter`: 

### ProductTradeItemResourceAttachment
Resource attachment

- `name`: String
- `format`: String
- `type`: String
- `value`: String

### ProductTradeItemResourceIcon
Resource icon

- `type`: ProductTradeItemResourceIconType
- `id`: String
- `title`: String
- `meta`: ProductTradeItemResourceIconMeta

### ProductTradeItemResources
Resources

- `attachments`: ProductTradeItemResourceAttachment
- `icons`: ProductTradeItemResourceIcon

### ProductTradeItemStorage
Storage data

- `instructions`: 
- `lifeSpan`: 

### ProductTradeItemUsage
Instructions, suggestions and statements about product usage

- `instructions`: 
- `ageDescription`: String
- `servingSuggestion`: String
- `preparationInstructions`: 
- `dosageInstructions`: 
- `precautions`: String
- `warnings`: 
- `hazardStatements`: String
- `signalWords`: String
- `duringPregnancy`: String
- `duringBreastFeeding`: String
- `bacteriaWarning`: Boolean

### ProductUnavailableForOrderIndication
indication of the availability of a product

- `status`: ProductUnavailableReason
- `availableFrom`: String

### ProductUnitInfo
Info per unit

- `price`: Money
- `description`: String

### ProductVariant
Variant (e.g. wine by the case vs wine per bottle)

- `type`: ProductVariantType
- `label`: String
- `product`: Product

### ProductVirtualBundleItem
Product virtual bundle item

- `quantity`: Int
- `product`: Product

### Promotion
Promotion

- `id`: String
- `hqId`: Int
- `title`: String
- `subtitle`: String
- `segmentType`: PromotionSegmentType
- `promotionType`: PromotionType
- `exampleText`: String
- `extraDescriptions`: 
- `images`: 
- `productCount`: Int
- `price`: PromotionPrice
- `category`: String
- `activationStatus`: PromotionActivationStatus
- `salesUnitSize`: String
- `periodStart`: String
- `periodEnd`: String
- `periodDescription`: String
- `storeOnly`: Boolean
- `webPath`: String
- `exceptionRule`: String
- `redemptionCount`: Int
- `promotionLabels`: 
- `rawPromotionLabels`: 
- `rawPromotionLabelType`: PromotionLabelType
- `legacyBonusMechanism`: String
- `productIds`: 
- `products`: 
- `singleProductPromotion`: Boolean
- `product`: Product

### PromotionImage
PromotionImage

- `url`: String
- `title`: String
- `height`: Int
- `width`: Int

### PromotionLabel
Promotion label, which might contain information about the discount or other labels (ex: virtual bundle without a discount)

- `type`: PromotionLabelType
- `tiers`: PromotionLabelTier

### PromotionLabelTier
A promotion label tier

- `mechanism`: BonusSegmentDiscountLabelCodes
- `description`: String
- `alternateDescription`: String
- `originalDescription`: String
- `amount`: Float
- `incentiveType`: String
- `count`: Int
- `freeCount`: Int
- `actualCount`: Int
- `price`: Float
- `percentage`: Float
- `unit`: String

### PromotionPrice
Promotion price

- `now`: Money
- `was`: Money
- `label`: String

### PurchaseStampBalance
A type containing information on a single Koopzegel balance.

- `points`: PurchaseStampBalancePoints
- `money`: PurchaseStampBalanceMoney
- `constants`: PurchaseStampConstants

### PurchaseStampBalanceMoney
A type that contains information regarding the monetary value of a Koopzegel balance.

- `invested`: Money
- `interest`: Money
- `payout`: Money

### PurchaseStampBalancePoints
A type containing information on the points that are embedded in a Koopzegel balance.

- `totalPoints`: Int
- `currentBookletPoints`: Int
- `fullBooklets`: Int

### PurchaseStampBookletTarget
A type containing a target to achieve on a Koopzegel booklet.

- `points`: Int
- `interest`: Money

### PurchaseStampBookletTargets
A relatively constant type containing the booklet target for Koopzegel booklets.

- `halfBooklet`: Int
- `fullBooklet`: Int

### PurchaseStampConstants
A relatively constant type containing the constants for Koopzegels, such as the targets
and the price of a single Koopzegel stamp.

- `price`: Money
- `partialBookletTarget`: PurchaseStampBookletTarget
- `fullBookletTarget`: PurchaseStampBookletTarget

### PurchaseStampMigration
The migration status of the currently authenticated member.

- `status`: PurchaseStampMigrationStatus

### PurchaseStampSavingGoal
A type representing a purchase stamp saving goal.

- `amount`: Money
- `name`: String

### PurchaseStampSavingGoalMutationResult
A type representing the mutation result of a purchase stamp saving goal.
Either the purchase stamp saving goal is created or deleted.
In case of deletion, it contains only the status or error state of the mutation.

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: PurchaseStampSavingGoal

### PurchaseStampSecret
A type containing a Koopzegel secret. These secrets are used to generate TOTP codes, which
are used in the Koopzegel redeem process.

- `id`: UUID
- `secret`: String

### PurchaseStampSecretMutationResult
A type containing a Koopzegel secret. These secrets are used to generate TOTP codes, which
are used in the Koopzegel redeem process.

- `status`: MutationResultStatus
- `errorMessage`: String
- `secret`: PurchaseStampSecret

### PurchaseStampSecretState
A type containing information if a secret is (still) valid. Secrets can become invalid after,
for example, password changes, usage or simply because they don't exist (anymore) remotely.

- `id`: UUID
- `valid`: Boolean

### PurchaseStampTransactionPage
A page containing Koopzegel transactions. This type also has a pagination object attached to it
with information on the pagination object that was passed along.

- `pagination`: LoyaltyPointsTransactionPagination
- `transactions`: 

### QRCodeAction
Whenever BE returns a QR Code string, QRCodeAction is used

- `qrCodeData`: String
- `pspData`: String
- `authorizationId`: String

### Quantities
Quantities

- `claimable`: Int

### Query

- `addressSearch`: AddressSearch
- `advertisementPersonalized`: 
- `advertisementNonPersonalized`: 
- `advertisementPersonalizedV2`: 
- `advertisementNonPersonalizedV2`: 
- `ahMemberCompanyRegistrationV2`: MemberCompanyRegistrationResult
- `ahMemberCompanyRegistrationSearchPagination`: MemberCompanyRegistrationSearchResult
- `ahManagedCompanyMemberDetails`: MemberCompanyDetails
- `basket`: Basket
- `bonusPersonalPromotionBundles`: 
- `bonusPeriods`: 
- `bonusCategories`: 
- `bonusPromotions`: 
- `checkoutValidateOrder`: CheckoutOrderValidation
- `ciamCaptchaConfiguration`: CaptchaConfiguration
- `ciamPasskeySettings`: CiamPasskeySettings
- `ciamPasswordResetValidateCode`: ValidationResponse
- `ciamPasswordValidation`: PasswordRulesValidated
- `member`: Member
- `memberIsOnboarded`: Boolean
- `memberLoginState`: MemberLoginState
- `memberGetNewsletterStatus`: GetNewsletterStatus
- `memberUnsubscribeGetBlockerReasons`: 
- `memberAddCardCheck`: AddMemberCard
- `memberHasAccount`: MemberHasAccount
- `memberPhoneNumberIsValid`: MemberPhoneNumberIsValid
- `applicationName`: String
- `products`: 
- `version`: String
- `milesBalance`: MilesBalance
- `milesDoRedirectToLmn`: RedirectToLmn
- `milesTransactions`: 
- `milesCharities`: 
- `contentCMSBrandHeader`: ContentBrandHeaderData
- `contentCMSComponents`: 
- `contentCMSFlexPage`: ContentFlexPage
- `contentCMSFlexPageFromWebUrl`: ContentFlexPage
- `contentCMSHomeVideoList`: ContentVideoList
- `contentCMSMobileComponent`: ContentBaseMobileCMSComponent
- `contentCMSMobileLoyaltyCampaign`: ContentLoyaltyCampaignDocument
- `contentCMSMobileNextBestActionCard`: ContentMobileNextBestActionCard
- `contentCMSMobilePageTemplate`: ContentMobilePageTemplate
- `contentCMSPage`: ContentPage
- `contentCMSTargetedDocument`: ContentBaseTargetedContentCMSDocument
- `contentCustomerAttributes`: ContentCustomerAttributes
- `contentDeliveryGrid`: ContentDeliveryGrid
- `contentFederatedDomains`: ContentFederatedDomains
- `contentFooterLinks`: ContentFooterLinks
- `contentMegaMenuLinks`: ContentMegaMenuLinks
- `contentOptIn`: ContentOptIn
- `contentProductSuggestions`: 
- `contentResourceBundle`: ContentResourceBundle
- `contentThemesConfiguration`: ContentThemesConfiguration
- `contentUsps`: ContentUspGroup
- `contentAllerhandeCMSPage`: ContentPage
- `contentAllerhandeFlexPage`: ContentAllerhandeFlexPage
- `conversation`: Conversation
- `conversationGetFileUrls`: ConversationFiles
- `conversationSummary`: ConversationSummary
- `cookBook`: CookBook
- `cookBookCollectedRecipe`: CookBookCollectedRecipe
- `cookBookLastAdded`: 
- `cookBookMember`: CookBook
- `cookBookMemberBlockedMembers`: 
- `cookBookMemberMessage`: CookBookMemberMessage
- `cookBookMemberMessages`: CookBookMemberMessagesResult
- `cookBookMemberMessagingAllowed`: Boolean
- `cookBookMemberProfile`: CookBookMemberProfile
- `cookBookMemberRecipe`: CookBookMemberRecipe
- `cookBookMemberUnreadCount`: Int
- `cookBookRecipe`: CookBookRecipe
- `cookBookRecipeLastAdded`: 
- `cookBookRecipeTopVisited`: 
- `cookBookRecipes`: CookBookRecipes
- `cookBookRecipesCollected`: CookBookRecipesCollected
- `cookBookRecipesCollectedMember`: CookBookRecipesCollectedMember
- `cookBookRecipesMember`: CookBookRecipesMember
- `cookBookSearch`: CookBookSearchResult
- `cookBookSearchRecipe`: CookBookSearchRecipeResult
- `cookBookTopVisited`: 
- `customerCareChannels`: 
- `customerCareEmployeeLoadOfToday`: CustomerCareEmployeeLoad
- `customerCarePersonalizedOrdersLane`: 
- `customerCareRecentlyVisitedStores`: 
- `customerCareSettlementsPage`: CustomerCareSettlementsPage
- `customerCareSpecialDays`: 
- `customerCareStoreByPurchaseId`: CustomerCareStore
- `customerServiceSearchAutocomplete`: CustomerServiceSuggestionResponse
- `customerServiceSearch`: CustomerServiceSearchResponse
- `customerServiceSearchFaqItems`: 
- `customerServiceSearchAccordion`: CustomerServiceSearchAccordionResponse
- `customer`: Customer
- `deliveryBundles`: 
- `deliveryBundleSubscription`: DeliveryBundleSubscription
- `deliveryBundleCalculateChangePrice`: DeliveryBundlePriceToChange
- `deliveryPrevious`: Delivery
- `feedbackOrders`: 
- `feedback`: Feedback
- `feedbackImageWithUrl`: 
- `feedbackOptionsV2`: 
- `favoriteListV2`: 
- `invoices`: Invoices
- `invoicesV2`: InvoicesV2
- `invoiceReports`: 
- `loyaltyChallenge`: LoyaltyChallenge
- `loyaltyChallenges`: 
- `loyaltyPointsBalances`: 
- `loyaltyPointsTransactions`: LoyaltyPointsTransactionPage
- `loyaltyProgram`: LoyaltyProgram
- `loyaltyPrograms`: 
- `messengerGetConversation`: MessengerConversation
- `messengerGetGroupedConversation`: MessengerGroupedConversation
- `messengerGetFileUrls`: MessengerFiles
- `notifications`: 
- `notificationsGetDeliveryNotification`: DeliveryNotification
- `order`: Order
- `orderBudgetV2`: OrderBudgetV2QueryResult
- `orderDeliverySlots`: 
- `orderExpandFulfillment`: ExpandOrder
- `orderExpandFulfillments`: ExpandOrder
- `orderFulfillment`: Fulfillment
- `orderFulfillments`: FulfillmentsResult
- `orderFulfillmentsByDateRange`: 
- `orderLeadTime`: 
- `orderMergeOptions`: 
- `orderMinimumValue`: Float
- `orderPickupSlots`: 
- `orderReceipt`: OrderReceipt
- `orderReport`: OrderReport
- `orderReportTotal`: OrderReportTotal
- `orderSamples`: 
- `orderValueLimits`: OrderValueLimts
- `paymentsOptions`: 
- `paymentsAuthorizationStatus`: CheckoutOrderSubmissionPayment
- `paymentsGetGiftCards`: 
- `paymentsGetDCTCards`: 
- `paymentsGetOrderTransactions`: PaymentsOrderTransactions
- `paymentsGetFullBookletsCount`: Int
- `paymentsPaymentRequestById`: PaymentRequest
- `paymentsGetPaymentRequestStatus`: PaymentProcess
- `pickupLocations`: 
- `pickupLocation`: PickupLocation
- `totalPrice`: PriceTotal
- `purchaseStampBalance`: PurchaseStampBalance
- `purchaseStampCheckSecretState`: PurchaseStampSecretState
- `purchaseStampSavingGoal`: PurchaseStampSavingGoal
- `purchaseStampServerTime`: DateTime
- `purchaseStampTransactions`: PurchaseStampTransactionPage
- `purchaseStampMigrationStatus`: PurchaseStampMigrationStatus
- `questionnaireForm`: QuestionnaireForm
- `memberRecipeProductSuggestions`: 
- `recipe`: Recipe
- `recipeAllerhandeMagazine`: AllerhandeMagazine
- `recipeAllerhandeMagazines`: 
- `recipeAutoSuggestions`: 
- `recipeBonusLane`: RecipeBonusLane
- `recipeCollectionCategories`: 
- `recipeCollectionCategory`: RecipeCollectionCategory
- `recipeHighlightedThemes`: RecipeHighlightedThemes
- `recipeLanes`: 
- `recipeMemberRating`: Int
- `recipeMemberRecipe`: MemberRecipe
- `recipePreferences`: MemberRecipePreferences
- `recipePreparationSection`: RecipePreparationSection
- `recipeProductSuggestionsV2`: 
- `recipeRecommendationLane`: RecipeRecommendationLane
- `recipeRecommendations`: 
- `recipeRecommendationsV2`: RecipeRecommendationResult
- `recipeSearchV2`: RecipeSearchResult
- `recipeShoppedRecipes`: 
- `recipeStories`: 
- `recipeSuggestionsForProductId`: 
- `recipeSummaries`: 
- `recipeSummary`: RecipeSummary
- `recipeTopRecipes`: 
- `recipeVariants`: 
- `recipeVideo`: RecipeVideo
- `recipeVideoAutoSuggestions`: 
- `recipeVideoSearch`: RecipeVideoSearchResult
- `recipeVideosRelatedToVideo`: 
- `recipeWeekPlannerHistoryOverview`: WeekPlannerOverview
- `recipeWeekPlannerOverview`: WeekPlannerOverview
- `recipeWeekPlannerSuggestions`: WeekPlannerRecipeSuggestion
- `recipesRelatedToRecipe`: 
- `recipesRelatedToVideo`: 
- `scrapeRecipe`: ScrapedRecipe
- `recommendationsMissingBonus`: 
- `recommendationsForgottenProducts`: 
- `recommendationsForgottenProductsV2`: 
- `settlementsTotal`: SettlementsTotal
- `stampSharingGiftLinks`: 
- `stampSharingRequestLink`: StampShareable
- `stampSharingRequestValidate`: StampShareableValidationResult
- `storesInformation`: Stores
- `storesSearch`: StoresResultType
- `storesCityCount`: 
- `storesGeoLocationSuggestions`: 
- `storesFavouriteStore`: Stores
- `subscriptionCurrent`: SubscriptionCurrent
- `subscriptionBundlesAvailableV2`: 
- `subscriptionPaymentIDealIssuers`: 
- `subscriptionPaymentStatus`: SubscriptionPollPaymentStatus
- `subscriptionFixedDeliverySlotsV2`: 
- `subscriptionCalculateSavings`: SubscriptionSavings
- `subscriptionCalculateChange`: SubscriptionChange
- `subscriptionPremiumSavingsV2`: SubscriptionPremiumSavings
- `subscriptionEstimatedPremiumSavings`: SubscriptionEstimatedPremiumSavings
- `subscriptionSummary`: SubscriptionSummary
- `subscriptionFixedDeliverySlotCurrent`: SubscriptionFixedDeliverySlotCurrent
- `subscriptionEligibleForB2B`: Boolean
- `subscriptionSavingPotential`: SubscriptionSavingPotential
- `subscriptionSavingsCarousel`: 
- `subscriptionAvailableDeliveryDaysV2`: SubscriptionAvailableDeliveryDays
- `subscriptionPaymentOptions`: 
- `suppliers`: 
- `testEcho`: String
- `testEchoService`: String
- `testHttpHeaders`: 
- `videoBbw`: VideoBbw
- `targetedOfferChallenge`: TargetedOfferChallenge
- `productReturnOrders`: 
- `productReturnProducts`: 
- `productComplaintOrders`: 
- `productComplaintOrderProducts`: 
- `productComplaintStorePurchases`: 
- `productComplaintStoreProducts`: 
- `depositComplaintOrders`: 
- `depositComplaintDepositItems`: 
- `messageCenterGetMessages`: 
- `messageCenterGetUnreadMessagesInfo`: MessageCenterUnreadMessagesInfo
- `entryPointComponent`: EntryPointComponent
- `groceryList`: GroceryListGetResult
- `folders`: 
- `orderRulesMaximumValue`: Float
- `companyAccountsManagedAccounts`: 
- `communicationConsents`: CommunicationConsentsResponse
- `communicationConsentCommunicationTopicsDescription`: CommunicationTopicsDescriptionResponse
- `communicationConsentAvailablePushNotificationChannels`: NotificationChannelsResponse
- `lifestyleCheckIntroduction`: LifestyleIntroduction
- `lifestyleCheckIntake`: LifestyleCheckIntake
- `lifestyleCheckChapter`: LifestyleCheckChapter
- `lifestyleCheckChaptersMeta`: 
- `lifestyleCheckScore`: LifestyleCheckScore
- `lifestyleCheckScoreV2`: LifestyleCheckScore
- `lifestyleAdvice`: LifestyleAdvice
- `lifestyleNavigation`: LifestyleCheckNavigationSuggestion
- `lifestyleHubInfo`: LifestyleHubInfo
- `lifestyleOptInState`: Boolean
- `evalueAccountBalance`: EvalueAccountBalance
- `evalueTransactionHistory`: EvalueTransactionHistory
- `posReceiptsPage`: PosReceiptPage
- `posReceiptDetails`: PosReceipt
- `posReceiptPdf`: PdfResponse
- `scanIngredientsMapIngredientsToProducts`: 
- `selfServiceContentBySearchTerm`: SelfServiceContentQueryResponse
- `selfServiceContentByPath`: SelfServiceContentPathResponse
- `bargainCategories`: 
- `bargainItems`: 
- `previouslyBoughtBargainItems`: 
- `_entities`: 
- `_service`: _Service
- `product`: Product
- `productConvertId`: Int
- `productOtherSorts`: Product
- `productAlternatives`: ProductRecommendations
- `productCrossSellsV2`: ProductRecommendations
- `productSearch`: ProductSearchResult
- `productBrands`: 
- `productTopLevelTaxonomies`: 
- `productTaxonomyByName`: ProductTaxonomy
- `productTaxonomyById`: ProductTaxonomy
- `productTaxonomiesByIds`: ProductTaxonomy
- `productPurchaseHistory`: ProductPurchaseHistory
- `productTaxonomyPath`: ProductTaxonomyPath
- `productAutosuggest`: 
- `experimentalSearchProducts`: ExperimentalSearchPayload
- `search`: 
- `searchTaxonomies`: SearchTaxonomy
- `searchTaxonomiesById`: SearchTaxonomy
- `searchFacetConfig`: 
- `searchProducts`: SearchPayload
- `searchProductsExperimental`: SearchProductCardPayload
- `searchPreviouslyBought`: SearchPayload
- `searchCategory`: SearchPayload
- `searchBrand`: SearchPayload
- `searchContextual`: SearchContextualPayload
- `sourceVersions`: String

### QuestionnaireCheckboxElement
Type representing a checkbox element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `options`: 
- `otherOption`: String
- `otherPlaceholder`: String
- `maxLength`: Int

### QuestionnaireForm
Type representing a questionnaire form

- `topic`: String
- `title`: String
- `elements`: 

### QuestionnaireParagraphElement
Type representing a paragraph element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String

### QuestionnaireRadioElement
Type representing a radio element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `options`: 
- `otherOption`: String
- `otherPlaceholder`: String
- `maxLength`: Int

### QuestionnaireRate10Element
Type representing a 10-point rating element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `low`: String
- `high`: String
- `required`: Boolean

### QuestionnaireRate5Element
Type representing a 5-point rating element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `low`: String
- `high`: String
- `required`: Boolean

### QuestionnaireTextAreaElement
Type representing a text area element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `maxLength`: Int
- `placeholder`: String

### QuestionnaireTextElement
Type representing a text element in a questionnaire

- `id`: ID
- `type`: String
- `text`: String
- `required`: Boolean
- `maxLength`: Int
- `placeholder`: String

### RawPromotionLabel
Unparsed promotion label data

- `mechanism`: PromotionLabelMechanism
- `defaultDescription`: String
- `price`: Float
- `percentage`: Float
- `amount`: Float
- `count`: Int
- `freeCount`: Int
- `actualCount`: Int
- `deliveryType`: String
- `unit`: String

### Recipe
We have two types of entities in recipe domain: Recipe and RecipeSummary
Recipe has more data compare to RecipeSummary and they have different data sources:
RecipeSummar comes from Elastic
Recipe comes from PostgreSQL DB

Will be extended with more fields

- `alternateTitle`: String
- `author`: RecipeAuthor
- `classifications`: 
- `cookTime`: Int
- `courses`: 
- `cuisines`: 
- `description`: String
- `flags`: 
- `group`: String
- `href`: String
- `id`: Int
- `images`: 
- `ingredients`: 
- `kitchenAppliances`: 
- `magazines`: 
- `meta`: RecipeMeta
- `modifiedAt`: String
- `noIndex`: Boolean
- `nutriScore`: RecipeNutriScore
- `nutritions`: RecipeNutritionInfo
- `ovenTime`: Int
- `preparation`: RecipePreparation
- `publishedAt`: String
- `rating`: RecipeRating
- `seoCanonical`: String
- `servings`: RecipeDetailsServing
- `slugifiedTitle`: String
- `spiciness`: Int
- `tags`: 
- `tips`: 
- `title`: String
- `variants`: 
- `videos`: RecipeVideos
- `waitTime`: Int

### RecipeAlternativeSection
Alternative Sections, alternative products are divided into sections, for example Biological, Product with less waste, Best priced, etc..

- `title`: String
- `description`: String
- `productSuggestions`: 

### RecipeAuthor
Recipe source

- `origin`: RecipeOrigin
- `brand`: RecipeBrand

### RecipeBonusLane
Bonus recipes list

- `bonusRecipes`: 

### RecipeBrand
Recipe brand

- `name`: String
- `logo`: ImageSet
- `supplier`: RecipeBrandSupplier

### RecipeBrandSupplier
Recipe brand supplier

- `name`: String

### RecipeCollectionAssignRecipeToCategoriesResult
Mutation result

- `status`: MutationResultStatus
- `result`: RecipeCollectionCategoryRecipe
- `errorMessage`: String

### RecipeCollectionCategory
Recipe collection - all member recipes
Collection is divided into categories
Each category contains recipes (it's only Allerhande)

- `id`: Int
- `name`: String
- `isDefault`: Boolean
- `list`: 
- `recipes`: 

### RecipeCollectionCategoryRecipe
Contains informaton about collected recipe
(Will be extended with type property in future)

- `id`: Int
- `type`: CollectionCategoryRecipeType

### RecipeCollectionCreateCategoryResult
Mutation result

- `status`: MutationResultStatus
- `result`: RecipeCollectionCategory
- `errorMessage`: String

### RecipeCollectionDeleteCategoryResult
Mutation result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### RecipeCollectionRemoveRecipeFromCategoriesResult
Mutation result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### RecipeCollectionUpdateCategoryResult
Mutation result

- `status`: MutationResultStatus
- `result`: RecipeCollectionCategory
- `errorMessage`: String

### RecipeCollectionUpdateRecipeWithCategoriesResult
Mutation result

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### RecipeDetailsServing
Recipe Serving Details

- `number`: Int
- `servingType`: RecipeServingType
- `min`: Int
- `max`: Int
- `scale`: Int
- `isChangeable`: Boolean

### RecipeFilter
Recipe filter

- `group`: String
- `values`: 

### RecipeHighlightedThemes
Highlighted recipe themes created in Hippo CMS by content team

- `title`: String
- `themes`: 

### RecipeImage
Extended model for recipe image

- `rendition`: RenditionKey
- `url`: String
- `width`: Int
- `height`: Int

### RecipeImageRenditions
Recipe Image Renditions

- `d220x162`: ImageSet
- `d302x220`: ImageSet
- `d440x324`: ImageSet
- `d445x297`: ImageSet
- `d612x450`: ImageSet
- `d680x320`: ImageSet
- `d890x594`: ImageSet
- `d1024x748`: ImageSet
- `d1224x900`: ImageSet
- `d1600x560`: ImageSet
- `d1920x1080`: ImageSet
- `d2048x1496`: ImageSet

### RecipeIngredient
Recipe ingredient type

- `id`: Int
- `name`: SingularPluralName
- `quantity`: Float
- `quantityUnit`: SingularPluralName
- `text`: String
- `servingsScale`: Float
- `curatedBonus`: Boolean

### RecipeKitchenAppliance
Recipe Kitchen Appliance

- `name`: String
- `quantity`: Int
- `scalable`: Boolean
- `text`: String

### RecipeLane
CMS configured recipe lane

- `id`: String
- `name`: String
- `recipeSummary`: 

### RecipeMagazine
Recipe magazine issue

- `title`: String
- `date`: String
- `number`: String
- `issueSlug`: String
- `type`: String

### RecipeMeta
Recipe seo meta

- `description`: String
- `title`: String

### RecipeNutrition
Nutrition data

- `name`: String
- `value`: Float
- `unit`: String

### RecipeNutritionInfo
Nutrition data set

- `energy`: RecipeNutrition
- `carbohydrates`: RecipeNutrition
- `sugar`: RecipeNutrition
- `sodium`: RecipeNutrition
- `protein`: RecipeNutrition
- `fat`: RecipeNutrition
- `saturatedFat`: RecipeNutrition
- `fibers`: RecipeNutrition

### RecipeOrigin
Recipe origin

- `type`: RecipeOriginType
- `hostName`: String
- `url`: String

### RecipePreparation
Recipe preparation details

- `steps`: 
- `summary`: 

### RecipePreparationSection
Recipe preparation section

- `recipeTitle`: String
- `recipeImages`: 
- `servingSize`: Int
- `cookingTime`: Int
- `preparationSteps`: 

### RecipePreparationStep
Preparation step information

- `text`: String
- `ingredients`: 
- `kitchenTimer`: Int
- `tips`: 
- `images`: 
- `videos`: 
- `kitchenAppliances`: 

### RecipeProductSuggestionsIngredient
Recipe Ingredient

- `id`: Int
- `name`: String
- `quantityFloat`: Float
- `quantityUnit`: String
- `index`: Int
- `additionalInfo`: String
- `completeText`: String
- `description`: SingularPluralName
- `rawIngredientText`: String

### RecipeRating
Defines recipe rating

- `count`: Int
- `average`: Float

### RecipeRecommendation
Recipe recommendation type

- `id`: Int
- `recipe`: RecipeSummary

### RecipeRecommendationLane
All necessary data to render recipe recommendation lane on koken tab.
Contains recommended recipes and member preferences

- `availableWeeks`: 
- `bonusRecipes`: 
- `genericRecipeRecommendations`: 
- `memberPreferencePreviewOptions`: 
- `preferenceTilePosition`: Int
- `recipeRecommendations`: 
- `selectedNutritionAndAllergyPreferences`: 
- `servingSize`: Int

### RecipeRecommendationResult
Recipe recommendation result

- `recipeRecommendations`: 
- `bonusRecipes`: 
- `genericRecipeRecommendations`: 
- `preferenceTilePosition`: Int

### RecipeSearchResult
Recipe search data

- `page`: PageInfo
- `filters`: 
- `result`: 
- `correctedSearchTerm`: String

### RecipeSearchResultFilter
Recipe search data filter

- `name`: String
- `label`: String
- `group`: String
- `count`: Int
- `selected`: Boolean

### RecipeSearchResultFilterGroup
Recipe search data filter group

- `label`: String
- `name`: String
- `filters`: 

### RecipeSetMemberRatingResult
Set member's rating for recipe mutation result type

- `status`: MutationResultStatus
- `result`: UpdatedRecipeRating
- `errorMessage`: String

### RecipeShoppableProductSuggestion
Recipe product suggestion

- `id`: Int
- `quantity`: Float
- `proposer`: RecipeShoppableProductSuggestionProposer
- `product`: Product

### RecipeShoppableSaveSelectionResult
Result type for the mutations that saves user selections of the ingredients to products mappings

- `status`: MutationResultStatus
- `result`: Boolean
- `errorMessage`: String

### RecipeShoppedRecipe
Recipe shopped recipe

- `recipe`: RecipeSummary
- `lastShoppedAt`: DateTime

### RecipeSource
Recipe source

- `url`: String
- `hostName`: String
- `type`: RecipeSourceType

### RecipeStoryActivePeriod
Defines when recipe story will be visible/active

- `from`: String
- `to`: String

### RecipeStoryCampaign
Recipe story a/b testing campaign

- `controlVariant`: Boolean
- `campaignId`: Int
- `variantId`: Int

### RecipeStoryCollection
Recipe Story

- `id`: Int
- `title`: String
- `theme`: RecipeStoryCollectionTheme
- `type`: RecipeStoryCollectionType
- `campaign`: RecipeStoryCampaign
- `active`: RecipeStoryActivePeriod
- `coverImage`: ImageSet
- `storyItems`: 

### RecipeStoryFlexPage
Recipe Story Flex Page

- `id`: Int
- `flexPageId`: String
- `title`: String
- `description`: String
- `heroImage`: ImageSet
- `videoId`: Int
- `video`: RecipeVideo

### RecipeStoryFlexPageStoryItem
Recipe Story flex page story item type

- `type`: RecipeStoryItemType
- `data`: RecipeStoryFlexPage

### RecipeStoryProductBundle
Recipe Story productBundle

- `id`: Int
- `title`: String
- `sponsored`: Boolean
- `products`: 
- `heroImage`: ImageSet
- `videoId`: Int
- `video`: RecipeVideo

### RecipeStoryProductBundleItem
Recipe Story productBundle Items

- `id`: Int
- `optional`: Boolean
- `product`: Product

### RecipeStoryProductBundleStoryItem
Recipe Story productBundle story item type

- `type`: RecipeStoryItemType
- `data`: RecipeStoryProductBundle

### RecipeStoryRecipe
Recipe Story Recipe

- `id`: Int
- `recipeId`: Int
- `recipe`: RecipeSummary
- `videoId`: Int
- `video`: RecipeVideo
- `sponsored`: Boolean

### RecipeStoryRecipeStoryItem
Recipe Story recipe story item type

- `type`: RecipeStoryItemType
- `data`: RecipeStoryRecipe

### RecipeSummary
We have two types of entities in recipe domain: Recipe and RecipeSummary
RecipeSummary has less data compare to Recipe and they have different data sources:
RecipeSummary comes from Elastic
Recipe comes from PostgreSQL DB

- `id`: Int
- `title`: String
- `slug`: String
- `alternateTitle`: String
- `diet`: 
- `courses`: 
- `publishedAt`: String
- `modifiedAt`: String
- `time`: RecipeTime
- `rating`: RecipeRating
- `images`: 
- `serving`: RecipeSummaryServing
- `nutrition`: RecipeNutritionInfo
- `nutriScore`: RecipeNutriScore
- `author`: RecipeAuthor
- `flags`: 
- `tags`: 
- `matchedIngredientsNumber`: Int

### RecipeSummaryServing
Recipe serving information

- `number`: Int
- `servingType`: RecipeServingType
- `min`: Int
- `max`: Int
- `scale`: Int
- `isChangeable`: Boolean

### RecipeSupplier
Company that provided the recipe

- `name`: String
- `logo`: 

### RecipeTag
Recipe tag

- `key`: String
- `value`: String

### RecipeTheme
Recipe theme represents a document with type recipeGroup in hippo (also called themes by the content team)

- `title`: String
- `headerRecipeId`: Int
- `headerRecipe`: RecipeSummary
- `highlightedRecipeId`: Int
- `highlightedRecipe`: RecipeSummary
- `themeSearchParams`: RecipeThemeSearchParams

### RecipeThemeSearchParams
Query to search the recipes related to this theme

- `searchText`: String
- `sortBy`: RecipeSearchSortOption
- `filters`: 
- `priorityRecipeIds`: 

### RecipeTime
Defines time to spend for each stage of cooking

- `cook`: Int
- `oven`: Int
- `wait`: Int

### RecipeTip
Recipe tip

- `type`: String
- `value`: String

### RecipeToProductSuggestionsResult
Map Product to Suggestion result

- `ingredient`: RecipeProductSuggestionsIngredient
- `optional`: Boolean
- `productSuggestion`: RecipeShoppableProductSuggestion
- `alternativeSections`: 
- `index`: Int

### RecipeTopRecipes
Top recipes list created in Hippo CMS by content team

- `id`: Int
- `title`: String
- `category`: String
- `overviewTitle`: String
- `images`: 
- `recipeSummary`: 
- `textColor`: String
- `useGradient`: Boolean

### RecipeUpdateWeekPlannerItemsResult
Mutation result

- `status`: MutationResultStatus
- `errorMessage`: String

### RecipeVideo
Recipe Video (from Blue Billy Wig)

- `id`: Int
- `title`: String
- `href`: String
- `slug`: String
- `description`: String
- `descriptionHtml`: String
- `category`: String
- `duration`: String
- `publication`: String
- `images`: RecipeVideoImages
- `views`: Int
- `streams`: RecipeVideoStreams

### RecipeVideoImages
Recipe Video Images

- `sd`: String
- `thumbnail`: String

### RecipeVideoSearchResult
Result of searching recipe videos

- `page`: PageInfo
- `filters`: 
- `result`: 
- `correctedSearchTerm`: String

### RecipeVideoStreams
Recipe Video Streams

- `sd`: String
- `hd`: String

### RecipeVideos
Recipe preparation details

- `preparation`: RecipeVideo
- `secondaryPreparation`: RecipeVideo
- `tips`: 

### RecommendationWeek
Weeks for which recommendations are available for user

- `startDate`: String
- `endDate`: String
- `activeWeek`: Boolean

### RedirectAction
Whenever FE needs to be redirected then a RedirectAction is used.

- `redirectUri`: String

### RedirectToLmn
The URLs related to the Loyalty Management Netherlands b.v. (LMN) authentication flow.

- `loginUrl`: String
- `exitUrl`: String
- `registerUrl`: String

### RuleValidated
Contains the rule and result against user password

- `code`: String
- `isValid`: Boolean

### SaveMemberRecipeMutationResult
Save member recipe result

- `status`: MutationResultStatus
- `result`: MemberRecipe
- `errorMessage`: String

### Savings
Calculated Saving for either Instore or WithSubscriptions

- `min`: Money
- `max`: Money

### SavingsPeriod

- `year`: Int
- `startMonth`: Int
- `endMonth`: Int

### ScrapedRecipe
Recipe scraped from external source

- `title`: String
- `description`: String
- `courses`: String
- `preparation`: RecipePreparation
- `imageUrl`: String
- `author`: RecipeAuthor

### SearchContextualPayload
The result of a contextual search

- `refinedQuery`: String
- `products`: 

### SearchFacetConfig
Available Facet Configuration

- `label`: String
- `name`: String
- `type`: SearchFacetType

### SearchFacetOption
A facet option type

- `value`: PrimitiveTypeScalar
- `name`: String
- `matches`: Int
- `suboptions`: 
- `selected`: Boolean
- `image`: ImageSet

### SearchFacetPayload
Facet payload type

- `label`: String
- `name`: String
- `type`: SearchFacetType
- `rangeSelectedFacetValues`: SearchRangeSelectedFacetValues
- `options`: 

### SearchFacetsPayload
Search facets payload contains all the available facets

- `facets`: 
- `quickFilters`: 

### SearchPayload
Search payload type

- `products`: 
- `facets`: SearchFacetsPayload
- `similarProducts`: 
- `sponsoring`: SearchSponsoringPayload
- `totalFound`: Int
- `searchId`: String

### SearchProductCard
Search product card type

- `id`: Int
- `title`: String
- `additionalInformation`: String
- `brand`: String
- `category`: String
- `hqId`: Int
- `imagePack`: 
- `icons`: 
- `availability`: ProductCardAvailability
- `salesUnitSize`: String
- `summary`: String
- `highlights`: String
- `highlight`: String
- `interactionLabel`: String
- `virtualBundleProducts`: ProductCardVirtualBundleItem
- `webPath`: String
- `isSample`: Boolean
- `isMedicine`: Boolean
- `isMedicalDevice`: Boolean
- `hasListPrice`: Boolean
- `minBestBeforeDays`: Int
- `isDeactivated`: Boolean
- `ageCheck`: Boolean
- `shopType`: ProductCardShopType
- `privateLabel`: Boolean
- `isSponsored`: Boolean
- `priceV2`: ProductCardPriceV2
- `priceV2Original`: ProductCardPriceV2Original
- `externalWebshopUrl`: String
- `variant`: ProductCardVariant
- `variants`: 

### SearchProductCardPayload
Search payload type

- `products`: 
- `facets`: SearchFacetsPayload
- `sponsoring`: SearchSponsoringPayload
- `totalFound`: Int
- `searchId`: String

### SearchRangeSelectedFacetValues
The min and/or max values for a range selected facet type

- `min`: PrimitiveTypeScalar
- `max`: PrimitiveTypeScalar

### SearchSponsoringPayload
SearchSponsoringPayload stores sponsoring related data (eg: 'auctionId')

- `auctionId`: String

### SearchSuggestion
A suggested section of search results with suggested search results.

- `label`: String
- `value`: String
- `href`: String
- `type`: SearchSuggestionType
- `icon`: String
- `suggestions`: 

### SearchTaxonomy
A representation of a product taxonomy.
Primarily used for displaying sub-taxonomy buttons in taxonomy browsing
and in search results.

- `id`: Int
- `name`: String
- `slug`: String
- `parents`: 
- `children`: 
- `images`: 
- `productIds`: 
- `totalProductCount`: Int

### SelfServiceAccordions
SelfServiceAccordions config component

- `query`: String
- `accordions`: 

### SelfServiceArticle
SelfServiceArticle configured component response

- `content`: String

### SelfServiceBreadCrumbLink
BreadCrumb link

- `title`: String
- `href`: String

### SelfServiceBreadCrumbs
SelfServiceBreadCrumbs configured component response

- `breadcrumbs`: 

### SelfServiceContentError
Error Status for customer care module to be able to handle errors

- `code`: Int
- `message`: String
- `type`: SelfServiceContentErrorTypes

### SelfServiceContentPathResponse
Self Service Content Response

- `query`: String
- `buttons`: SelfServiceQuickEntryButtons
- `links`: SelfServiceLinks
- `accordions`: SelfServiceAccordions
- `metadata`: SelfServiceMetaData
- `breadcrumbs`: SelfServiceBreadCrumbs
- `article`: SelfServiceArticle
- `error`: SelfServiceContentError

### SelfServiceContentQueryResponse
Self Service Content Response

- `query`: String
- `breadcrumbs`: SelfServiceBreadCrumbs
- `buttons`: SelfServiceQuickEntryButtons
- `links`: SelfServiceLinks
- `accordions`: SelfServiceAccordions
- `error`: SelfServiceContentError

### SelfServiceLinks
SelfServiceLinks config component

- `links`: 

### SelfServiceMetaData
SelfServiceMetaData configured component response

- `data`: SelfServiceMetaDataData

### SelfServiceMetaDataData
MetaData data

- `title`: String
- `description`: String
- `canonical`: String
- `follow`: Boolean
- `index`: Boolean

### SelfServiceQuickEntryButtons
SelfServiceQuickEntryButtons configured component response

- `buttons`: 

### ServiceCharge
The price of a slot

- `price`: Money
- `defaultPrice`: Money
- `discountLevel`: DiscountLevel

### Settlement
Settlement

- `id`: Int
- `type`: String
- `order`: Order
- `product`: Product
- `quantity`: Int
- `comment`: String
- `amount`: Money
- `refundMethod`: SettlementRefundMethod
- `imageIds`: 
- `images`: 
- `cancellable`: Boolean
- `settlementFraction`: Float
- `depositLabel`: String

### SettlementAfterDelivery
CostOverview Settlement item

- `title`: String
- `price`: Money
- `amount`: Int

### SettlementImage
Settlement image

- `url`: String
- `imageId`: Int

### SettlementImageV2
SettlementV2 image

- `imageId`: String
- `url`: String

### SettlementItems
Divided settlement items by type

- `restitutions`: 
- `deposits`: 
- `serviceMemos`: 
- `otherSettlements`: 

### SettlementV2
Settlement V2

- `id`: String
- `type`: String
- `order`: Order
- `product`: Product
- `quantity`: Int
- `comment`: String
- `amount`: Money
- `imageIdsV2`: 
- `imagesV2`: 
- `cancellable`: Boolean
- `settlementFraction`: Float
- `feedbackLabel`: String
- `title`: String
- `subtitle`: String

### SettlementsTotal
List of settlements containing settlement's and a message + a conclusion for all settlements

- `settlements`: 
- `totalAmount`: Money

### SharedEntityMutationResult
Result of a mutation done to a favorite list's shared entity member's permissions

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: String

### SingularPluralName
GraphQL equivalent for Noun

- `singular`: String
- `plural`: String

### Slot
Slot for delivery or pickup order

- `date`: DateTime
- `dateFormatted`: String
- `startTime`: DateTime
- `startTimeFormatted`: String
- `endTime`: DateTime
- `endTimeFormatted`: String
- `day`: Int
- `deliveryLocationId`: Int
- `isFullyBooked`: Boolean
- `nudgeType`: 
- `pickupLocationId`: Int
- `serviceCharge`: ServiceCharge
- `shiftCode`: String

### SlotDay
Day containing slots

- `date`: DateTime
- `dateFormatted`: String
- `slots`: 
- `isFullyBooked`: Boolean

### SolidPantryColorSet
Solid color set (only has one primary color)

- `fillStyle`: ThemeFillStyle
- `primaryColor`: String
- `textColor`: String

### SpotlightCardAdvertisement
Spotlight Item Advertisement

- `metadata`: AdvertisementMetadata
- `card`: ContentSpotlightCard

### StampShareable
A gifting or requesting document.

- `shareableId`: String
- `programId`: Int
- `expiresAt`: String
- `links`: StampShareableLinks
- `initiatorDetails`: StampShareableInitiatorDetails

### StampShareableInitiatorDetails
Details of the initiator who is gifting/requesting stamps

- `firstName`: String
- `sharingAmount`: Int

### StampShareableLinks
Links of a gifting or requesting document.

- `href`: String
- `deeplink`: String

### StampShareableMutationResult
Mutation result of creating or delete a gift link.

- `status`: MutationResultStatus
- `errorMessage`: String
- `result`: StampShareable

### StampShareableValidationResult
Validation result of a ShareableLink

- `shareableId`: String
- `valid`: Boolean
- `invalidReason`: StampShareErrorReason

### StampTransferBalanceChange
Chage of the stamp balance

- `old`: Int
- `new`: Int
- `change`: Int

### StampTransferMutationResult
Mutation result ot transfering stamps.

- `status`: MutationResultStatus
- `errorMessage`: String
- `errorReason`: StampShareErrorReason
- `result`: StampTransferResult

### StampTransferResult
Result ot transfering stamps.

- `shareableId`: String
- `programId`: Int
- `balanceChange`: StampTransferBalanceChange
- `initiatorDetails`: StampShareableInitiatorDetails

### Stores
Store information for store details

- `id`: Int
- `name`: String
- `phone`: String
- `address`: StoresAddress
- `storeType`: StoresType
- `geoLocation`: GeoLocation
- `openingDays`: 
- `services`: 
- `distance`: Float
- `productAvailability`: ProductAvailabilityStatus

### StoresAddress
Address of the store

- `street`: String
- `houseNumber`: String
- `houseNumberExtra`: String
- `postalCode`: String
- `city`: String
- `countryCode`: String

### StoresCityCount
Store count per place

- `city`: String
- `count`: Int

### StoresGeoLocationSuggestion
Suggestion for a possible location

- `description`: String
- `match`: StoresGeoLocationSuggestionMatch
- `geoLocation`: GeoLocation

### StoresGeoLocationSuggestionMatch
Matched substring of the result

- `position`: Int
- `length`: Int

### StoresOpeningDay
opening day

- `dayName`: String
- `date`: String
- `openingHour`: StoresOpeningHour
- `type`: StoresOpeningDayType

### StoresOpeningHour
opening hour

- `openFrom`: String
- `openUntil`: String

### StoresOpeningHourCount
For each kind of opening hour, the amount of times it occurs in the search

- `type`: StoresOpeningHourType
- `count`: Int

### StoresResultAggregationType
Aggregated results after searching for stores

- `services`: 
- `openingHours`: 
- `storeTypes`: 

### StoresResultType
Result of searching for stores

- `result`: 
- `pageInfo`: PageInfo
- `aggregation`: StoresResultAggregationType

### StoresService
Services stores can provide

- `code`: ID
- `description`: String
- `shortDescription`: String

### StoresServiceCount
Services and amount of stores that provide them

- `service`: StoresService
- `count`: Int

### StoresTypeCount
For each kind of store, the amount of times it occurs in the search

- `type`: StoresType
- `count`: Int

### SubscriptionAvailability
The availability of a certain subscription. For example: a delivery bundle cannot be sold to a member with a non-deliverable zipcode

- `available`: Boolean
- `gridType`: String

### SubscriptionAvailable
Available subscription for purchase, including terms and pricing

- `id`: ID
- `code`: String
- `type`: String
- `description`: String
- `duration`: Int
- `price`: SubscriptionPrice
- `priceBeforeDiscount`: SubscriptionPrice
- `deliveryScheme`: String
- `discounts`: 
- `isTrialSubscription`: Boolean
- `postTrialSubscription`: SubscriptionPostTrial
- `isPromoCodeApplicable`: Boolean

### SubscriptionAvailableDeliveryDays
The available delivery days per subscription code

- `availableDeliveryDays`: 

### SubscriptionBundle
A bundle containing a possible combination of subscriptions

- `bundleId`: ID
- `subscriptionsAvailable`: 
- `bundlePrice`: SubscriptionPrice
- `bundlePriceBeforeDiscount`: SubscriptionPrice
- `connectedBundle`: SubscriptionBundle

### SubscriptionBundlePromoCodeApplyResult
Result after applying a promotion code to a subscriptionBundle

- `result`: 
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionBundleSubscribeMutationResult
Result of subscribing to a subscription bundle

- `result`: SubscriptionPaymentIDeal
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionCancelMutationResult
Result after cancelling your subscription

- `result`: SubscriptionCurrent
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionCancelMutationResultTemp
Result after cancelling you subscription. Directly cancelled returns null subscription and cancelled after end of period returns a subscription with
a certain end date

- `result`: SubscriptionPlan
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionCancellation
Subscription cancellation specification

- `earliestDate`: Date
- `latestDate`: Date
- `date`: Date
- `isImmediatelyCancellable`: Boolean

### SubscriptionChange
Comparison between current and new subscription

- `currentSubscription`: SubscriptionAvailable
- `newSubscription`: SubscriptionAvailable
- `residualPrice`: Money
- `finalPrice`: Money

### SubscriptionCodeDeliveryDays
The available delivery days for a specific subscription code

- `subscriptionCode`: String
- `deliveryDays`: 

### SubscriptionCurrent
The current subscription of a certain type of a member, such as delivery bundle and premium.

- `id`: ID
- `definitionId`: Int
- `description`: String
- `type`: String
- `code`: String
- `statusCode`: Int
- `deliveryScheme`: String
- `extension`: SubscriptionExtension
- `duration`: Int
- `startDate`: String
- `endDate`: String
- `cancellation`: SubscriptionCancellation
- `isActive`: Boolean
- `isChangeable`: Boolean
- `isCancellable`: Boolean
- `isTrialSubscription`: Boolean
- `postTrialSubscription`: SubscriptionPostTrial
- `originalPrice`: Money
- `sellPrice`: Money
- `isPromoCodeApplicable`: Boolean
- `isEligibleForSavingWarranty`: Boolean
- `isReactivateEnabled`: Boolean

### SubscriptionCurrentPromoCodeApplyResult
Result after applying a promotion code to a current subscription

- `result`: SubscriptionCurrent
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionDefinitionChangeMutationResult
Result after changing your subscription

- `result`: SubscriptionCurrent
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionDefinitionChangeMutationResultTemp
Result after changing your subscription

- `result`: SubscriptionPlan
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionDiscount
Subscription discount information

- `code`: String
- `description`: String
- `discount`: Money
- `promoCode`: String

### SubscriptionEstimatedPremiumSavings
Estimated premium savings of remaining subscription days based on the members current purchase history

- `subscriptionActiveDays`: Int
- `estimatedYearlySavedAmount`: Float
- `estimatedYearlySavedStampsAmount`: Int
- `estimatedRemainingDaysSavedAmount`: Float
- `estimatedRemainingDaysSavedStampsAmount`: Int
- `averageDailySavedAmount`: Float
- `averageDailySavedStampsAmount`: Float

### SubscriptionExtension
Data regarding the extension of the subscription. Previously known as SubscriptionRenewal

- `date`: String
- `originalPrice`: Money
- `sellPrice`: Money
- `discounts`: 

### SubscriptionFixedDeliverySlot
The available fixed delivery slot for a specific subscription. Only available for BEZORG and B2B type.

- `date`: String
- `from`: String
- `until`: String
- `score`: Float
- `nudgeTypes`: 
- `subscriptionValidityType`: String

### SubscriptionFixedDeliverySlotCurrent
The current chosen fixed delivery slot of the member

- `startDate`: String
- `slotStartTime`: String
- `slotEndTime`: String
- `isSustainableSlot`: Boolean

### SubscriptionNextSubscription
Description of next subscription

- `id`: ID
- `code`: String
- `description`: String
- `startDate`: String
- `price`: Money

### SubscriptionPaymentIDeal
Result of starting of payment

- `id`: String
- `status`: SubscriptionPaymentStatus
- `paymentStatus`: SubscriptionPaymentStatus
- `url`: String
- `mutationId`: String
- `action`: Action

### SubscriptionPaymentIDealIssuer
Available iDeal issuers

- `id`: Int
- `name`: String
- `brandUrl`: String

### SubscriptionPaymentOption
Available payment option for the member

- `option`: String
- `status`: String
- `issuers`: 

### SubscriptionPlan
A subscription. Identical to MemberSubscription. Use this instead. Note: Subscription is a reserved type in GraphQL for a real-time websocket connection, so we call this SubscriptionPlan instead.

- `id`: ID
- `type`: String
- `code`: String
- `statusCode`: Int
- `deliveryScheme`: String
- `price`: Money
- `renewal`: SubscriptionRenewal
- `duration`: Int
- `startDate`: String
- `endDate`: String
- `cancellation`: SubscriptionCancellation
- `isActive`: Boolean
- `isChangeable`: Boolean
- `isCancellable`: Boolean
- `nextSubscription`: SubscriptionNextSubscription

### SubscriptionPollPaymentStatus
The payment status of a subscription. May contain response data regarding polling.

- `id`: String
- `mutationId`: String
- `paymentStatus`: PaymentStatus
- `action`: Action

### SubscriptionPostTrial
Description of the subscription a member automatically gets after a trial subscription

- `id`: ID
- `code`: String
- `duration`: Int
- `description`: String
- `startDate`: String
- `sellPrice`: Money
- `monthlySellPrice`: Money
- `discounts`: 

### SubscriptionPremiumSavings
The amount saved due to the premium subscription

- `lastUpdated`: String
- `subscriptionStartDate`: String
- `bioDiscountAmount`: Money
- `bonusBoxDiscountAmount`: Money
- `bonusBoxDiscountPremiumExclusiveAmount`: Money
- `additionalAirmilesObtained`: Int
- `additionalAirmilesObtainedInMoney`: Money
- `additionalPurchaseStampsObtained`: Int
- `addtionalPurchaseStampsPurchaseAmount`: Money
- `additionalPurchaseStampsInterest`: Money
- `additionalSavingsStampsObtained`: Int
- `deliveryBundleDiscountAmount`: Money
- `terraDiscountAmount`: Money
- `totalSavedAmount`: Money
- `totalAdditionalStamps`: Int
- `periods`: 

### SubscriptionPrice
Subscription pricing specification

- `fullDuration`: Money
- `monthly`: Money
- `monthlyEquivalent`: Money

### SubscriptionReactivateResult
Result after applying reactivation to a current trial subscription

- `result`: SubscriptionCurrent
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionRenewal
Subscription next renewal pricing specification

- `date`: String
- `originalPrice`: Money
- `sellPrice`: Money

### SubscriptionSavingPotential
The potential savings a member can get by getting a subscription. Hence, the discount they are missing out on.

- `premiumSavingPotential`: Float
- `hasOrderedRecently`: Boolean

### SubscriptionSavings
Subscription savingsamount

- `inStore`: Savings
- `withSubscriptions`: Savings

### SubscriptionSavingsCarousel
Carousel for premium savings

- `id`: ID
- `index`: Int
- `title`: String
- `imageUrl`: String
- `link`: SubscriptionSavingsCarouselLink
- `theme`: SubscriptionSavingsCarouselTheme
- `isCompleted`: Boolean
- `isUpdatedByRequirements`: Boolean

### SubscriptionSavingsCarouselLink
The link of the carousel. Contains URL and type of the link ("DEEPLINK", "WEB")

- `url`: String
- `type`: SubscriptionSavingsCarouselLinkType

### SubscriptionSavingsCarouselUpdateMutationResult
Result of updating a premium savings carousel

- `result`: 
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionSetFixedDeliverySlotMutationResult
Result after setting your fixed delivery slot

- `result`: SubscriptionFixedDeliverySlotCurrent
- `status`: MutationResultStatus
- `errorMessage`: String

### SubscriptionSummary
A summary of the current subscription.

- `hasHadSubscription`: Boolean
- `currentSubscriptionCode`: String

### Supplier
Store supplier information

- `id`: Int
- `name`: String
- `address`: SupplierAddress
- `category`: String
- `geoLocation`: GeoLocation
- `supplierType`: String
- `inBetterForNatureProgram`: Boolean

### SupplierAddress
Address of the supplier

- `address`: String
- `postalCode`: String
- `city`: String
- `countryCode`: String

### TargetedOffer
A type containing a targeted offer. This is an offer that is only valid for a single member,
for example a Bonus Box or Digital Scratching offer.

- `id`: UUID
- `externalIds`: TargetedOfferExternalIds
- `status`: TargetedOfferStatus
- `validityPeriod`: TargetedOfferPeriod

### TargetedOfferAllocationMutationResult
A type containing the result of an allocation mutation. Allocations might fail if the member
is not eligible for (another) targeted offer for a program. For example, they might not
have enough points available to trade for an allocated offer.

- `status`: MutationResultStatus
- `errorMessage`: String
- `errorReason`: TargetedOfferAllocationErrorReason
- `targetedOffer`: TargetedOffer

### TargetedOfferChallenge
A type containing the allocation challenge (tie-breaker question) used in the Belgian domain.
Contains a question and a list of possible answers (multiple choice).

- `id`: Int
- `question`: String
- `answers`: 

### TargetedOfferChallengeAnswer
A type containing a possible answer for an allocation challenge (tie-breaker question) used in
the Belgian domain.

- `id`: Int
- `label`: String
- `answer`: String

### TargetedOfferExternalIds
A type containing various secondary, external, IDs of a targeted offer.

- `offerId`: UUID
- `programId`: Int
- `hqId`: String

### TargetedOfferPeriod
A type containing a timespan used by targeted offers. The start and end are both dates, and
as such do not contain time information.

- `start`: Date
- `end`: Date

### TaxonomyImage
A product brand

- `width`: Int
- `height`: Int
- `url`: String
- `imageFormat`: String

### TestHttpHeader
Http Header (has name & value)

- `name`: String
- `value`: String

### TestMutationResult
A MutationResult with an Int

- `status`: MutationResultStatus
- `result`: Int
- `errorMessage`: String

### TimeSlot
Timeslot

- `timeV2`: CustomerCareTime
- `load`: Load
- `active`: Boolean

### TrackAndTraceV2
Track and Trace

- `orderId`: Int
- `orderType`: OrderType
- `type`: TrackAndTraceType
- `message`: String
- `etaBlock`: EtaBlock
- `realisedDeliveryTime`: String

### TransactionAmountCurrency
Amount of the transaction

- `value`: Float
- `currency`: EvalueCurrencyEnum

### UpdateFeedbackResult
Update feedback result

- `status`: MutationResultStatus
- `errorMessage`: String
- `feedbackId`: String

### UpdatedRecipeRating
Defines recipe rating after update

- `count`: Int
- `average`: Float

### UsageWithExtra
Preparation instructions formatted for web ui

- `extra`: String
- `contentLines`: 

### UserChallengeAction
User Challenge that is performed by the user doing something

- `transactionId`: String
- `authorizationId`: String
- `data`: String
- `sdk`: PaymentSDK
- `issuerId`: String

### ValidationResponse
The returned value after GetMemberCodeResponse is normalized

- `status`: CodeStatus

### VideoBbw
Blue Billywig video

- `id`: String
- `title`: String
- `description`: String
- `author`: String
- `originalFilename`: String
- `src`: String
- `thumbnails`: VideoBbwThumbnail
- `thumbnail`: ContentImage
- `assets`: 
- `width`: String
- `height`: String
- `updatedDate`: String
- `createdDate`: String

### VideoBbwAsset
Stream file/version/rendition

- `src`: String
- `length`: String
- `width`: String
- `height`: String
- `bandwidth`: String
- `status`: String

### VideoBbwThumbnail
Image related to the video

- `src`: String
- `width`: String
- `height`: String
- `main`: String

### WeekPlannerDay
Week planner day

- `date`: String
- `items`: 

### WeekPlannerMemberRecipeItem
Week planner member recipe item

- `id`: Int
- `type`: WeekPlannerItemType
- `recipeId`: Int
- `recipeTitle`: String
- `memberRecipe`: MemberRecipe

### WeekPlannerNoteItem
Week planner member recipe item

- `id`: Int
- `type`: WeekPlannerItemType
- `text`: String

### WeekPlannerOverview
Week planner overview

- `days`: 

### WeekPlannerProductItem
Week planner member recipe item

- `id`: Int
- `type`: WeekPlannerItemType
- `productId`: Int
- `productTitle`: String
- `product`: Product

### WeekPlannerRecipeItem
Week planner recipe item

- `id`: Int
- `type`: WeekPlannerItemType
- `recipeId`: Int
- `recipeTitle`: String
- `recipe`: RecipeSummary

### WeekPlannerRecipeSuggestion
Week planner recipe suggestion

- `recipeSuggestions`: 
- `preferencesOptedIn`: Boolean

## INPUT_OBJECT Types (145)

### AddCardCheckInput
The given input to check if this member-card is not already connected to another member

### AddressCompleteInput
Address input type, all fields mandatory just as the Address type

### AddressPartialInput
Partial Address input type, all fields optional

### AddressSearchInput
Address input type for searching addresses

### AndroidConfigurationInput
This type represents specific for Android configuration settings for push notifications.

### Answer
Answer to an Lifestylecheck Question

### BancontactMobileTransactionRequest
Request for Bancontact Mobile Transaction

### BasketDelete
To delete an item from basket

### BasketInput
Basket input type

### BasketItemInput
Type of items in the basket

### BasketItemRecommendationsInput
Arguments for checkout recommendations endpoint for basketitems. These are the items that are currently in your the basket.

### BasketMutation
We can mutate both products and "vage termen" to the Basket but also to List

### CaptchaVerifyInput
Input parameters usde to verify the query/mutation is not made
by a robot

### CheckoutConfirmDctInput
DCT Payment Options Information

### CheckoutConfirmDctInputV4
DCT Payment Options Information

### CheckoutConfirmGiftCardInfo
NPL Giftcard

### CheckoutConfirmIDEALInput
input for creating new iDEAL Payment

### CheckoutConfirmOrderPayload
Required body payload for submitting the order unified with WEB, iOS and Android

### CheckoutConfirmOrderPayloadV4
Required body payload for submitting the order unified with WEB, iOS and Android

### CheckoutConfirmSddInput
Sepa Direct Debit Payment Options Information

### CommunicationConsentInput
This type represents the input for adding a communication consent.

### ConsentInput
Default information needed to add a consent

### ContentCuratedListItemInput
Arguments to get product suggestions

### ContentDeliveryGridOptions
CMS Delivery Grid Arguments

### ContentFlexPageOptions
Content CMS flex page options

### ContentFlexPageWebUrlOptions
Content CMS flex page options

### ContentFooterLinksOptions
CMS Footer Links Arguments

### ContentMegaMenuLinksOptions
CMS MegaMenu Links Arguments

### ContentMobileComponentOptions
Content mobile CMS options

### ContentOptInOptions
CMS Opt In Arguments

### ContentPageOptions
CMS Page Arguments

### ContentResourceBundleOptions
Resource bundle input options

### ContentTargetedDocumentOptions
Content CMS targeted document options

### ConversationHandlerContextMutation
Handler context

### CookBookInputRecipeIngredient
Same as cookBookRecipeIngredient

### CookBookInputRecipeMeta
Same as CookBookRecipeMeta but we need input

### CookBookInputRecipeTimes
Same as CookBookRecipeTimes

### CookBookMemberAddProfileInput
input when adding new cookbook profile

### CookBookMemberUpdateProfileInput
input when updating a profile, when null not updated

### CookieConsentInput
Input representing Cookie Consent

### CookieConsentVersionInput
Input representing Cookie Consent version

### CreatePayPaymentRequest
Request payload for paying a payment request

### CreatePaymentRequestTransactions
Which transactions to create per request

### CustomerContactDetailsInput
Customer Contact Details

### DeviceInput
This type represents properties of a device that sends its notification settings.

### DeviceNotificationSettingsInput
This type represents device notification settings.

### EnrollAssetRequestInput
Input to enroll personal asset

### ExperimentalPageInput
Pagination descriptor

### ExperimentalSearchFacetInput
Facet input type allows filtering search results

### ExperimentalSearchInput
Search base input type

### ExperimentalSearchIntentInput
Search intent input

### ExperimentalSearchProductsInput
Search products input

### ExperimentalSearchRangeSelectedFacetInput
The input min and/or max values for a RANGE selected facet

### ExperimentalSearchValuesSelectedFacetInput
The input values for a selected facet (when the facet type is not RANGE)

### FavoriteListProductMutation
Add a product to a favorite list or update the quantity of an existing one.

### FullfillmentDateRange
Date range start and end date

### GeoLocationBoundariesInput

### GeoLocationInput
Geolocation as input

### GeoLocationPointInput

### GiftCardCredentials
Required information for adding a giftcard

### IDEALTransactionRequest
Request for iDEAL Transaction

### IOSConfigurationInput
This type represents specific for iOS configuration settings for push notifications.

### IngredientToOverride

### InvoiceInput
Input given to fetch invoices

### InvoiceV2Input
Input given to fetch invoices

### KeyValueInput
Generic key value input for subscriptionPaymentIDealStatusV2 queryParam

### MandateData
Data of the mandate to be enrolled

### MemberAddressInput
Member address input

### MemberBusinessActivityInput
Member business activities input

### MemberCompanyDetailsInput
Member company details input

### MemberConsentsInput
Provide the consents which should be added or removed when doing the memberConsents mutation

### MemberFoodPreferencesInput
Member food preference input

### MemberFoodPreferencesV2Input
Member food preference v2 input

### MemberPatchEntitiesInput
Member patch entities input

### MemberPatchEntityInput
Member patch entity input

### MemberPatchIdEntitiyInput
Member patch id entity input

### MemberPatchIdEntityAddInput
Member patch id entity add input

### MemberPersonalDetailsInput
Member personal details input

### MemberRecipeAuthorInput
Member recipe author input

### MemberRecipeInput
Member recipe input

### MemberRecipeOriginInput
Member recipe origin input

### MemberRecipeProductSuggestionsInput
memberRecipeProductSuggestions input type

### MessageCenterMarkingMessageReadInput
Input required to mark message center message as read or unmark message as read

### MessageCenterMessageToUpdate
Message to be updated

### MessageCenterUpdateIsDeletedByUserInput
Input required to update if message was deleted by user

### MessengerCreateConversationContext
Context from the member

### MilesAccountFields
Input for an easy onboarding miles account

### NameInput
Member name input

### NewGroceryItem
Arguments for creating a new grocery list item

### NewNote
Arguments for creating a new note

### OffsetLimitPagination
Offset / Limit pagination input type

@example Using a default value, so that pagination input is not required.
extend type Query {
   queryWithPagination(
     pagination: OffsetLimitPagination! = { offset: 0, limit: 50 }
   ): [SomeResultType!]!
}

### OrderDeliveryAddress
Address where the order should be delivered.

### OrderSlot
Slot of the order.

### PageInput
Pagination descriptor

### PageSizePagination
Page / Size pagination input type

@example Using a default value, so that pagination input is not required.
extend type Query {
   queryWithPagination(
     pagination: PageSizePagination! = { page: 0, size: 50 }
   ): [SomeResultType!]!
}

### PaymentAmountRequest
Amount of the transaction that is requested

### PaymentMethodUpdateInput
The input values to submit the new payment method for a B2B user, either by invoice or payment slip

### PaymentRequestAuthorizeInput
Input to authorize transaction in payment request

### PaymentsAuthorizeInput
Input to authorize transaction

### PaymentsChallengeResult
Data of the challenge result provided by PSP

### PaymentsDctCardOnboardingInput
Input required to onboard DCT Tokens

### PickupLocationsFiltersInput
Filters for pickup locations

### PriceLineItem
a productId with the quantity

### ProductCrossSellsV2Input
Input type for the productCrossSellsV2 query
V2 endpoint filters out unavailable products whereas V1 does not

### ProductSearchInput
Product search input

This is a temporary alternative for the old basketProductCardsByName query.
Should be replaced by a proper ProductSearch when it is available.

### ProductsInput
Input for 'products' query.
Optionally has a date.

### PromotionSearchInput
Promotion Search Input

### PushNotificationsSitespectExperiment
Input for the Sitespect push experiment

### PushNotificationsTokenUpdateInput
Input for the pushtokenSaveToken mutation

### QuestionAnswers
Customer provided/picked Answers to an Lifestylecheck Question

### QuestionnaireAnswer
Questionnaire answer

### QuestionnaireFormSubmitInput
Questionnaire form submit input

### RecipeProductSuggestionV2Input
recipeProductSuggestionsV2 input type

### RecipeSearchParams
Query options for search request

### RecipeSearchQueryFilter
Filter options for recipe search query

### RecipeShoppableIngredientName
Recipe ingredient and quantity unit comes in singular and plural names

### RecipeShoppableSaveSelectionParams
Input type for the mutations that saves user selections of the ingredients to products mappings

### RecipeShoppableSelection
Single selection of the ingredient to product mappings

### RecipeShoppableSelectionIngredient
Ingredient type for mappings selection

### RecipeShoppableSelectionProduct
Selection product type for mappings

### RecipeSourceInput
Recipe source input

### RecipeVideoSearchParams
Input for searching recipe videos

### ReferenceMandateData
Data of the reference mandate to be enrolled

### ReferencePaymentData
Reference Payment Data

### SavingsPeriodInput

### ScannedIngredient
ScannedIngredient

### ScannedIngredients
The ingredients that are scanned by our AI.

### SearchBrandInput
Search products for given brand input

### SearchCategoryInput
Search products in category input

### SearchContextualInput
Search Contextual input

### SearchFacetInput
Facet input type allows filtering search results

### SearchInput
Search base input type

### SearchIntentInput
Search intent input

### SearchPreviouslyBoughtInput
Search previously bought input

### SearchProductsInput
Search products input

### SearchRangeSelectedFacetInput
The input min and/or max values for a RANGE selected facet

### SearchValuesSelectedFacetInput
The input values for a selected facet (when the facet type is not RANGE)

### StoresFilterInput
Filter options for stores

### SubscriptionSavingsOptions
Input options for calculating the hypothetical savings after having a subscription

### SubscriptionSlotInput
Date and time slot for a subscription slot

### TestWithArgumentsInput
Input type for a testWithArguments mutation

### UnsubscribeBlockerReasonInput
Given input to fetch blocker reasons

### UpdateRecipeWeekPlannerInput
Week planner update input

### WeekPlannerDayInput
Week planner day update input

### WeekPlannerItemInput
Week planner day item update input

## ENUM Types (288)

### ActivatePersonalPromotionMessage
ActivatePersonalPromotionStatus

### AddSMSReminderStatus
Additional status information that is returned from the service

### AdditionalReceiptItemType
Type of the AdditionalReceiptItem

### AdvertisementDeviceType
An advertisement's device type. This refers to the type of device that is used to render the advertisement on.

### BargainMarkdownType

### BasketOrderState
States order can be in

### BasketShoppingType
Enum for the basket shopping type eg: "DELIVERY", "PICKUP"

### BasketSortTypeInput
Possible basket sorting types

### BonusCategoryType
Bonus category type

### BonusFilterType
Bonus filter types. When passed it allows to filter products that have bonus price of specified type

### BonusLaneType
Bonus Lane type that can be free delivery or personal bonus

### BonusPromotionType
Promotion type

### BonusSegmentActivationStatus
Activation statuses for segment

### BonusSegmentDiscountLabelCodes
BonusSegmentDiscountLabelCodes

### BonusSegmentDiscountShieldEmphasis
BonusSegmentDiscountShieldEmphasis

### BonusSegmentDiscountShieldTheme
BonusSegmentDiscountShieldTheme

### BonusSegmentState
BonusSegmentState contains the states possible for a bonus segment

### BonusSegmentType
BonusSegmentType

### BonusTheme
Bonus Themes

### BudgetType

### ButtonType
ButtonType

### CacheControlScope
cache control

### CategoriesFilterSet
Predefined filter sets for categories

### ChallengeType
Challenge Type

### Channel
The channel subscriptionBundleSubscribe mutation is called from. For example: WEB, ANDROID, IOS. Set to 'WEB' by default.

### ChannelAvailabilityStatus
Customer care channel availability status

### ChannelPlatform
Customer care channels supported platforms

### ChannelType
Customer care channel type

### CheckoutATPLimitType
Limits you can hit on the checkout

### CheckoutErrors
CheckoutErrors Enum, DCT errors and Possible validation errors combined.

### CiamDeleteAccountMutationResultErrorCode
Error codes for mutation.ciamDeleteAccount

### CiamPasswordChangeMutationResultErrorCode
Error codes for mutation.ciamPasswordChange

### CiamPhoneNumberVerifyCodeErrorCode
Possible error codes for `Mutation.ciamPhoneNumberVerifyCode`

### CiamUserNameChangeMutationResultErrorCode
Error codes for mutation.ciamUsernameChange

### CmsMode
CMS mode enum

### CodeStatus
The status of the code

### CollectionCategoryRecipeType
Types of recipes that can be saved in categories

### CommunicationConsentChannel
Enum representing communication consent channel

### CompanyMemberStatusCode
Company Member Status Code enum

### ContentAccordionBlockType
Accordion block types

### ContentAccordionType
Type of Accordion (Manual/Search)

### ContentAllerhandeGridlaneComponentType
Allerhande flex page component types

### ContentAllerhandeSpotlightType
Content Spotlight Item Type

### ContentArticleBlockType
Article content blocks types

### ContentBaseMobileCMSComponentType
CMS mobile types

### ContentBaseTargetedContentCMSDocumentType
CMS targeted content document types

### ContentCMSComponentStatus
CMS Component Render Status

### ContentCMSComponentType
CMS Component Types

### ContentComponentType
Root Component Types

### ContentErrorTypes
Error types for content module

### ContentFlexPageComponentType
Flex page component types

### ContentFlexPageHeaderType
Flex page header types

### ContentFormDropdownRenderType
Form Dropdown Render type

### ContentFormFieldType
Types of the form fields

### ContentHighlightLaneRenderType

### ContentHighlightListRenderType

### ContentImageVariantType
Content image variant

### ContentMediaBlockType
Content media types (IMAGE, YOUTUBE, BLUEBILLYWIG)

### ContentProductLaneRenderType
Content Product Lane Render Type

### ContentSpotlightVisualRenderType
Content image render type

### ContentSubmenuFocusPoint
Submenu header focus point

### ConversationEventType
Available event types in a conversation

### ConversationHandlerType
Conversation Handler types

### ConversationHyperLinkType
the type of link

### ConversationStatus
Conversation status

### CookBookModerated
Moderation state for a cookbook

### CookieConsentLevel
Enum that represents supported Cookie Consent levels

### CustomerMemberType
The type of member

### CustomerServiceDocumentType
Search response document type

### DebitType
Debit Type

### DeliveryBundleDeliveryDays
Chosen days to receive deliveries

### DeliveryBundleDuration
Contract length of Delivery Bundle subscription

### DeliveryBundlePaymentStatus
Payment status of a delivery bundle

### DeliveryMethod
Method of Delivery

### DeliveryStatus
Status of delivery

### DeliveryTrackTraceInformationMessageType
The type of the message

### DeliveryTrackTraceInformationOrderType
The type of order

### DeliveryTrackTraceInformationState
List of possible states the order or delivery can have

### DepositComplaintDepositType
The type of deposit

### DepositComplaintExplanationType
Explanation type

### DepositComplaintReimbursementOption
Optional reimbursement

### DepositComplaintShoppingType
Shopping type of the order

### DepositComplaintStatus
Status of the order for deposit complaints

### DeviceType
This type represents device type.

### DiscountLevel
Service charge discount levels

### DomainsEnum
All available domains where the users email will be checked on

### EnrollAssetType
Asset type

### EntryPointContentVariantType
Enum type referencing the content variant of an entry point content. Mostly used for
internal purposes to differentiate between the different content types, as GraphQL's
query language already supports fetching based on type.

### EntryPointPropertyType
Enum for resolving the type of a EntryPointProperty.

### EntryPointTargetLinkType
Enum type referencing the target type of an entry point target.

### EvalueCurrencyEnum
Payment Transaction Currencies

### EvalueTransactionStatus
Transaction Status

### EvalueTransactionType
Transaction Type

### ExperimentalBonusPromotionType

### ExperimentalBonusSegmentType

### ExperimentalBonusTheme

### ExperimentalProductAvailabilityStatus

### ExperimentalProductIcon

### ExperimentalProductImageAngle

### ExperimentalProductPropertyIcon

### ExperimentalProductShopType

### ExperimentalProductUnavailableReason

### ExperimentalProductVariantType

### ExperimentalSearchFacetType
Possible facet types

### ExperimentalSearchIntentType
Possible search products intent types

### ExperimentalSearchProductsSortType
Possible sorting types for search

### ExplanationType
Explanation type

### FavoriteListSharedBy
Shared by enum

### FeedbackChannel
Feedback channel

### FeedbackFormType
Feedback form type for ces feedback

### FeedbackLoyaltyProgram
The type of loyalty program

### FeedbackOrderShoppingType
ShoppingType of the order eligible for feedback

### FeedbackShoppingType
The type of Shopping

### FeedbackSource
The source of feedback

### FeedbackType
The type of feedback.

### FolderState
Publication state

### FolderType
Folder type gets created by the backend based on the slug name

### FulfillmentCancellationReasonType
Type of cancellation for a fulfillment.

### FulfillmentDeliveryMethod
The method of delivery for a fulfillment.

### FulfillmentStatus
Whether the fulfillment is still open or already closed (or both).

### Gender
Gender options available within Ahold

### GroceryListType
Different types of grocery lists

### InvoiceOpenAmount
Invoices Open Amount filter

### InvoiceSortField
Invoices Sort Field

### InvoiceSortOrder
Invoices Sort order

### IosNotificationPermissionType
This type represents push notifications permission type for IOS device.

### LifestyleAdviceIconType
Icon Type of LifestyleAdvice

### LifestyleAdvicePath
Path of LifestyleAdvice

### LifestyleAdviceTopicLabel
Label of LifestyleAdviceTopic

### LifestyleCheckChapterIconType
Lifestylecheck Chapter Icon Type

### LifestyleCheckChapterStatus
State of Lifestylecheck Chapter

### LifestyleCheckIntakeAnswerIconType
Lifestylecheck Intake Answer Icon Type

### LifestyleCheckIntakeStatus
State of Lifestylecheck Intake

### LifestyleCheckNavigationIntent
Lifestylecheck Navigation Intent

### LifestyleCheckNavigationSuggestion
Lifestylecheck Navigation Suggestion

### LifestyleErrors
LifestyleErrors Enum

### LinkType
Link type

### LoyaltyCard
Returns the type of the loyalty card

### LoyaltyChallengeInfoLinkTargetType
An enumeration that contains the various types that a link to more information about the
challenge can be.

### LoyaltyChallengeStatus
A type containing the status of a challenge.

### LoyaltyChallengeStepStatus
A type containing the status of a step.

### LoyaltyChallengeStepTargetType
An enumeration that contains the various types of step targets.

### LoyaltyChallengeTermType
An enumeration that contains the various types that a term might related to. Used for
displaying the term correctly towards the member.

### LoyaltyPointsTransactionReason
@deprecated - Use LoyaltyPointsTransactionType
An enum to indicate the reason for a transaction.

### LoyaltyPointsTransactionType
An enum to indicate the type of a transaction.

### LoyaltyProgramCtaLinkType
Type of the CTA link.

### LoyaltyProgramState
Enum that indicates the program state.

### LoyaltyProgramType
Enum that indicates the program type.

### MemberDietPreferencesOptionType
All member diet options

### MemberDishPreferenceOptionType
All member dish types options

### MemberLoginState
Status of member

### MemberMeatFrequencyPreferenceOptionType
All member meat frequency options for FLEXITARIAN diet only

### MemberMeatPreferencesOptionType
All member protein options

### MemberNutritionPreferenceOptionType
All member nutrition preference options

### MemberProteinPreferencesOptionType
All member protein options

### MemberRecipeOriginType
Member recipe origin type

### MessengerChannelTypeEntrypoint
Chat channel entrypoint

### MessengerComponentInteractionComponentInteractionType
Messenger Component Interaction Component Interaction Type

### MessengerComponentInteractionComponentType
Messenger Component Interaction Component Type

### MessengerConversationEventType
All available event types in a conversation

### MessengerHandlerType
Messenger Handler types

### MessengerHyperLinkType
the type of link

### MessengerStatus
Status

### MilesErrorState
The different kinds of error state that a customer can be in with regards to the AH Miles program.

### MilesStatus
Enum with values of status for miles

### MutationResultStatus
Status of mutation result

### NotificationIcon
Enum for notification icons

### NotificationPriorityType
Enum Priority type of notification

### NotificationType
NotificationType

### NudgeType
Nudge types, which indicated the category of slots. SUSTAINABLE or EXCLUDE_INEFFICIENT. More to come soon

### OciShopSession
The method of an OCI shop session

### OperatingSystem
This type represents operating system.

### OrderDeliveryStatus
The status of an order's delivery.

### OrderShoppingType
Shopping type of the order

### OrderState
States order can be in

### OrderType
Order type

### PaymentActionType
Supported payment action types. QR_CODE for WEB(BEL) and REDIRECT for WEB(NLD)/ANDROID/IOS

### PaymentCardType
Card Type to update DCT Card (Required by Payments)

### PaymentChannel
Channel of processing

### PaymentIssuerAvailability
Identifies disruption in the issuer services

### PaymentMethod
Payment method that the customer can use

### PaymentMutationStatus
Payment Mutation Status

### PaymentOperationStatus
Defines status of an operation

### PaymentOperationType
Type of operation that is executed

### PaymentOptionStatus
Status of the PaymentOption

### PaymentProcessingSdkType
SDK key

### PaymentRequestStatus
Payment Request Status

### PaymentSDKType
Payment SDK Type

### PaymentStatus
Defining the status of an Payment

### PaymentTokenPspProvider
Provider of Psp

### PaymentTokenStatus
Card status of the payment token

### PaymentTransactionCurrencies
Payment Transaction Currencies

### PaymentTransactionStatus
Payment Transaction Status

### PaymentsAccessPermissionState
Status for an Personal Asset

### PaymentsPersonalAssetsAssetType
Asset type of personal assets

### PaymentsPersonalAssetsEnrollType
Enroll type for personal assets

### PaymentsPersonalAssetsStatus
Status of mandate

### PaymentsTokenOnboardingResponseCode
Replacement of GiftCard Response code

### PersonalPromotionErrorMessage
PersonalPromotionErrorMessage

### PickupLocationTypeCode
Pickup location type code

### PosReceiptStoreType
The type of store

### ProductAvailabilityStatus
availability status

### ProductCardAvailabilityStatus
availability status

### ProductCardBonusPromotionType
Promotion type

### ProductCardBonusSegmentType
BonusSegmentType

### ProductCardBonusTheme
Bonus Themes

### ProductCardIcon
Possible product icons

### ProductCardImageAngle
What angles are available for a product image

### ProductCardPromotionLabelType
Promotion Label Type

### ProductCardShopType
Possible product shop types

### ProductCardVariantType
Product variant type

### ProductComplaintClaimStatus
Product complaint claim status

### ProductComplaintShoppingType
Product complaint shopping type

### ProductIcon
Possible product icons

### ProductImageAngle
What angles are available for a product image

### ProductImageRenditions
Image renditions returned by the product image processor
in the form of D<width>x<height>

### ProductNutrientType
Types of nutrients that can be present in a product

### ProductPropertyIcon
Possible product property icons

### ProductReturnExplanationType
Explanation types

### ProductReturnShoppingType
Shopping type of the order

### ProductReturnStatus
Status of the order for product-returns

### ProductShopType
Possible product shop types

### ProductTradeItemCommunicationChannelType
Communication channel type

### ProductTradeItemResourceIconType
Resource icon type

### ProductUnavailableReason
reason why the product is unavailable

### ProductVariantType
Product variant type

### PromotionActivationStatus
Activation statuses for personal promotions

### PromotionLabelEmphasis
Promotion Label emphasis, used to indicate if one of the lines in the label should be larger than the others

### PromotionLabelMechanism
PromotionLabelMechanism

### PromotionLabelType
Promotion Label Type

### PromotionLabelVariant
Promotion label variant

### PromotionSegmentType
PromotionSegmentType

### PromotionType
Promotion type

### PromotionsFilterSet
Predefined filter sets for promotions

### Proposition
Type of proposition determines which products catalog will be used

### PurchaseStampMigrationStatus
Enum containing all the types of migration status.

### PushNotificationsDeviceType
Enum that represents type of the device that performs the mutation

### QuoteImageAlignmentType
The quote image alignment (left/right)

### RecipeFlag
Recipe flags indicating presence of certain tags, internalTags or classifications

### RecipeLaneArrangement
CMS configured recipe lane arrangement type

### RecipeNutriScore
Recipe nutri-score

### RecipeOriginType
Recipe origin type enum

### RecipeSearchSortOption
Sorting options for recipe search query

### RecipeServingType
Recipe servings types

### RecipeShoppableProductSuggestionProposer
Product suggestion proposer

### RecipeShoppableSelectionAction
Selection type enum. Specifies whether ingredient to product mapping was added, changed or removed

### RecipeSortOrderOption
Sorting order for recipe search query

### RecipeSourceType
Recipe source type enum

### RecipeStoryCollectionTheme
Recipe story theme enum

### RecipeStoryCollectionType
Enum thats contains all type of recipeStories

### RecipeStoryItemType
Enum that contains all types of recipeStoryItems

### RecipeVideoSearchSortOption
Sorting options for recipe search query

### RenderType
Render type

### RenditionKey
Recipe image rendition keys enum. See graphQlKey = backendKey

### SearchFacetType
Possible facet types

### SearchIntentType
Possible search products intent types

### SearchPreviouslyBoughtSortType
Possible previously bought sorting types

### SearchProductsConsent
Cookie consent categories supported by search personalization.

### SearchProductsSortType
Possible sorting types for search

### SearchSuggestionType
Suggestion types for search.

### SelfServiceContentErrorTypes
Error types for content module

### SettlementAfterDeliveryType
CostOverviewSettlement Item type

### SettlementRefundMethod
Settlement refund method

### SettlementType
Settlement feedback type

### SortBySort
Order in which the facade should return products.
This is sorted BEFORE splitting into pages

### StampShareErrorReason
Error reason for Sharing link validation and stamp transfer failiure.

### StoreAssortmentAvailabilityStatus
Store assortment availability status

### StoreSpecificAssortmentStatus
Store specific assortment status

### StoresOpeningDayType
The type of opening day

### StoresOpeningHourType
Time frames stores can be open

### StoresType
The type of store

### SubscriptionPaymentMethod
Supported subscription payment method types by bundle processor

### SubscriptionPaymentStatus
Possible payment status after payment

### SubscriptionSavingsCarouselLinkType
The type of the link: DEEPLINK, WEB

### SubscriptionSavingsCarouselTheme
Themes for the premium savings carousel

### SubscriptionType
Currently supported subscriptionTypes

### TargetedOfferAllocationErrorReason
An enumeration that contains the exact error reason that might be returned when attempting to
allocate a targeted offer to a member.

### TargetedOfferStatus
An enumeration containing the status of a targeted offer. This status indicates if and what
action a member has already taken for a particular targeted offer.

### TaxonomyImageRenditions
Image renditions returned by the product image processor
in the form of D<width>x<height>

### ThemeFillStyle
Pantry theme fill style

### TrackAndTraceType
Track and Trace type

### TransactionStatus
Transaction status enum

### TransactionType
Type that defines type of the transaction

### UnsubscribeBlockerReason
A block unsubscribe reason

### UpdateCardResultCode
Specific error code when card has been updated

### WeekDay
Weekday enum

### WeekPlannerItemType
Enum that contains all types of week planner items

### link__Purpose

## INTERFACE Types (23)

### Card
Base Card for Multi Module Cards

- `cardId`: String

### CheckoutChallengeAction
User needs to undertake a challenge to verify the payment

- `sdk`: PaymentSDK

### ContentAccordionBlockBase
Accordion block types

- `type`: ContentAccordionBlockType

### ContentAllerhandeGridlaneComponentBase
Allerhande Flex Page Gridlane Component Base

- `type`: ContentAllerhandeGridlaneComponentType

### ContentAllerhandeSpotlightItemBase
Content Spotlight Data Item Base

- `type`: ContentAllerhandeSpotlightType

### ContentArticleBlockBase
Article content blocks schemas

- `type`: ContentArticleBlockType

### ContentBaseCMSComponent
Common CMS component properties

- `anchorId`: String
- `id`: String
- `tagLevel`: Int
- `componentStatus`: ContentCMSComponentStatus
- `type`: ContentCMSComponentType
- `previewData`: ContentPreviewData

### ContentBaseMobileCMSComponent
Common CMS mobile component properties

- `id`: String
- `type`: ContentBaseMobileCMSComponentType
- `meta`: ContentMobileComponentMetaObject

### ContentBaseTargetedContentCMSDocument
Common targeted content CMS document properties

- `id`: String
- `type`: ContentBaseTargetedContentCMSDocumentType
- `meta`: ContentMobileComponentMetaObject

### ContentFlexPageComponentBase
Flex page component

- `type`: ContentFlexPageComponentType

### ContentFlexPageHeaderBase
Flex page header

- `type`: ContentFlexPageHeaderType

### ContentFormField
Common form field properties

- `type`: ContentFormFieldType
- `media`: ContentMediaBlockBase

### ContentMediaBlockBase
Media base interface

- `type`: ContentMediaBlockType

### ContentSpotlightVisualV2
Spotlight Visual

- `image`: ContentImage

### IDeliverySlot
Common Delivery slot interface

- `date`: String
- `from`: String
- `until`: String
- `isPreferred`: Boolean

### LifestyleCheckQuestionBase
Lifestylecheck Question Interface

- `id`: String
- `question`: String
- `selectedAnswers`: LifestyleCheckSelectedAnswer
- `isLastQuestion`: Boolean

### MutationResult
Result of performing mutation

- `status`: MutationResultStatus
- `errorMessage`: String

### Notification
Notification Base

- `id`: String
- `type`: NotificationType
- `link`: NotificationLink

### PantryColorSet
Pantry color set

- `fillStyle`: ThemeFillStyle
- `primaryColor`: String
- `textColor`: String

### QuestionnaireElement
Interface representing a generic questionnaire element

- `id`: ID
- `type`: String
- `text`: String

### RecipeStoryItemBase
Recipe Story Item Base Interface

- `type`: RecipeStoryItemType

### SelfServiceContentBaseResponse
Self Service Content Base Response

- `query`: String
- `breadcrumbs`: SelfServiceBreadCrumbs
- `buttons`: SelfServiceQuickEntryButtons
- `links`: SelfServiceLinks
- `accordions`: SelfServiceAccordions
- `error`: SelfServiceContentError

### WeekPlannerItemBase
Week Planner Item Base Interface

- `id`: Int
- `type`: WeekPlannerItemType

## UNION Types (16)

### Action
Payment Action types

### CheckoutPaymentAction
Payment Action that needs to be taken.

### ContentAccordionData
Accordion can be two types (Manual and Search)

### ContentAllerhandeImageCollectionLaneItem
Content Image Collection Lane Item

### ContentAllerhandeSpotlightItem
Content Spotlight Item

### ConversationEvent
List of events that can be in a conversation

### CustomerCareChannel
List of all customer care channels

### CustomerCareSettlements
union of all baskets for the mijn verrekeningen page

### EntryPointContentVariant
Union type of all entry point content variants.

### LifestyleCheckQuestion
LifestyleCheck Question

### MessengerConversationEvent
A subset of all possible messenger events

### PaymentProcessingAction
Payment Action

### PaymentsOrderCards
Return union for all cards

### RecipeStoryType
Union for types of recipe stories: product bundle story, recipe story and flex page story

### SubscriptionPayment
The return type of a payment

### TransactionDetails
Payment details

## SCALAR Types (20)

### BigInt
The `BigInt` scalar type represents non-fractional signed whole numeric values.

### Boolean
The `Boolean` scalar type represents `true` or `false`.

### Date
A date (yyyy-MM-dd)

### DateTime
A datetime in ISO8601 format

### EmailAddress
A field whose value conforms to the standard internet email address format as specified in HTML Spec: https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address.

### Float
The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](https://en.wikipedia.org/wiki/IEEE_floating_point).

### ID
The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as `"4"`) or integer (such as `4`) input value will be accepted as an ID.

### Int
The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.

### KeyValueValue
keyvaluevalue scalar

### PageSize
Number (Int) of items from results. Value can be 1 till 100

### PhoneNumber
A phone number

### PostalCode
A field whose value conforms to the standard postal code formats for United States, United Kingdom, Germany, Canada, France, Italy, Australia, Netherlands, Spain, Denmark, Sweden, Belgium, India, Austria, Portugal, Switzerland or Luxembourg.

### PrimitiveTypeScalar
Scalar for any primitive type

### ProductTradeItemResourceIconMeta
Arbitrary resource icon meta data

### String
The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.

### UUID
A UUIDv4

### federation__FieldSet

### federation__Policy

### federation__Scope

### link__Import

