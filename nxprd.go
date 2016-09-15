// Package nxprd is the main package of the GO wrapper for NXP's reader library.
// All NFC functionality is included in this package.
package nxprd

/*
#cgo CFLAGS: -I ${SRCDIR}/nxp/nxprdlib/NxpRdLib/intfs -I ${SRCDIR}/nxp/nxprdlib/NxpRdLib/types -I ${SRCDIR}/nxp/linux/comps/phPlatform/src/Posix -I ${SRCDIR}/nxp/linux/shared -I ${SRCDIR}/nxp/linux/comps/phOsal/src/Posix
#cgo LDFLAGS: -L ${SRCDIR}/nxp/build/linux -lNxpRdLibLinuxPN512 -lrt

//#define DEBUG

#include <phhwConfig.h>
#include <ph_Status.h>
#include <phPlatform.h>
#include <phbalReg.h>

#include <phpalI14443p3a.h>
#include <phpalI14443p4.h>
#include <phpalI14443p3b.h>
#include <phpalI14443p4a.h>
#include <phpalI18092mPI.h>
#include <phpalSli15693.h>
#include <phpalI18000p3m3.h>
#include <phpalMifare.h>
#include <phpalFelica.h>

#include <phalI18000p3m3.h>
#include <phalT1T.h>
#include <phalFelica.h>
#include <phalMfc.h>
#include <phalMful.h>
#include <phacDiscLoop.h>
#include <phKeyStore.h>

#define LISTEN_PHASE_TIME_MS        50

#define NUMBER_OF_KEYENTRIES        2
#define NUMBER_OF_KEYVERSIONPAIRS   2
#define NUMBER_OF_KUCENTRIES        1

#define DATA_BUFFER_LEN             16
#define MFC_BLOCK_DATA_SIZE         16

// SAK codes
#define sak_ul         0x00
#define sak_ulc        0x00
#define sak_mini       0x09
#define sak_mfc_1k     0x08
#define sak_mfc_4k     0x18
#define sak_mfp_2k_sl1 0x08
#define sak_mfp_4k_sl1 0x18
#define sak_mfp_2k_sl2 0x10
#define sak_mfp_4k_sl2 0x11
#define sak_mfp_2k_sl3 0x20
#define sak_mfp_4k_sl3 0x20
#define sak_desfire    0x20
#define sak_jcop       0x28
#define sak_layer4     0x20

// ATQ codes
#define atqa_ul        0x4400
#define atqa_ulc       0x4400
#define atqa_mfc       0x0200
#define atqa_mfp_s     0x0400
#define atqa_mfp_x     0x4200
#define atqa_desfire   0x4403
#define atqa_jcop      0x0400
#define atqa_mini      0x0400
#define atqa_nPA       0x0800

// MIFARE cards
#define mifare_ultralight    0x01
#define mifare_ultralight_c  0x02
#define mifare_classic       0x03
#define mifare_classic_1k    0x04
#define mifare_classic_4k    0x05
#define mifare_plus          0x06
#define mifare_plus_2k_sl1   0x07
#define mifare_plus_4k_sl1   0x08
#define mifare_plus_2k_sl2   0x09
#define mifare_plus_4k_sl2   0x0A
#define mifare_plus_2k_sl3   0x0B
#define mifare_plus_4k_sl3   0x0C
#define mifare_desfire       0x0D
#define jcop                 0x0F
#define mifare_mini          0x10
#define nPA                  0x11

// Used by GO wrapper
typedef enum DiscoverResult {
	DR_FOUND,
	DR_UNKNOWN,
	DR_NOT_FOUND
} DiscoverResult_t;

// Used by GO wrapper
typedef enum TagType
{
	TAT_1 = 1,
	TAT_2,
	TAT_3,
	TAT_4A,
	TAT_P2P,
	TAT_NFC_DEP_4A,
	TAT_UNDEFINED
} TagType_t;

// Used by GO wrapper
typedef enum TechType
{
	TET_A = 1,
	TET_B,
	TET_F,
	TET_V_15693_T5T,
	TET_18000p3m3_EPCGen2,
	TET_UNDEFINED
} TechType_t;

// Used by GO wrapper
typedef struct NFCParams
{
	int		 	sak;
	uint8_t     atq[255];
	uint8_t     atqSize;
	uint8_t 	uid[255];
	uint8_t 	uidSize;
	TechType_t 	techType;
	TagType_t	tagType;
} NFCParams_t;

NFCParams_t nfcParams;

phbalReg_Stub_DataParams_t 			sBalReader;
phPlatform_DataParams_t 			sPlatform;
phhalHw_Nfc_Ic_DataParams_t			sHal_Nfc_Ic;

uint8_t                         	bHalBufferTx[256];
uint8_t                         	bHalBufferRx[256];
uint8_t 							mfulDataBuffer[PHAL_MFUL_READ_BLOCK_LENGTH];

void                            	*pHal;

phpalI14443p3a_Sw_DataParams_t  	spalI14443p3a;
phpalI14443p4a_Sw_DataParams_t  	spalI14443p4a;
phpalI14443p3b_Sw_DataParams_t		spalI14443p3b;
phpalI14443p4_Sw_DataParams_t   	spalI14443p4;
phpalFelica_Sw_DataParams_t     	spalFelica;
phpalI18092mPI_Sw_DataParams_t  	spalI18092mPI;
phpalMifare_Sw_DataParams_t     	spalMifare;

phacDiscLoop_Sw_DataParams_t    	sDiscLoop;
phalMfc_Sw_DataParams_t         	salMfc;
phalMful_Sw_DataParams_t        	salMfu;
phalT1T_Sw_DataParams_t         	alT1T;

#ifndef NXPBUILD__PHHAL_HW_RC523
phpalSli15693_Sw_DataParams_t		spalSli15693;
phalI18000p3m3_Sw_DataParams_t     	salI18000p3m3;
phpalI18000p3m3_Sw_DataParams_t    	spalI18000p3m3;
#endif

phKeyStore_Sw_DataParams_t         	sSwkeyStore;
phKeyStore_Sw_KeyEntry_t           	sKeyEntries[NUMBER_OF_KEYENTRIES];
phKeyStore_Sw_KUCEntry_t           	sKUCEntries[NUMBER_OF_KUCENTRIES];
phKeyStore_Sw_KeyVersionPair_t     	sKeyVersionPairs[NUMBER_OF_KEYVERSIONPAIRS * NUMBER_OF_KEYENTRIES];

uint8_t                            	bDataBuffer[DATA_BUFFER_LEN];

uint8_t                            	bSak;
uint16_t                           	wAtqa;

const uint8_t GI[] = { 0x46,0x66,0x6D,
                       0x01,0x01,0x10,
                       0x03,0x02,0x00,0x01,
                       0x04,0x01,0xF1
                      };

static uint8_t    aData[50];

#ifndef NXPBUILD__PHHAL_HW_RC663
static uint8_t  sens_res[2]     = {0x04, 0x00};
static uint8_t  nfc_id1[3]      = {0xA1, 0xA2, 0xA3};
static uint8_t  sel_res         = 0x40;
static uint8_t  nfc_id3         = 0xFA;
static uint8_t  poll_res[18]    = {0x01, 0xFE, 0xB2, 0xB3, 0xB4, 0xB5,
                                   0xB6, 0xB7, 0xC0, 0xC1, 0xC2, 0xC3,
                                   0xC4, 0xC5, 0xC6, 0xC7, 0x23, 0x45 };
#endif

static uint16_t bSavePollTechCfg  = 0;

uint8_t Key[6]          = {0xFFU, 0xFFU, 0xFFU, 0xFFU, 0xFFU, 0xFFU};
uint8_t Original_Key[6] = {0xFFU, 0xFFU, 0xFFU, 0xFFU, 0xFFU, 0xFFU};

// Used in Discover and Discover_Init functions
uint16_t      wTagsDetected = 0;
uint16_t      wNumberOfTags = 0;
uint16_t      wEntryPoint;

#define DETECT_ERROR 1

#if DETECT_ERROR
    #define DEBUG_ERROR_PRINT(x) x
    #define PRINT_INFO(...) DEBUG_PRINTF(__VA_ARGS__)
#else
    #define DEBUG_ERROR_PRINT(x)
    #define PRINT_INFO(...)
#endif


static void PRINT_BUFF(uint8_t *pBuff, uint8_t num)
{
    uint32_t    i;

    for(i = 0; i < num; i++)
    {
        DEBUG_PRINTF(" %02X",pBuff[i]);
    }
}

static void PrintTagInfo(phacDiscLoop_Sw_DataParams_t *pDataParams, uint16_t wNumberOfTags, uint16_t wTagsDetected)
{
    uint8_t bIndex;
    uint8_t bTagType;

    if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_A))
    {
        if(pDataParams->sTypeATargetInfo.bT1TFlag)
        {
            DEBUG_PRINTF("\tTechnology  : Type A");
            DEBUG_PRINTF ("\n\t\tUID :");
            PRINT_BUFF( pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aUid,
                        pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].bUidSize);
            DEBUG_PRINTF ("\n\t\tSAK : 0x%02x",pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aSak);
            DEBUG_PRINTF ("\n\t\tType: Type 1 Tag\n");
        }
        else
        {
            DEBUG_PRINTF("\tTechnology  : Type A");
            for(bIndex = 0; bIndex < wNumberOfTags; bIndex++)
            {
                DEBUG_PRINTF ("\n\t\tCard: %d",bIndex + 1);
                DEBUG_PRINTF ("\n\t\tUID :");
                PRINT_BUFF( pDataParams->sTypeATargetInfo.aTypeA_I3P3[bIndex].aUid,
                            pDataParams->sTypeATargetInfo.aTypeA_I3P3[bIndex].bUidSize);
                DEBUG_PRINTF ("\n\t\tSAK : 0x%02x",pDataParams->sTypeATargetInfo.aTypeA_I3P3[bIndex].aSak);

                if ((pDataParams->sTypeATargetInfo.aTypeA_I3P3[bIndex].aSak & (uint8_t) ~0xFB) == 0)
                {
                    // Bit b3 is set to zero, [Digital] 4.8.2
                    // Mask out all other bits except for b7 and b6
                    bTagType = (pDataParams->sTypeATargetInfo.aTypeA_I3P3[bIndex].aSak & 0x60);
                    bTagType = bTagType >> 5;

                    switch(bTagType)
                    {
                    case PHAC_DISCLOOP_TYPEA_TYPE2_TAG_CONFIG_MASK:
                        DEBUG_PRINTF ("\n\t\tType: Type 2 Tag\n");
                        break;
                    case PHAC_DISCLOOP_TYPEA_TYPE4A_TAG_CONFIG_MASK:
                        DEBUG_PRINTF ("\n\t\tType: Type 4A Tag\n");
                        break;
                    case PHAC_DISCLOOP_TYPEA_TYPE_NFC_DEP_TAG_CONFIG_MASK:
                        DEBUG_PRINTF ("\n\t\tType: P2P\n");
                        break;
                    case PHAC_DISCLOOP_TYPEA_TYPE_NFC_DEP_TYPE4A_TAG_CONFIG_MASK:
                        DEBUG_PRINTF ("\n\t\tType: Type NFC_DEP and  4A Tag\n");
                        break;
                    default:
                        break;
                    }
                }
            }
        }
    }

    if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_B))
    {
        DEBUG_PRINTF("\tTechnology  : Type B");
        // Loop through all the Type B tags detected and print the Pupi
        for (bIndex = 0; bIndex < wNumberOfTags; bIndex++)
        {
            DEBUG_PRINTF ("\n\t\tCard: %d",bIndex + 1);
            DEBUG_PRINTF ("\n\t\tUID :");
            // PUPI Length is always 4 bytes
            PRINT_BUFF( pDataParams->sTypeBTargetInfo.aTypeB_I3P3[bIndex].aPupi, 0x04);
        }
        DEBUG_PRINTF("\n");
    }

    if( PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F212) ||
        PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F424))
    {
        DEBUG_PRINTF("\tTechnology  : Type F");

        // Loop through all the type F tags and print the IDm
        for (bIndex = 0; bIndex < wNumberOfTags; bIndex++)
        {
            DEBUG_PRINTF ("\n\t\tCard: %d",bIndex + 1);
            DEBUG_PRINTF ("\n\t\tUID :");
            PRINT_BUFF( pDataParams->sTypeFTargetInfo.aTypeFTag[bIndex].aIDmPMm,
                        PHAC_DISCLOOP_FELICA_IDM_LENGTH );
            if ((pDataParams->sTypeFTargetInfo.aTypeFTag[bIndex].aIDmPMm[0] == 0x01) &&
                (pDataParams->sTypeFTargetInfo.aTypeFTag[bIndex].aIDmPMm[1] == 0xFE))
            {
                // This is Type F tag with P2P capabilities
                DEBUG_PRINTF ("\n\t\tType: P2P");
            }
            else
            {
                // This is Type F T3T tag
                DEBUG_PRINTF ("\n\t\tType: Type 3 Tag");
            }

            if(pDataParams->sTypeFTargetInfo.aTypeFTag[bIndex].bBaud != PHAC_DISCLOOP_CON_BITR_212)
            {
                DEBUG_PRINTF ("\n\t\tBit Rate: 424\n");
            }
            else
            {
                DEBUG_PRINTF ("\n\t\tBit Rate: 212\n");
            }
        }
    }
#ifndef NXPBUILD__PHHAL_HW_RC523
    if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_V))
    {
        DEBUG_PRINTF("\tTechnology  : Type V / ISO 15693 / T5T");
        // Loop through all the Type V tags detected and print the UIDs
        for (bIndex = 0; bIndex < wNumberOfTags; bIndex++)
        {
            DEBUG_PRINTF ("\n\t\tCard: %d",bIndex + 1);
            DEBUG_PRINTF ("\n\t\tUID :");
            PRINT_BUFF( pDataParams->sTypeVTargetInfo.aTypeV[bIndex].aUid, 0x08);
        }
        DEBUG_PRINTF("\n");
    }

    if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_18000P3M3))
    {
        DEBUG_PRINTF("\tTechnology  : ISO 18000p3m3 / EPC Gen2");
        // Loop through all the 18000p3m3 tags detected and print the UII
        for (bIndex = 0; bIndex < wNumberOfTags; bIndex++)
        {
            DEBUG_PRINTF("\n\t\tCard: %d",bIndex + 1);
            DEBUG_PRINTF("\n\t\tUII :");
            PRINT_BUFF(
                pDataParams->sI18000p3m3TargetInfo.aI18000p3m3[bIndex].aUii,
                (pDataParams->sI18000p3m3TargetInfo.aI18000p3m3[bIndex].wUiiLength / 8));
        }
        DEBUG_PRINTF("\n");
    }
#endif
}

#if DETECT_ERROR
static void PrintErrorInfo(phStatus_t wStatus)
{
    DEBUG_PRINTF("\n ErrorInfo Comp:");

    switch(wStatus & 0xFF00)
    {
    case PH_COMP_BAL:
        DEBUG_PRINTF("\t PH_COMP_BAL");
        break;
    case PH_COMP_HAL:
        DEBUG_PRINTF("\t PH_COMP_HAL");
        break;
    case PH_COMP_PAL_ISO14443P3A:
        DEBUG_PRINTF("\t PH_COMP_PAL_ISO14443P3A");
        break;
    case PH_COMP_PAL_ISO14443P3B:
        DEBUG_PRINTF("\t PH_COMP_PAL_ISO14443P3B");
        break;
    case PH_COMP_PAL_ISO14443P4A:
        DEBUG_PRINTF("\t PH_COMP_PAL_ISO14443P4A");
        break;
    case PH_COMP_PAL_ISO14443P4:
        DEBUG_PRINTF("\t PH_COMP_PAL_ISO14443P4");
        break;
    case PH_COMP_PAL_FELICA:
        DEBUG_PRINTF("\t PH_COMP_PAL_FELICA");
        break;
    case PH_COMP_PAL_EPCUID:
        DEBUG_PRINTF("\t PH_COMP_PAL_EPCUID");
        break;
    case PH_COMP_PAL_SLI15693:
        DEBUG_PRINTF("\t PH_COMP_PAL_SLI15693");
        break;
    case PH_COMP_PAL_I18000P3M3:
        DEBUG_PRINTF("\t PH_COMP_PAL_I18000P3M3");
        break;
    case PH_COMP_PAL_I18092MPI:
        DEBUG_PRINTF("\t PH_COMP_PAL_I18092MPI");
        break;
    case PH_COMP_PAL_I18092MT:
        DEBUG_PRINTF("\t PH_COMP_PAL_I18092MT");
        break;
    case PH_COMP_PAL_GENERALTARGET:
        DEBUG_PRINTF("\t PH_COMP_PAL_GENERALTARGET");
        break;
    case PH_COMP_PAL_I14443P4MC:
        DEBUG_PRINTF("\t PH_COMP_PAL_I14443P4MC");
        break;
    case PH_COMP_AC_DISCLOOP:
        DEBUG_PRINTF("\t PH_COMP_AC_DISCLOOP");
        break;
    case PH_COMP_OSAL:
        DEBUG_PRINTF("\t PH_COMP_PAL_I14443P4MC");
        break;
    default:
            DEBUG_PRINTF("\t 0x%x",(wStatus & PH_COMPID_MASK));
            break;
    }

    DEBUG_PRINTF("\t type:");

    switch(wStatus & PH_ERR_MASK)
    {
    case PH_ERR_SUCCESS_INCOMPLETE_BYTE:
        DEBUG_PRINTF("\t PH_ERR_SUCCESS_INCOMPLETE_BYTE");
        break;
    case PH_ERR_IO_TIMEOUT:
        DEBUG_PRINTF("\t PH_ERR_IO_TIMEOUT");
        break;
    case PH_ERR_INTEGRITY_ERROR:
        DEBUG_PRINTF("\t PH_ERR_INTEGRITY_ERROR");
        break;
    case PH_ERR_COLLISION_ERROR:
        DEBUG_PRINTF("\t PH_ERR_COLLISION_ERROR");
        break;
    case PH_ERR_BUFFER_OVERFLOW:
        DEBUG_PRINTF("\t PH_ERR_BUFFER_OVERFLOW");
        break;
    case PH_ERR_FRAMING_ERROR:
        DEBUG_PRINTF("\t PH_ERR_FRAMING_ERROR");
        break;
    case PH_ERR_PROTOCOL_ERROR:
        DEBUG_PRINTF("\t PH_ERR_PROTOCOL_ERROR");
        break;
    case PH_ERR_RF_ERROR:
        DEBUG_PRINTF("\t PH_ERR_RF_ERROR");
        break;
    case PH_ERR_EXT_RF_ERROR:
        DEBUG_PRINTF("\t PH_ERR_EXT_RF_ERROR");
        break;
    case PH_ERR_NOISE_ERROR:
        DEBUG_PRINTF("\t PH_ERR_NOISE_ERROR");
        break;
    case PH_ERR_ABORTED:
        DEBUG_PRINTF("\t PH_ERR_ABORTED");
        break;
    //case PH_ERR_RF_TURNOFF:
    //    DEBUG_PRINTF("\t PH_ERR_RF_TURNOFF");
    //    break;
    case PH_ERR_INTERNAL_ERROR:
        DEBUG_PRINTF("\t PH_ERR_INTERNAL_ERROR");
        break;
    case PH_ERR_INVALID_DATA_PARAMS:
        DEBUG_PRINTF("\t PH_ERR_INVALID_DATA_PARAMS");
        break;
    case PH_ERR_INVALID_PARAMETER:
        DEBUG_PRINTF("\t PH_ERR_INVALID_PARAMETER");
        break;
    case PH_ERR_PARAMETER_OVERFLOW:
        DEBUG_PRINTF("\t PH_ERR_PARAMETER_OVERFLOW");
        break;
    case PH_ERR_UNSUPPORTED_PARAMETER:
        DEBUG_PRINTF("\t PH_ERR_UNSUPPORTED_PARAMETER");
        break;
    case PH_ERR_OSAL_ERROR:
        DEBUG_PRINTF("\t PH_ERR_OSAL_ERROR");
        break;
    case PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED:
        DEBUG_PRINTF("\t PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED");
        break;
    case PHAC_DISCLOOP_COLLISION_PENDING:
        DEBUG_PRINTF("\t PHAC_DISCLOOP_COLLISION_PENDING");
        break;
    default:
        DEBUG_PRINTF("\t 0x%x",(wStatus & PH_ERR_MASK));
        break;
    }
}
#endif

// Print technology being resolved
void PrintTechnology(uint8_t TechType)
{
    switch(TechType)
    {
    case PHAC_DISCLOOP_POS_BIT_MASK_A:
        DEBUG_PRINTF ("\tResolving Type A... \n");
        break;

    case PHAC_DISCLOOP_POS_BIT_MASK_B:
        DEBUG_PRINTF ("\tResolving Type B... \n");
        break;

    case PHAC_DISCLOOP_POS_BIT_MASK_F212:
        DEBUG_PRINTF ("\tResolving Type F with baud rate 212... \n");
        break;

    case PHAC_DISCLOOP_POS_BIT_MASK_F424:
        DEBUG_PRINTF ("\tResolving Type F with baud rate 424... \n");
        break;

    case PHAC_DISCLOOP_POS_BIT_MASK_V:
        DEBUG_PRINTF ("\tResolving Type V... \n");
        break;

    default:
        break;
    }
}

phStatus_t NfcRdLibInit(void) {
	phStatus_t status;

	status = phbalReg_Stub_Init(
		&sBalReader,
		sizeof(phbalReg_Stub_DataParams_t));
	CHECK_STATUS(status);

	status = phPlatform_Init(&sPlatform);
	CHECK_SUCCESS(status);

	status = phOsal_Event_Init();
	CHECK_STATUS(status);

	Set_Interrupt();

	#ifdef NXPBUILD__PHHAL_HW_PN5180
	status = phbalReg_SetConfig(
		&sBalReader,
		PHBAL_REG_CONFIG_HAL_HW_TYPE,
		PHBAL_REG_HAL_HW_PN5180);
	#endif
	#ifdef NXPBUILD__PHHAL_HW_RC523
	status = phbalReg_SetConfig(
		&sBalReader,
		PHBAL_REG_CONFIG_HAL_HW_TYPE,
		PHBAL_REG_HAL_HW_RC523);
	#endif
	#ifdef NXPBUILD__PHHAL_HW_RC663
	status = phbalReg_SetConfig(
		&sBalReader,
		PHBAL_REG_CONFIG_HAL_HW_TYPE,
		PHBAL_REG_HAL_HW_RC663);
	#endif
	CHECK_STATUS(status);

	status = phbalReg_SetPort(
		&sBalReader,
		SPI_CONFIG);
	CHECK_STATUS(status);

	status = phbalReg_OpenPort(&sBalReader);
	CHECK_STATUS(status);

	status = phhalHw_Nfc_IC_Init(
		&sHal_Nfc_Ic,
		sizeof(phhalHw_Nfc_Ic_DataParams_t),
		&sBalReader,
		0,
		bHalBufferTx,
		sizeof(bHalBufferTx),
		bHalBufferRx,
		sizeof(bHalBufferRx));

	sHal_Nfc_Ic.sHal.bBalConnectionType = PHHAL_HW_BAL_CONNECTION_SPI;

	Configure_Device(&sHal_Nfc_Ic);

	pHal = &sHal_Nfc_Ic.sHal;

	status = phpalI14443p3a_Sw_Init(
		&spalI14443p3a,
		sizeof(phpalI14443p3a_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	status = phpalI14443p4a_Sw_Init(
		&spalI14443p4a,
		sizeof(phpalI14443p4a_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	status = phpalI14443p4_Sw_Init(
		&spalI14443p4,
		sizeof(phpalI14443p4_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	status = phpalI14443p3b_Sw_Init(
		&spalI14443p3b,
		sizeof(spalI14443p3b),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	status = phpalFelica_Sw_Init(
		&spalFelica,
		sizeof(phpalFelica_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	status = phpalI18092mPI_Sw_Init(
		&spalI18092mPI,
		sizeof(phpalI18092mPI_Sw_DataParams_t),
		pHal);
	CHECK_STATUS(status);

	status = phpalMifare_Sw_Init(
		&spalMifare,
		sizeof(phpalMifare_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal,
		&spalI14443p4);
	CHECK_STATUS(status);

	#ifndef NXPBUILD__PHHAL_HW_RC523
	status = phpalI18000p3m3_Sw_Init(
		&spalI18000p3m3,
		sizeof(phpalI18000p3m3_Sw_DataParams_t),
		pHal);
	CHECK_STATUS(status);

	status = phalI18000p3m3_Sw_Init(
		&salI18000p3m3,
		sizeof(phalI18000p3m3_Sw_DataParams_t),
		&spalI18000p3m3);
	CHECK_STATUS(status);

	status = phpalSli15693_Sw_Init(
		&spalSli15693,
		sizeof(phpalSli15693_Sw_DataParams_t),
		pHal);
	CHECK_STATUS(status);
	#endif

	status = phalT1T_Sw_Init(
		&alT1T,
		sizeof(phalT1T_Sw_DataParams_t),
		&spalI14443p3a);
	CHECK_STATUS(status);

	status = phpalMifare_Sw_Init(&spalMifare, sizeof(phpalMifare_Sw_DataParams_t), &sHal_Nfc_Ic.sHal, NULL);
	CHECK_STATUS(status);

	// Initialize the keystore component
	// Not used at the moment
	status = phKeyStore_Sw_Init(
								&sSwkeyStore,
								sizeof(phKeyStore_Sw_DataParams_t),
								&sKeyEntries[0],
								NUMBER_OF_KEYENTRIES,
								&sKeyVersionPairs[0],
								NUMBER_OF_KEYVERSIONPAIRS,
								&sKUCEntries[0],
								NUMBER_OF_KUCENTRIES);
	CHECK_SUCCESS(status);

	status = phalMful_Sw_Init(&salMfu, sizeof(phalMful_Sw_DataParams_t), &spalMifare, NULL, NULL, NULL);
	CHECK_STATUS(status);

	status = phacDiscLoop_Sw_Init(
		&sDiscLoop,
		sizeof(phacDiscLoop_Sw_DataParams_t),
		&sHal_Nfc_Ic.sHal);
	CHECK_STATUS(status);

	#ifdef NXPBUILD__PHHAL_HW_RC523
	status = phhalHw_Rc523_SetListenParameters(
		&sHal_Nfc_Ic.sHal,
		&sens_res[0],
		&nfc_id1[0],
		sel_res,
		&poll_res[0],
		nfc_id3);
	CHECK_SUCCESS(status);
	#endif

	#ifdef NXPBUILD__PHHAL_HW_PN5180
	status = phhalHw_Pn5180_SetListenParameters(
		&sHal_Nfc_Ic.sHal,
		&sens_res[0],
		&nfc_id1[0],
		sel_res,
		&poll_res[0],
		nfc_id3);
	CHECK_SUCCESS(status);
	#endif

	sDiscLoop.pPal1443p3aDataParams   = &spalI14443p3a;
	sDiscLoop.pPal1443p3bDataParams   = &spalI14443p3b;
	sDiscLoop.pPal1443p4aDataParams   = &spalI14443p4a;
	sDiscLoop.pPal14443p4DataParams   = &spalI14443p4;
	#ifndef NXPBUILD__PHHAL_HW_RC523
	sDiscLoop.pPal18000p3m3DataParams = &spalI18000p3m3;
	sDiscLoop.pAl18000p3m3DataParams  = &salI18000p3m3;
	sDiscLoop.pPalSli15693DataParams  = &spalSli15693;
	#endif
	sDiscLoop.pPal18092mPIDataParams  = &spalI18092mPI;
	sDiscLoop.pPalFelicaDataParams    = &spalFelica;
	sDiscLoop.pAlT1TDataParams        = &alT1T;
	sDiscLoop.pHalDataParams          = &sHal_Nfc_Ic.sHal;

	sDiscLoop.sTypeATargetInfo.sTypeA_P2P.pGi       = (uint8_t *)GI;
	sDiscLoop.sTypeATargetInfo.sTypeA_P2P.bGiLength = sizeof(GI);

	sDiscLoop.sTypeFTargetInfo.sTypeF_P2P.pGi       = (uint8_t *)GI;
	sDiscLoop.sTypeFTargetInfo.sTypeF_P2P.bGiLength = sizeof(GI);

	sDiscLoop.sTypeATargetInfo.sTypeA_P2P.pAtrRes   = aData;

	sDiscLoop.sTypeFTargetInfo.sTypeF_P2P.pAtrRes   = aData;

	sDiscLoop.sTypeATargetInfo.sTypeA_I3P4.pAts     = aData;

	return PH_ERR_SUCCESS;
}

#ifdef NXPBUILD__PHHAL_HW_RC663
phStatus_t ConfigureLPCD(void)
{
    phStatus_t status;
    uint8_t bValueI;
    uint8_t bValueQ;

    status = phhalHw_Rc663_Cmd_Lpcd_GetConfig(pHal, &bValueI, &bValueQ);
    CHECK_STATUS(status);

    status = phhalHw_Rc663_Cmd_Lpcd_SetConfig(
        pHal,
        PHHAL_HW_RC663_CMD_LPCD_MODE_POWERDOWN,
        bValueI,
        bValueQ,
        1,
        100);

    status = phacDiscLoop_SetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_ENABLE_LPCD, PH_ON);
    CHECK_STATUS(status);

    return status;
}
#endif

#ifdef NXPBUILD__PHHAL_HW_PN5180
phStatus_t ConfigureLPCD(void)
{
    phStatus_t status;
    uint16_t wConfig = PHHAL_HW_CONFIG_LPCD_REF;
    uint16_t wValue;

    status = phhalHw_Pn5180_Int_LPCD_GetConfig(pHal, wConfig, &wValue);
    CHECK_STATUS(status);

    wValue = PHHAL_HW_PN5180_LPCD_MODE_POWERDOWN;
    wConfig = PHHAL_HW_CONFIG_LPCD_MODE;

    status = phhalHw_Pn5180_Int_LPCD_SetConfig(
        pHal,
        wConfig,
        wValue
	);

    status = phacDiscLoop_SetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_ENABLE_LPCD, PH_ON);
    CHECK_STATUS(status);

    return status;
}
#endif

void Discover_Init() {
	void 		  *pDataParams 	= &sDiscLoop;
	phStatus_t    status 		= PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED;

	status = phacDiscLoop_GetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_PAS_POLL_TECH_CFG, &bSavePollTechCfg);
	CHECK_STATUS(status);

#ifdef NXPBUILD__PHHAL_HW_RC523
	wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#else
	wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_POLL;
	status = PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED;
#endif
}

DiscoverResult_t Discover()
{
    phStatus_t    status = PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED;
	void 		  *pDataParams 	= &sDiscLoop;
	uint8_t		  bIndex;
	uint16_t      wValue;

#ifndef NXPBUILD__PHHAL_HW_RC523
    if((status & PH_ERR_MASK) == PHAC_DISCLOOP_LPCD_NO_TECH_DETECTED)
    {
        status = ConfigureLPCD();
        CHECK_STATUS(status);
    }
#endif
    status = phacDiscLoop_SetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_NEXT_POLL_STATE, PHAC_DISCLOOP_POLL_STATE_DETECTION);
    CHECK_STATUS(status);

    status = phacDiscLoop_SetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_PAS_POLL_TECH_CFG, bSavePollTechCfg);
    CHECK_STATUS(status);

    status = phhalHw_FieldOff(pHal);
    CHECK_STATUS(status);

    status = phacDiscLoop_Run(pDataParams, wEntryPoint);
    if((status & PH_ERR_MASK) == PHAC_DISCLOOP_MULTI_TECH_DETECTED)
    {
        DEBUG_PRINTF (" \n Multiple technology detected: \n");

        status = phacDiscLoop_GetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
        CHECK_STATUS(status);

        if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_A))
        {
            DEBUG_PRINTF (" \tType A detected... \n");
        }
        if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_B))
        {
            DEBUG_PRINTF (" \tType B detected... \n");
        }
        if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F212))
        {
            DEBUG_PRINTF (" \tType F detected with baud rate 212... \n");
        }
        if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F424))
        {
            DEBUG_PRINTF (" \tType F detected with baud rate 424... \n");
        }
#ifndef NXPBUILD__PHHAL_HW_RC523
        if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_V))
        {
            DEBUG_PRINTF(" \tType V / ISO 15693 / T5T detected... \n");
        }
#endif
        for(bIndex = 0; bIndex < PHAC_DISCLOOP_PASS_POLL_MAX_TECHS_SUPPORTED; bIndex++)
        {
            if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, (1 << bIndex)))
            {
                status = phacDiscLoop_SetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_PAS_POLL_TECH_CFG, (1 << bIndex));
                CHECK_STATUS(status);
                break;
            }
        }

        //Print the technology resolved
#ifdef DEBUG
        PrintTechnology((1 << bIndex));
#endif
        // Set Discovery Poll State to collision resolution
        status = phacDiscLoop_SetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_NEXT_POLL_STATE, PHAC_DISCLOOP_POLL_STATE_COLLISION_RESOLUTION);
        CHECK_STATUS(status);

        // Restart discovery loop in poll mode from collision resolution phase
        status = phacDiscLoop_Run(pDataParams, wEntryPoint);
    }

    if((status & PH_ERR_MASK) == PHAC_DISCLOOP_MULTI_DEVICES_RESOLVED)
    {
        // Get Detected Technology Type
        status = phacDiscLoop_GetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
        CHECK_STATUS(status);

        // Get number of tags detected
        status = phacDiscLoop_GetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_NR_TAGS_FOUND, &wNumberOfTags);
        CHECK_STATUS(status);

        DEBUG_PRINTF (" \n Multiple cards resolved: %d cards \n",wNumberOfTags);
#ifdef DEBUG
        PrintTagInfo(pDataParams, wNumberOfTags, wTagsDetected);
#endif
        if(wNumberOfTags > 1)
        {
            // Get 1st Detected Technology and Activate device at index 0
            for(bIndex = 0; bIndex < PHAC_DISCLOOP_PASS_POLL_MAX_TECHS_SUPPORTED; bIndex++)
            {
                if(PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, (1 << bIndex)))
                {
                    DEBUG_PRINTF("\t Activating one card...\n");
                    status = phacDiscLoop_ActivateCard(pDataParams, bIndex, 0);
                    break;
                }
            }

            if(((status & PH_ERR_MASK) == PHAC_DISCLOOP_DEVICE_ACTIVATED) ||
               ((status & PH_ERR_MASK) == PHAC_DISCLOOP_PASSIVE_TARGET_ACTIVATED) ||
               ((status & PH_ERR_MASK) == PHAC_DISCLOOP_MERGED_SEL_RES_FOUND))
            {
                // Get Detected Technology Type
                status = phacDiscLoop_GetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
                CHECK_STATUS(status);
#ifdef DEBUG
                PrintTagInfo(pDataParams, 0x01, wTagsDetected);
#endif
#ifdef NXPBUILD__PHHAL_HW_RC523
                wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
				// Change to DR_FOUND? How to reproduce multiple cards resolved?
				return DR_UNKNOWN;
			}
            else
            {
                PRINT_INFO("\t\tCard activation failed...\n");
            }
        }
    }
    else if (((status & PH_ERR_MASK) == PHAC_DISCLOOP_NO_TECH_DETECTED) ||
            ((status & PH_ERR_MASK) == PHAC_DISCLOOP_NO_DEVICE_RESOLVED))
    {
        // Switch to LISTEN Mode
#ifdef NXPBUILD__PHHAL_HW_RC523
         wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_EXTERNAL_RFON)
    {
        // If external RF is detected during POLL, return back so that the application
        // can restart the loop in LISTEN mode
#ifdef NXPBUILD__PHHAL_HW_RC523
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_EXTERNAL_RFOFF)
    {
        // Enters here if in the target/card mode and external RF is not available
        // Wait for LISTEN timeout till an external RF is detected.
        // Application may choose to go into standby at this point.
#ifdef NXPBUILD__PHHAL_HW_RC523
        status = phOsal_Event_Consume(E_PH_OSAL_EVT_RF, E_PH_OSAL_EVT_SRC_ISR);
        CHECK_STATUS(status);

        status = phhalHw_SetConfig(pHal, PHHAL_HW_CONFIG_RFON_INTERRUPT, PH_ON);
        CHECK_STATUS(status);

        status = phOsal_Event_WaitAny(E_PH_OSAL_EVT_RF, LISTEN_PHASE_TIME_MS , NULL);
        if((status & PH_ERR_MASK) == PH_ERR_IO_TIMEOUT)
        {
			// With RC523 board we're always landing here after the first call
			// to this function. Listen mode not really supported by RC523?
			wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_POLL;
        }
        else
        {
            wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
        }
#endif

    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_ACTIVATED_BY_PEER)
    {
        DEBUG_PRINTF (" \n Device activated in listen mode... \n");
#ifdef NXPBUILD__PHHAL_HW_RC523
        // On successful activated by Peer, switch to LISTEN mode
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_ACTIVE_TARGET_ACTIVATED)
    {
        DEBUG_PRINTF (" \n Active target detected... \n");
#ifdef NXPBUILD__PHHAL_HW_RC523
        // Target Activated successful, switch to LISTEN mode
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_PASSIVE_TARGET_ACTIVATED)
    {
        DEBUG_PRINTF (" \n Passive target detected... \n");

        // Get Detected Technology Type
        status = phacDiscLoop_GetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
        CHECK_STATUS(status);

#ifdef DEBUG
        PrintTagInfo(pDataParams, 1, wTagsDetected);
#endif
#ifdef NXPBUILD__PHHAL_HW_RC523
        // Target Activated successful, switch to LISTEN mode
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_MERGED_SEL_RES_FOUND)
    {
        DEBUG_PRINTF (" \n Device having T4T and NFC-DEP support detected... \n");

        // Get Detected Technology Type
        status = phacDiscLoop_GetConfig(&sDiscLoop, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
        CHECK_STATUS(status);

#ifdef DEBUG
        PrintTagInfo(pDataParams, 1, wTagsDetected);
#endif
#ifdef NXPBUILD__PHHAL_HW_RC523
        // Switch to LISTEN mode
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
    }
    else if((status & PH_ERR_MASK) == PHAC_DISCLOOP_DEVICE_ACTIVATED)
    {
        DEBUG_PRINTF (" \n Card detected and activated successfully... \n");
        status = phacDiscLoop_GetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_NR_TAGS_FOUND, &wNumberOfTags);
        CHECK_STATUS(status);

        // Get Detected Technology Type
        status = phacDiscLoop_GetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_TECH_DETECTED, &wTagsDetected);
        CHECK_STATUS(status);

#ifdef DEBUG
        PrintTagInfo(pDataParams, wNumberOfTags, wTagsDetected);
#endif
#ifdef NXPBUILD__PHHAL_HW_RC523
        // On successful activation, switch to LISTEN mode
        wEntryPoint = PHAC_DISCLOOP_ENTRY_POINT_LISTEN;
#endif
		return DR_FOUND;
    }
    else
    {
        if((status & PH_ERR_MASK) == PHAC_DISCLOOP_FAILURE)
        {
            status = phacDiscLoop_GetConfig(pDataParams, PHAC_DISCLOOP_CONFIG_ADDITIONAL_INFO, &wValue);
            CHECK_STATUS(status);
            DEBUG_ERROR_PRINT(PrintErrorInfo(wValue));
        }
        else
        {
            DEBUG_ERROR_PRINT(PrintErrorInfo(status));
        }
    }
	return DR_NOT_FOUND;
}

//
// Functions used exclusively by the GO wrapper
//

uint32_t DetectMifare(void)
{
	if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_A))
	{
		// Check for MIFARE UL
		if (sDiscLoop.sTypeATargetInfo.aTypeA_I3P3[0].aSak == 0)
		{
			return mifare_ultralight;
		}
	}
	return 0;
}

void NFCParams_Retrieve(void)
{
	phacDiscLoop_Sw_DataParams_t *pDataParams = &sDiscLoop;

	nfcParams.sak = -1;
	nfcParams.atqSize = 0;
	nfcParams.uidSize = 0;
	nfcParams.techType = TET_UNDEFINED;
	nfcParams.tagType = TAT_UNDEFINED;

	if (wNumberOfTags == 0) {
		return;
	}

	if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_A))
	{
		if(pDataParams->sTypeATargetInfo.bT1TFlag)
		{
			nfcParams.techType = TET_A;
			nfcParams.tagType = TAT_1;
			nfcParams.uidSize = pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].bUidSize;
			nfcParams.atqSize = 2;
			memcpy(nfcParams.uid, pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aUid, nfcParams.uidSize);
			nfcParams.sak = pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aSak;
			memcpy(nfcParams.atq, pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aAtqa, nfcParams.atqSize);
		}
		else
		{
			nfcParams.techType = TET_A;
			nfcParams.atqSize = 2;
			// Only use the first tag at index 0
			nfcParams.uidSize = pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].bUidSize;
			memcpy(nfcParams.uid, pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aUid, nfcParams.uidSize);
			nfcParams.sak = pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aSak;
			memcpy(nfcParams.atq, pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aAtqa, nfcParams.atqSize);

			if ((pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aSak & (uint8_t) ~0xFB) == 0)
			{
				// Bit b3 is set to zero, [Digital] 4.8.2
				// Mask out all other bits except for b7 and b6
				uint8_t bTagType = (pDataParams->sTypeATargetInfo.aTypeA_I3P3[0].aSak & 0x60);
				bTagType = bTagType >> 5;

				switch(bTagType)
				{
				case PHAC_DISCLOOP_TYPEA_TYPE2_TAG_CONFIG_MASK:
					nfcParams.tagType = TAT_2;
					break;
				case PHAC_DISCLOOP_TYPEA_TYPE4A_TAG_CONFIG_MASK:
					nfcParams.tagType = TAT_4A;
					break;
				case PHAC_DISCLOOP_TYPEA_TYPE_NFC_DEP_TAG_CONFIG_MASK:
					nfcParams.tagType = TAT_P2P;
					break;
				case PHAC_DISCLOOP_TYPEA_TYPE_NFC_DEP_TYPE4A_TAG_CONFIG_MASK:
					nfcParams.tagType = TAT_NFC_DEP_4A;
					break;
				default:
					nfcParams.tagType = TAT_UNDEFINED;
					break;
				}
			}
		}
	}

	if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_B))
	{
		nfcParams.techType = TET_B;
		// Only use the first type B tag at index 0
		// PUPI Length is always 4 bytes
		nfcParams.uidSize = 4;
		nfcParams.atqSize = pDataParams->sTypeBTargetInfo.aTypeB_I3P3[0].bAtqBLength;
		memcpy(nfcParams.uid, pDataParams->sTypeBTargetInfo.aTypeB_I3P3[0].aPupi, nfcParams.uidSize);
		memcpy(nfcParams.atq, pDataParams->sTypeBTargetInfo.aTypeB_I3P3[0].aAtqB, nfcParams.atqSize);

	}

	if( PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F212) ||
		PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_F424))
	{
		nfcParams.techType = TET_F;

		// Only use the first type F tag at index 0
		nfcParams.uidSize = PHAC_DISCLOOP_FELICA_IDM_LENGTH;
		memcpy(nfcParams.uid, pDataParams->sTypeFTargetInfo.aTypeFTag[0].aIDmPMm, nfcParams.uidSize);

		if ((pDataParams->sTypeFTargetInfo.aTypeFTag[0].aIDmPMm[0] == 0x01) &&
			(pDataParams->sTypeFTargetInfo.aTypeFTag[0].aIDmPMm[1] == 0xFE))
		{
			// This is Type F tag with P2P capabilities
			nfcParams.tagType = TAT_P2P;
		}
		else
		{
			// This is Type F T3T tag
			nfcParams.tagType = TAT_3;
		}
	}
	#ifndef NXPBUILD__PHHAL_HW_RC523
	if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_V))
	{
		nfcParams.techType = TET_V_15693_T5T;
		// Only use the first type V tag at index 0
		nfcParams.uidSize = 8;
		memcpy(nfcParams.uid, pDataParams->sTypeVTargetInfo.aTypeV[0].aUid, nfcParams.uidSize);
	}

	if (PHAC_DISCLOOP_CHECK_ANDMASK(wTagsDetected, PHAC_DISCLOOP_POS_BIT_MASK_18000P3M3))
	{
		nfcParams.techType = TET_18000p3m3_EPCGen2;
		// Only use the first 18000p3m3 tag at index 0
		nfcParams.uidSize = pDataParams->sI18000p3m3TargetInfo.aI18000p3m3[0].wUiiLength / 8;
		memcpy(nfcParams.uid, pDataParams->sI18000p3m3TargetInfo.aI18000p3m3[0].aUii, nfcParams.uidSize);
	}
#endif
}

phStatus_t MifareUL_Read_Block(uint8_t blockIdx)
{
	phStatus_t status;
	void *buffer = &mfulDataBuffer;

	status = phalMful_Read(&salMfu, blockIdx, buffer);
	CHECK_SUCCESS(status);
}

phStatus_t MifareUL_Write_Block(uint8_t blockIdx, void *data)
{
	phStatus_t status;

	status = phalMful_Write(&salMfu, blockIdx, data);
	CHECK_SUCCESS(status);
}
*/
import "C"
import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"math"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// Used for synchronizing public methods
var libLock sync.Mutex

// Reader defines an interface for reading data from a nfc capable card/tag.
type Reader interface {
	// ReadBlock reads a block of bytes from a nfc card/tag. idx defines the
	// block number or index. The memory layout, i.e. the number and the size of
	// blocks may differ between the various tags and cards.
	// If data can't be read, a NxpError is returned along with a nil slice.
	// If data could be read, error will be always nil.
	ReadBlock(idx int) ([]byte, error)
	//Returns the entire contents of a card or tag as a NDEF struct
	ReadNdef() (Ndef, error)
}

// Writer defines an interface for writing data to a nfc capable card/tag.
type Writer interface {
	// WriteBlock writes a block (slice) of bytes to a nfc card/tag.
	// The position of the block is defined by idx and may vary between
	// different cards and tags. Also the length of the data slice depends on
	// the card/tag being used.
	// If the data can't be written, a NxpError is returned. If no error occurred,
	// nil is returned.
	WriteBlock(idx int, data []byte) error
	//Writes the NDEF struct to the tag or card
	WriteNdef(ndef *Ndef) error
	//Writes the string of a particular language to the card or tag
	WriteString(payload string, language language) error
}

// MifareULReader implements the Reader interface for the Mifare Ultralight
// card.
type MifareULReader struct {
}

const (
	// MifareULBlockLength defines the block length in bytes of a Mifare
	// Ultralight card, i.e. a block (page) of a Mifare Ultralight contains 4
	// bytes.
	MifareULBlockLength = 4
)

// ReadBlock implements reading a block of bytes for a Mifare Ultralight card.
// idx defines the block to be read. See also the Reader interface.
func (r *MifareULReader) ReadBlock(idx int) ([]byte, error) {
	libLock.Lock()
	defer libLock.Unlock()

	// The actual read operation is carried out by a C function (defined in the
	// comments for the pseudo C package).
	// which stores the data in a global C array.
	status := C.MifareUL_Read_Block(C.uint8_t(idx))
	if status != C.PH_ERR_SUCCESS {
		return nil, createLibErr(int(status))
	}
	buffer := make([]byte, MifareULBlockLength)
	for i := 0; i < MifareULBlockLength; i++ {
		// Get the bytes from the global C array
		buffer[i] = byte(C.mfulDataBuffer[i])
	}
	return buffer, nil
}

//Index of bytes used to convert the read bytes to a NDEF struct.
const (
	messageOffsetIndex     = 0
	messageTypeLengthIndex = 1
	payloadLengthIndex     = 2
	payloadTypeStartIndex  = 3
	idStartIndex           = 5
	UserBlockStartIndex    = 4
	UserBlockEndIndex      = 15
	MB                     = 128
	ME                     = 64
	CF                     = 32
	SR                     = 16
	IL                     = 8
)

//A struct which represents the NDEF message
type Ndef struct {
	MaxSize            int
	IsReadOnly         bool
	TotalMessageLength int
	NdefData           []NdefRecord
}

//A struct which represents each NDEF message
type NdefRecord struct {
	IsStartRecord          bool
	IsEndRecord            bool
	ChunkFlag              bool
	IsShortRecord          bool
	IsIdLengthFieldPresent bool
	PayloadLength          int
	Payload                []byte
	Tnf                    tnf
	Type                   string
	RecordNumber           int
	TypeLength             int
	IdLength               int
	Id                     string
	Language               language
}

//This function reads the data from the card using the ReadBlock method and
//converts the read bytes into a NDEF struct.
func (r *MifareULReader) ReadNdef() (Ndef, error) {
	var ndef Ndef
	buf := new(bytes.Buffer)
	for i := UserBlockStartIndex; i <= UserBlockEndIndex; i++ {
		tempBuf, err := r.ReadBlock(i)
		if err != nil {
			return ndef, err

		}
		buf.Write(tempBuf)
	}
	temp := buf.Bytes()

	//The second block gives information if the card is read only or writable.
	data, err := r.ReadBlock(2)
	//This gives the Max Size of the user block that can be used to for user data.
	intval, _ := strconv.Atoi(strconv.FormatInt(int64(data[1]), 16))
	//This value holds the total length of the NDEF message.
	totalMessageLength := int(temp[1])
	//Slicing of the first two bytes since it is already processed and the following bytes is the actual NDEF message.
	//The first byte is discarded since it is always 3 for NDEF messages
	temp = temp[2:]

	ndef = Ndef{
		NdefData:           createNdefMessages(temp),
		MaxSize:            intval,
		TotalMessageLength: totalMessageLength,
		IsReadOnly:         !(int(data[2]) == 0 && int(data[3]) == 0)}
	return ndef, err

}

//The below methods are to check if a particular bit is set.Let us consider an example below.
//To check if the message begin bit is set we should check if the MSB is set to true.
//2^7=128 or 10000000 in binary
//So if we & 128 with 128 we get back 128.
//10000000 & 10000000 = 10000000
//So if we & 128 with some other number we always get 0
//10000000 & 01111111 = 00000000
func isMessageBeginSet(recordByte int) bool {
	return (recordByte & MB) == MB
}
func isMessageEndSet(recordByte int) bool {
	return (recordByte & ME) == ME
}
func isShortRecord(recordByte int) bool {
	return (recordByte & SR) == SR
}
func isChunkFlagSet(recordByte int) bool {
	return (recordByte & CF) == CF
}
func isIdLengthFieldPresent(recordByte int) bool {
	return (recordByte & IL) == IL
}

//TNF can take value from 0 to 7.It can take the below values.
//Empty=0x00
//NFC Forum well-known type [NFC RTD]=0x01
//Media-type as defined in RFC 2046 [RFC 2046]=0x02
//Absolute URI as defined in RFC 3986 [RFC 3986]=0x03
//NFC Forum external type [NFC RTD]=0x04
//Unknown=0x05
//Unchanged=0x06
//Reserved=0x07
func getTnf(recordByte int) tnf {
	return tnf(recordByte & 7)
}

//The below method processes the bytes read from the reader and returns an array of NdefRecords
func createNdefMessages(ndefBytes []byte) []NdefRecord {
	isEnd := false
	i := 0
	var ndefMessage []NdefRecord
	for !isEnd {
		isBegin := isMessageBeginSet(int(ndefBytes[messageOffsetIndex]))
		isEnd = isMessageEndSet(int(ndefBytes[messageOffsetIndex]))
		isChunkFlagSet := isChunkFlagSet(int(ndefBytes[messageOffsetIndex]))
		isShortRecord := isShortRecord(int(ndefBytes[messageOffsetIndex]))
		isIdLengthFieldPresent := isIdLengthFieldPresent(int(ndefBytes[messageOffsetIndex]))
		typeLength := int(ndefBytes[messageTypeLengthIndex])
		tnf := getTnf(int(ndefBytes[messageOffsetIndex]))
		payloadLength := int(ndefBytes[payloadLengthIndex])
		ndefRecord := NdefRecord{
			IsStartRecord:          isBegin,
			IsEndRecord:            isEnd,
			ChunkFlag:              isChunkFlagSet,
			IsShortRecord:          isShortRecord,
			IsIdLengthFieldPresent: isIdLengthFieldPresent,
			RecordNumber:           i,
			Tnf:                    tnf,
			PayloadLength:          payloadLength,
			TypeLength:             typeLength}
		//The below if block is executed if the IDLength field is present. Please see the ndef documentation for better understanding.
		if isIdLengthFieldPresent {
			tempPayloadTypeStartIndex := payloadTypeStartIndex + 1
			idLength := int(ndefBytes[payloadTypeStartIndex])
			payloadType := string(ndefBytes[tempPayloadTypeStartIndex : tempPayloadTypeStartIndex+typeLength])
			id := string(ndefBytes[idStartIndex : idStartIndex+idLength])
			payload := ndefBytes[tempPayloadTypeStartIndex+typeLength+idLength : tempPayloadTypeStartIndex+typeLength+idLength+payloadLength]
			if payloadType == "T" {
				languageLength := int(payload[0])
				ndefRecord.Language = language(payload[1 : languageLength+1])
			}
			ndefRecord.Payload = payload
			ndefRecord.Type = payloadType
			ndefRecord.IdLength = idLength
			ndefRecord.Id = id
			ndefMessage = append(ndefMessage, ndefRecord)
			ndefBytes = ndefBytes[payloadTypeStartIndex+typeLength+idLength+payloadLength:]

		} else {
			payloadType := string(ndefBytes[payloadTypeStartIndex : payloadTypeStartIndex+typeLength])
			payload := ndefBytes[payloadTypeStartIndex+typeLength : payloadTypeStartIndex+typeLength+payloadLength]
			if payloadType == "T" {
				languageLength := int(payload[0])
				ndefRecord.Language = language(payload[1 : languageLength+1])
			}
			ndefRecord.Payload = payload
			ndefRecord.Type = payloadType
			ndefMessage = append(ndefMessage, ndefRecord)
			ndefBytes = ndefBytes[payloadTypeStartIndex+typeLength+payloadLength:]
		}

		i++
	}
	return ndefMessage
}

// MifareULWriter implements the Writer interface for the Mifare Ultralight
// card.
type MifareULWriter struct {
}

// WriteBlock implements writing a block of bytes for the Mifare Ultralight card.
// idx defines the memory position / block number.
// See also the Writer interface.
func (w *MifareULWriter) WriteBlock(idx int, data []byte) error {
	libLock.Lock()
	defer libLock.Unlock()

	var buffer [MifareULBlockLength]byte
	// Only MifareULBlockLength bytes will be copied to buffer.
	copy(buffer[:], data)
	pbuffer := unsafe.Pointer(&buffer[0])
	// Hand over a void pointer of the array containing the bytes to a C
	// function, which will do the actual work.
	status := C.MifareUL_Write_Block(C.uint8_t(idx), pbuffer)
	if status != C.PH_ERR_SUCCESS {
		return createLibErr(int(status))
	}
	return nil
}

func (w *MifareULWriter) WriteNdef(ndef *Ndef) error {
	ndefBytes := createBytesFromNdefStruct(ndef)
	err := writeToDevice(ndefBytes, w)
	return err

}

func (w *MifareULWriter) WriteString(payload string, language language) error {
	ndef := createDefaultNdefForString(payload, language)
	return w.WriteNdef(&ndef)
}

type tnf int

const (
	Empty tnf = iota
	Nfcrtd
	Rfc2046
	Rfc3986
	Nfcrtdexternal
	Unknowntnf
	Unchanged
	Reserved
)

var tnfs = [...]string{
	"Empty",
	"Nfcrtd",
	"Rfc2046",
	"Rfc3986",
	"Nfcrtdexternal",
	"Unknown",
	"Unchanged",
	"Reserved",
}

func (tn tnf) String() string {
	return tnfs[tn]
}

type language string

const (
	En language = "en"
	De          = "de"
	It          = "it"
	Nl          = "nl"
	Fr          = "fr"
	Ru          = "ru"
	Kr          = "kr"
	Cn          = "cn"
	Uk          = "uk"
	Ca          = "ca"
	Es          = "es"
)

//This function creates the first byte of the NDEF based on the respective values in the struct.
//Let us consider an example of Message Begin i.e MB flag. If the MB flag is set we should be setting the MSB to true.
//10000000 -->This is the value in binay we need assuming all the other 7 bits are set to false(Obviously for this example)
//This is 128 in decimal system.
//If the MB flag is set to true we are | it with 128 so that MSB is set.
//00000000 | 10000000 --> 10000000
//So if we have to set MB flag we just have to | the existing bits with 128 to set it.
func getFirstByte(ndefRecord NdefRecord) byte {
	var b byte
	if ndefRecord.IsStartRecord {
		b = b | MB

	}
	if ndefRecord.IsEndRecord {
		b = b | ME
	}
	if ndefRecord.ChunkFlag {
		b = b | CF
	}
	if ndefRecord.IsShortRecord {
		b = b | SR
	}
	if ndefRecord.IsIdLengthFieldPresent {
		b = b | IL
	}
	b = b | byte(ndefRecord.Tnf)
	return b

}

//This function creates a byte slice from the Ndef struct
func createBytesFromNdefStruct(ndef *Ndef) []byte {
	buf := new(bytes.Buffer)
	for _, record := range ndef.NdefData {
		firstByte := getFirstByte(record)
		buf.WriteByte(firstByte)
		buf.WriteByte(byte(record.TypeLength))
		buf.WriteByte(byte(record.PayloadLength))
		buf.Write([]byte(record.Type))
		if record.Type == "T" && len(record.Language) > 0 {
			buf.WriteByte(byte(len(string(record.Language))))
			buf.Write([]byte(string(record.Language)))
		}
		buf.Write(record.Payload)
	}
	//This is the character fe in hex.This needs to be added to the end of NDEF messages
	buf.WriteByte(byte(254))
	return buf.Bytes()
}

//This function writes the bytes to the device
func writeToDevice(ndefBytes []byte, writer *MifareULWriter) error {
	var data = make([]byte, 4, 4)
	//These 2 bytes need total be added at the begining of an NDEF message.
	//The 0 byte should be set to 3 for NDEF messages
	data[0] = byte(3)
	//The 1 byte should be set to the length of the message proceeding this byte.
	//The 1 is subtracted to exclude the fe we added at the end.
	data[1] = byte(len(ndefBytes) - 1)
	//This is where the actual NDEF message begins.
	data[2] = ndefBytes[0]
	data[3] = ndefBytes[1]
	err := writer.WriteBlock(4, data)
	if err != nil {
		return err
	}
	ndefBytes = ndefBytes[2:]
	exitCondition := int(math.Ceil(float64(len(ndefBytes)) / float64(4)))
	for i := 5; i < 5+exitCondition; i++ {
		err = writer.WriteBlock(i, ndefBytes)
		bytesToWrite := int(math.Min(4, float64(len(ndefBytes))))
		ndefBytes = ndefBytes[bytesToWrite:]
	}
	return err
}

func createDefaultNdefForString(payload string, language language) Ndef {
	ndef := Ndef{
		NdefData: []NdefRecord{{IsStartRecord: true,
			IsEndRecord:            true,
			ChunkFlag:              false,
			IsShortRecord:          true,
			IsIdLengthFieldPresent: false,
			Tnf:        1,
			TypeLength: 1,
			Type:       "T",
			Language:   language,
			Payload:    []byte(payload),
			//This is to handle the language
			PayloadLength: (len(payload) + len(string(language)) + 1)}}}
	return ndef
}

// DeviceParams holds various parameters of a nfc card/tag.
type DeviceParams struct {
	// Select AcKnowledge. Can be -1 if not available.
	SAK int
	// Holds ATQA for Tech Type A and ATQB for Tech Type B
	// ATQA = Answer To reQuest for Type A
	// ATQB = Answer To reQuest for Type B
	// ATQ can be nil, if not available
	ATQ []byte
	// Unified IDentifier of the card/tag. Can be nil, if not available.
	UID []byte
	// TechType identifies the technology type, e.g. A, B, P2P, etc.
	TechType TechType
	// TagType identifies the tag type, i.e. 1, 2, etc.
	TagType TagType
	// DevType
	DevType DeviceType
}

// Device identifies a card/tag and contains parameters, such as the tag type
// or technology and Reader/Writer implementations for reading and writing
// bytes. Please note that Reader or Writer can be nil. This happens if an
// card/tag is identified for which no Reader/Writer implementation exists.
// Params will never be nil.
type Device struct {
	// Params holds general nfc parameters of the card/tag, such as SAK or UID.
	Params DeviceParams
	// Reader reads bytes from the card/tag. Can be nil if no specific
	// implementation exists for the card/tag.
	Reader Reader
	// Writer writes bytes to the card/tag. Can be nil if no specific
	// implementation exists for the card/tag.
	Writer Writer
}

// TagType defines the nfc tag type, such as 1, 2, 3, 4A etc.
type TagType int

// Tag types which are supported by this wrapper.
const (
	TagType1 TagType = 1 + iota
	TagType2
	TagType3
	TagType4A
	TagTypeP2P
	TagTypeNFCDEP4A
	TagTypeUndefined
)

var tagTypes = [...]string{
	"TagType1",
	"TagType2",
	"TagType3",
	"TagType4A",
	"TagTypeP2P",
	"TagTypeNFCDEP4A",
	"TagTypeUndefined",
}

// String returns the name of a TagType (emulating an enum).
func (tt TagType) String() string {
	return tagTypes[tt-1]
}

// TechType defines the technology of the card/tag, e.g. A, B, etc..
type TechType int

// Technology types which are supported by this wrapper.
const (
	TechA TechType = 1 + iota
	TechB
	TechF
	TechV15693T5T
	Tech18000p3m3EPCGen2
	TechUndefined
)

var techTypes = [...]string{
	"TechA",
	"TechB",
	"TechF",
	"TechV15693T5T",
	"Tech18000p3m3EPCGen2",
	"TechUndefined",
}

// String returns the name of a TechType (emulating an enum).
func (tt TechType) String() string {
	return techTypes[tt-1]
}

// DeviceType defines manufacturer specific cards/tags which are supported by
// this wrapper.
type DeviceType int

// All cards/tags which are supported by this wrapper.
const (
	// Mifare Ultralight card
	MifareUL DeviceType = 1 + iota
	Unknown
)

var devTypes = [...]string{
	"MifareUL",
	"Unknown",
}

// String returns the name of a DeviceType (emulating an enum).
func (dt DeviceType) String() string {
	return devTypes[dt-1]
}

// createDevParams basically calls a C function to collect all necessary
// parameters and converts them to a go struct. A DeviceParams struct is always
// returned.
func createDevParams() DeviceParams {
	C.NFCParams_Retrieve()
	var atq []byte
	var uid []byte
	if C.nfcParams.uidSize > 0 {
		uidSize := int(C.nfcParams.uidSize)
		uid = make([]byte, uidSize)
		for i := 0; i < uidSize; i++ {
			uid[i] = byte(C.nfcParams.uid[i])
		}
	}
	if C.nfcParams.atqSize > 0 {
		atqSize := int(C.nfcParams.atqSize)
		atq = make([]byte, atqSize)
		for i := 0; i < atqSize; i++ {
			atq[i] = byte(C.nfcParams.atq[i])
		}
	}
	return DeviceParams{
		SAK:      int(C.nfcParams.sak),
		ATQ:      atq,
		UID:      uid,
		TechType: TechType(C.nfcParams.techType),
		TagType:  TagType(C.nfcParams.tagType),
		DevType:  Unknown,
	}
}

// Init initializes the C library. It must be called once before other functions
// of this wrapper are used. A NxpError is returned in case of an error, otherwise
// nil is returned.
func Init() error {
	// After initialization the C library needs to be called from the same OS
	// thread
	runtime.LockOSThread()

	libLock.Lock()
	defer libLock.Unlock()

	C.Set_Interface_Link()
	C.Reset_reader_device()

	// Note: The calling OS thread will be stored by the C library for event
	// handling
	status := C.NfcRdLibInit()
	if status != C.PH_ERR_SUCCESS {
		return createLibErr(int(status))
	}

	C.Discover_Init()

	return nil
}

// DeInit cleans up the C library and should be called if the wrapper is not
// used anymore, e.g. at the end of the application.
func DeInit() {
	libLock.Lock()
	C.Cleanup_Interface_Link()
	libLock.Unlock()

	// The go runtime is now free to re-bind this goroutine to another OS
	// thread
	runtime.UnlockOSThread()
}

// Discover selects a card/tag that is in the range of the breakout board and
// returns a  pointer to a Device structure. If an error occurred while detecting
// and activating the card/tag, a NxpError along with a nil value for the Device
// is returned. Please note that not all cards/tags may be supported and thus
// the Reader or Writer variables within the Device struct may be nil. However,
// if a card/tag is found and no error occurred you'll always get a Device
// struct containing the card's/tag's parameters.
// Discover has an optional timeout parameter. The default value is 1000, which
// means Discover waits up to 1000ms to discover/detect a card/tag. During
// that time the caller will be blocked. The timeout value has to be of type int
// and will be interpreted as ms. So, if you want to give some more time for
// card/tag discovery, for instance 20s, you can call Discover(20000).
func Discover(args ...interface{}) (*Device, error) {
	libLock.Lock()
	defer libLock.Unlock()

	// Default timeout value is 1000ms.
	var timeoutVal int = 1000

	for _, arg := range args {
		switch t := arg.(type) {
		case int:
			timeoutVal = t
		default:
			panic("Discover: Unknown argument. Only int for timeout in ms is supported")
		}
	}

	log.WithField("timeout", timeoutVal).Debug("Discover")
	timeout := time.After(time.Duration(timeoutVal) * time.Millisecond)
	for {
		select {
		case <-timeout:
			log.Debug("Timeout occurred")
			return nil, createExtErr(TimeoutErr, "Discover: timeout occurred")
		default:
			res := C.Discover()
			if res == C.DR_FOUND {
				log.Debug("Detected and activated card/device")
				params := createDevParams()
				// We don't know about the Reader Writer yet, so pass nil for them.
				dev := &Device{params, nil, nil}
				if C.DetectMifare() == C.mifare_ultralight {
					// We detected a Mifare Ultralight card, so we set
					// Reader, Writer and Params accordingly.
					log.WithField("cardType", MifareUL.String()).Debug("Discover")
					dev.Params.DevType = MifareUL
					dev.Reader = &MifareULReader{}
					dev.Writer = &MifareULWriter{}
				}
				log.WithField("cardType", Unknown.String()).Debug("Discover")
				return dev, nil
			}
		}
	}
}
