package report

const (

	//XMLHead ...
	XMLHead = `<?xml version="1.0" encoding="utf-8"?>
	<?mso-application progid="Word.Document"?>

	<w:wordDocument xmlns:w="http://schemas.microsoft.com/office/word/2003/wordml" xmlns:aml="http://schemas.microsoft.com/aml/2001/core" xmlns:dt="uuid:C2F41010-65B3-11d1-A29F-00AA00C14882" xmlns:ve="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:wx="http://schemas.microsoft.com/office/word/2003/auxHint" xmlns:wsp="http://schemas.microsoft.com/office/word/2003/wordml/sp2" xmlns:sl="http://schemas.microsoft.com/schemaLibrary/2003/core" w:macrosPresent="no" w:embeddedObjPresent="no" w:ocxPresent="no" xml:space="preserve">
	  <w:ignoreSubtree w:val="http://schemas.microsoft.com/office/word/2003/wordml/sp2"/>
	  <o:DocumentProperties>
	    <o:Author>scnace</o:Author>
	    <o:LastAuthor>admin</o:LastAuthor>
	    <o:Revision>2</o:Revision>
	    <o:TotalTime>1</o:TotalTime>
	    <o:Created>2016-09-12T01:48:00Z</o:Created>
	    <o:LastSaved>2016-09-12T01:48:00Z</o:LastSaved>
	    <o:Pages>1</o:Pages>
	    <o:Words>1</o:Words>
	    <o:Characters>11</o:Characters>
	    <o:Lines>1</o:Lines>
	    <o:Paragraphs>1</o:Paragraphs>
	    <o:CharactersWithSpaces>11</o:CharactersWithSpaces>
	    <o:Version>12</o:Version>
	  </o:DocumentProperties>
	  <w:fonts>
	    <w:defaultFonts w:ascii="Calibri" w:fareast="宋体" w:h-ansi="Calibri" w:cs="Times New Roman"/>
	    <w:font w:name="Times New Roman">
	      <w:panose-1 w:val="02020603050405020304"/>
	      <w:charset w:val="00"/>
	      <w:family w:val="Roman"/>
	      <w:pitch w:val="variable"/>
	      <w:sig w:usb-0="E0002AFF" w:usb-1="C0007841" w:usb-2="00000009" w:usb-3="00000000" w:csb-0="000001FF" w:csb-1="00000000"/>
	    </w:font>
	    <w:font w:name="宋体">
	      <w:altName w:val="SimSun"/>
	      <w:panose-1 w:val="02010600030101010101"/>
	      <w:charset w:val="86"/>
	      <w:family w:val="auto"/>
	      <w:pitch w:val="variable"/>
	      <w:sig w:usb-0="00000003" w:usb-1="288F0000" w:usb-2="00000016" w:usb-3="00000000" w:csb-0="00040001" w:csb-1="00000000"/>
	    </w:font>
	    <w:font w:name="Cambria Math">
	      <w:panose-1 w:val="02040503050406030204"/>
	      <w:charset w:val="00"/>
	      <w:family w:val="Roman"/>
	      <w:pitch w:val="variable"/>
	      <w:sig w:usb-0="E00002FF" w:usb-1="420024FF" w:usb-2="00000000" w:usb-3="00000000" w:csb-0="0000019F" w:csb-1="00000000"/>
	    </w:font>
	    <w:font w:name="Calibri">
	      <w:panose-1 w:val="020F0502020204030204"/>
	      <w:charset w:val="00"/>
	      <w:family w:val="Swiss"/>
	      <w:pitch w:val="variable"/>
	      <w:sig w:usb-0="E00002FF" w:usb-1="4000ACFF" w:usb-2="00000001" w:usb-3="00000000" w:csb-0="0000019F" w:csb-1="00000000"/>
	    </w:font>
	    <w:font w:name="@宋体">
	      <w:panose-1 w:val="02010600030101010101"/>
	      <w:charset w:val="86"/>
	      <w:family w:val="auto"/>
	      <w:pitch w:val="variable"/>
	      <w:sig w:usb-0="00000003" w:usb-1="288F0000" w:usb-2="00000016" w:usb-3="00000000" w:csb-0="00040001" w:csb-1="00000000"/>
	    </w:font>
	  </w:fonts>
	  <w:styles>
	    <w:versionOfBuiltInStylenames w:val="7"/>
	    <w:latentStyles w:defLockedState="off" w:latentStyleCount="267">
	      <w:lsdException w:name="Normal"/>
	      <w:lsdException w:name="heading 1"/>
	      <w:lsdException w:name="heading 2"/>
	      <w:lsdException w:name="heading 3"/>
	      <w:lsdException w:name="heading 4"/>
	      <w:lsdException w:name="heading 5"/>
	      <w:lsdException w:name="heading 6"/>
	      <w:lsdException w:name="heading 7"/>
	      <w:lsdException w:name="heading 8"/>
	      <w:lsdException w:name="heading 9"/>
	      <w:lsdException w:name="toc 1"/>
	      <w:lsdException w:name="toc 2"/>
	      <w:lsdException w:name="toc 3"/>
	      <w:lsdException w:name="toc 4"/>
	      <w:lsdException w:name="toc 5"/>
	      <w:lsdException w:name="toc 6"/>
	      <w:lsdException w:name="toc 7"/>
	      <w:lsdException w:name="toc 8"/>
	      <w:lsdException w:name="toc 9"/>
	      <w:lsdException w:name="caption"/>
	      <w:lsdException w:name="Title"/>
	      <w:lsdException w:name="Default Paragraph Font"/>
	      <w:lsdException w:name="Subtitle"/>
	      <w:lsdException w:name="Strong"/>
	      <w:lsdException w:name="Emphasis"/>
	      <w:lsdException w:name="Table Grid"/>
	      <w:lsdException w:name="Placeholder Text"/>
	      <w:lsdException w:name="No Spacing"/>
	      <w:lsdException w:name="Light Shading"/>
	      <w:lsdException w:name="Light List"/>
	      <w:lsdException w:name="Light Grid"/>
	      <w:lsdException w:name="Medium Shading 1"/>
	      <w:lsdException w:name="Medium Shading 2"/>
	      <w:lsdException w:name="Medium List 1"/>
	      <w:lsdException w:name="Medium List 2"/>
	      <w:lsdException w:name="Medium Grid 1"/>
	      <w:lsdException w:name="Medium Grid 2"/>
	      <w:lsdException w:name="Medium Grid 3"/>
	      <w:lsdException w:name="Dark List"/>
	      <w:lsdException w:name="Colorful Shading"/>
	      <w:lsdException w:name="Colorful List"/>
	      <w:lsdException w:name="Colorful Grid"/>
	      <w:lsdException w:name="Light Shading Accent 1"/>
	      <w:lsdException w:name="Light List Accent 1"/>
	      <w:lsdException w:name="Light Grid Accent 1"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 1"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 1"/>
	      <w:lsdException w:name="Medium List 1 Accent 1"/>
	      <w:lsdException w:name="Revision"/>
	      <w:lsdException w:name="List Paragraph"/>
	      <w:lsdException w:name="Quote"/>
	      <w:lsdException w:name="Intense Quote"/>
	      <w:lsdException w:name="Medium List 2 Accent 1"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 1"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 1"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 1"/>
	      <w:lsdException w:name="Dark List Accent 1"/>
	      <w:lsdException w:name="Colorful Shading Accent 1"/>
	      <w:lsdException w:name="Colorful List Accent 1"/>
	      <w:lsdException w:name="Colorful Grid Accent 1"/>
	      <w:lsdException w:name="Light Shading Accent 2"/>
	      <w:lsdException w:name="Light List Accent 2"/>
	      <w:lsdException w:name="Light Grid Accent 2"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 2"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 2"/>
	      <w:lsdException w:name="Medium List 1 Accent 2"/>
	      <w:lsdException w:name="Medium List 2 Accent 2"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 2"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 2"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 2"/>
	      <w:lsdException w:name="Dark List Accent 2"/>
	      <w:lsdException w:name="Colorful Shading Accent 2"/>
	      <w:lsdException w:name="Colorful List Accent 2"/>
	      <w:lsdException w:name="Colorful Grid Accent 2"/>
	      <w:lsdException w:name="Light Shading Accent 3"/>
	      <w:lsdException w:name="Light List Accent 3"/>
	      <w:lsdException w:name="Light Grid Accent 3"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 3"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 3"/>
	      <w:lsdException w:name="Medium List 1 Accent 3"/>
	      <w:lsdException w:name="Medium List 2 Accent 3"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 3"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 3"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 3"/>
	      <w:lsdException w:name="Dark List Accent 3"/>
	      <w:lsdException w:name="Colorful Shading Accent 3"/>
	      <w:lsdException w:name="Colorful List Accent 3"/>
	      <w:lsdException w:name="Colorful Grid Accent 3"/>
	      <w:lsdException w:name="Light Shading Accent 4"/>
	      <w:lsdException w:name="Light List Accent 4"/>
	      <w:lsdException w:name="Light Grid Accent 4"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 4"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 4"/>
	      <w:lsdException w:name="Medium List 1 Accent 4"/>
	      <w:lsdException w:name="Medium List 2 Accent 4"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 4"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 4"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 4"/>
	      <w:lsdException w:name="Dark List Accent 4"/>
	      <w:lsdException w:name="Colorful Shading Accent 4"/>
	      <w:lsdException w:name="Colorful List Accent 4"/>
	      <w:lsdException w:name="Colorful Grid Accent 4"/>
	      <w:lsdException w:name="Light Shading Accent 5"/>
	      <w:lsdException w:name="Light List Accent 5"/>
	      <w:lsdException w:name="Light Grid Accent 5"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 5"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 5"/>
	      <w:lsdException w:name="Medium List 1 Accent 5"/>
	      <w:lsdException w:name="Medium List 2 Accent 5"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 5"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 5"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 5"/>
	      <w:lsdException w:name="Dark List Accent 5"/>
	      <w:lsdException w:name="Colorful Shading Accent 5"/>
	      <w:lsdException w:name="Colorful List Accent 5"/>
	      <w:lsdException w:name="Colorful Grid Accent 5"/>
	      <w:lsdException w:name="Light Shading Accent 6"/>
	      <w:lsdException w:name="Light List Accent 6"/>
	      <w:lsdException w:name="Light Grid Accent 6"/>
	      <w:lsdException w:name="Medium Shading 1 Accent 6"/>
	      <w:lsdException w:name="Medium Shading 2 Accent 6"/>
	      <w:lsdException w:name="Medium List 1 Accent 6"/>
	      <w:lsdException w:name="Medium List 2 Accent 6"/>
	      <w:lsdException w:name="Medium Grid 1 Accent 6"/>
	      <w:lsdException w:name="Medium Grid 2 Accent 6"/>
	      <w:lsdException w:name="Medium Grid 3 Accent 6"/>
	      <w:lsdException w:name="Dark List Accent 6"/>
	      <w:lsdException w:name="Colorful Shading Accent 6"/>
	      <w:lsdException w:name="Colorful List Accent 6"/>
	      <w:lsdException w:name="Colorful Grid Accent 6"/>
	      <w:lsdException w:name="Subtle Emphasis"/>
	      <w:lsdException w:name="Intense Emphasis"/>
	      <w:lsdException w:name="Subtle Reference"/>
	      <w:lsdException w:name="Intense Reference"/>
	      <w:lsdException w:name="Book Title"/>
	      <w:lsdException w:name="Bibliography"/>
	      <w:lsdException w:name="TOC Heading"/>
	    </w:latentStyles>
	    <w:style w:type="paragraph" w:default="on" w:styleId="a">
	      <w:name w:val="Normal"/>
	      <wx:uiName wx:val="正文"/>
	      <w:pPr>
	        <w:widowControl w:val="off"/>
	        <w:jc w:val="both"/>
	      </w:pPr>
	      <w:rPr>
	        <wx:font wx:val="Calibri"/>
	        <w:kern w:val="2"/>
	        <w:sz w:val="21"/>
	        <w:sz-cs w:val="22"/>
	        <w:lang w:val="EN-US" w:fareast="ZH-CN" w:bidi="AR-SA"/>
	      </w:rPr>
	    </w:style>
	    <w:style w:type="character" w:default="on" w:styleId="a0">
	      <w:name w:val="Default Paragraph Font"/>
	      <wx:uiName wx:val="默认段落字体"/>
	    </w:style>
	    <w:style w:type="table" w:default="on" w:styleId="a1">
	      <w:name w:val="Normal Table"/>
	      <wx:uiName wx:val="普通表格"/>
	      <w:rPr>
	        <wx:font wx:val="Calibri"/>
	        <w:lang w:val="EN-US" w:fareast="ZH-CN" w:bidi="AR-SA"/>
	      </w:rPr>
	      <w:tblPr>
	        <w:tblInd w:w="0" w:type="dxa"/>
	        <w:tblCellMar>
	          <w:top w:w="0" w:type="dxa"/>
	          <w:left w:w="108" w:type="dxa"/>
	          <w:bottom w:w="0" w:type="dxa"/>
	          <w:right w:w="108" w:type="dxa"/>
	        </w:tblCellMar>
	      </w:tblPr>
	    </w:style>
	    <w:style w:type="list" w:default="on" w:styleId="a2">
	      <w:name w:val="No List"/>
	      <wx:uiName wx:val="无列表"/>
	    </w:style>
	    <w:style w:type="paragraph" w:styleId="a3">
	      <w:name w:val="header"/>
	      <wx:uiName wx:val="页眉"/>
	      <w:basedOn w:val="a"/>
	      <w:link w:val="Char"/>
	      <w:rsid w:val="00C82541"/>
	      <w:pPr>
	        <w:pBdr>
	          <w:bottom w:val="single" w:sz="6" wx:bdrwidth="15" w:space="1" w:color="auto"/>
	        </w:pBdr>
	        <w:tabs>
	          <w:tab w:val="center" w:pos="4153"/>
	          <w:tab w:val="right" w:pos="8306"/>
	        </w:tabs>
	        <w:snapToGrid w:val="off"/>
	        <w:jc w:val="center"/>
	      </w:pPr>
	      <w:rPr>
	        <wx:font wx:val="Calibri"/>
	        <w:sz w:val="18"/>
	        <w:sz-cs w:val="18"/>
	      </w:rPr>
	    </w:style>
	    <w:style w:type="character" w:styleId="Char">
	      <w:name w:val="页眉 Char"/>
	      <w:basedOn w:val="a0"/>
	      <w:link w:val="a3"/>
	      <w:rsid w:val="00C82541"/>
	      <w:rPr>
	        <w:sz w:val="18"/>
	        <w:sz-cs w:val="18"/>
	      </w:rPr>
	    </w:style>
	    <w:style w:type="paragraph" w:styleId="a4">
	      <w:name w:val="footer"/>
	      <wx:uiName wx:val="页脚"/>
	      <w:basedOn w:val="a"/>
	      <w:link w:val="Char0"/>
	      <w:rsid w:val="00C82541"/>
	      <w:pPr>
	        <w:tabs>
	          <w:tab w:val="center" w:pos="4153"/>
	          <w:tab w:val="right" w:pos="8306"/>
	        </w:tabs>
	        <w:snapToGrid w:val="off"/>
	        <w:jc w:val="left"/>
	      </w:pPr>
	      <w:rPr>
	        <wx:font wx:val="Calibri"/>
	        <w:sz w:val="18"/>
	        <w:sz-cs w:val="18"/>
	      </w:rPr>
	    </w:style>
	    <w:style w:type="character" w:styleId="Char0">
	      <w:name w:val="页脚 Char"/>
	      <w:basedOn w:val="a0"/>
	      <w:link w:val="a4"/>
	      <w:rsid w:val="00C82541"/>
	      <w:rPr>
	        <w:sz w:val="18"/>
	        <w:sz-cs w:val="18"/>
	      </w:rPr>
	    </w:style>
	  </w:styles>
	  <w:shapeDefaults>
	    <o:shapedefaults v:ext="edit" spidmax="3074"/>
	    <o:shapelayout v:ext="edit">
	      <o:idmap v:ext="edit" data="2"/>
	    </o:shapelayout>
	  </w:shapeDefaults>
	  <w:docPr>
	    <w:view w:val="print"/>
	    <w:zoom w:percent="100"/>
	    <w:doNotEmbedSystemFonts/>
	    <w:bordersDontSurroundHeader/>
	    <w:bordersDontSurroundFooter/>
	    <w:defaultTabStop w:val="420"/>
	    <w:drawingGridVerticalSpacing w:val="156"/>
	    <w:displayHorizontalDrawingGridEvery w:val="0"/>
	    <w:displayVerticalDrawingGridEvery w:val="2"/>
	    <w:punctuationKerning/>
	    <w:characterSpacingControl w:val="CompressPunctuation"/>
	    <w:optimizeForBrowser/>
	    <w:validateAgainstSchema/>
	    <w:saveInvalidXML w:val="off"/>
	    <w:ignoreMixedContent w:val="off"/>
	    <w:alwaysShowPlaceholderText w:val="off"/>
	    <w:hdrShapeDefaults>
	      <o:shapedefaults v:ext="edit" spidmax="3074"/>
	    </w:hdrShapeDefaults>
	    <w:footnotePr>
	      <w:footnote w:type="separator">
	        <w:p wsp:rsidR="00BB52B0" wsp:rsidRDefault="00BB52B0" wsp:rsidP="00C82541">
	          <w:r>
	            <w:separator/>
	          </w:r>
	        </w:p>
	      </w:footnote>
	      <w:footnote w:type="continuation-separator">
	        <w:p wsp:rsidR="00BB52B0" wsp:rsidRDefault="00BB52B0" wsp:rsidP="00C82541">
	          <w:r>
	            <w:continuationSeparator/>
	          </w:r>
	        </w:p>
	      </w:footnote>
	    </w:footnotePr>
	    <w:endnotePr>
	      <w:endnote w:type="separator">
	        <w:p wsp:rsidR="00BB52B0" wsp:rsidRDefault="00BB52B0" wsp:rsidP="00C82541">
	          <w:r>
	            <w:separator/>
	          </w:r>
	        </w:p>
	      </w:endnote>
	      <w:endnote w:type="continuation-separator">
	        <w:p wsp:rsidR="00BB52B0" wsp:rsidRDefault="00BB52B0" wsp:rsidP="00C82541">
	          <w:r>
	            <w:continuationSeparator/>
	          </w:r>
	        </w:p>
	      </w:endnote>
	    </w:endnotePr>
	    <w:compat>
	      <w:spaceForUL/>
	      <w:balanceSingleByteDoubleByteWidth/>
	      <w:doNotLeaveBackslashAlone/>
	      <w:ulTrailSpace/>
	      <w:doNotExpandShiftReturn/>
	      <w:adjustLineHeightInTable/>
	      <w:breakWrappedTables/>
	      <w:snapToGridInCell/>
	      <w:wrapTextWithPunct/>
	      <w:useAsianBreakRules/>
	      <w:dontGrowAutofit/>
	      <w:useFELayout/>
	    </w:compat>
	    <wsp:rsids>
	      <wsp:rsidRoot wsp:val="00C82541"/>
	      <wsp:rsid wsp:val="00BB52B0"/>
	      <wsp:rsid wsp:val="00C82541"/>
	    </wsp:rsids>
	  </w:docPr>
	  <w:body>
	    <wx:sect>`
	//XMLTitle == 居中大标题
	XMLTitle = `<w:p>
		<w:pPr>
			<w:pStyle w:val="a2"/>
			<w:jc w:val="center"/>
		</w:pPr>
		<w:r>
			<w:t>%s</w:t>
		</w:r>
	</w:p>
`
	//XMLTitle1 == 标题1
	XMLTitle1 = `<w:p>
    <w:pPr>
      <w:pStyle w:val="a2"/>
    </w:pPr>
    <w:r>
      <w:t>%s</w:t>
    </w:r>
  </w:p>
`
	//XMLTitle2 == 标题2
	XMLTitle2 = `<w:p>
    <w:pPr>
      <w:pStyle w:val="a3"/>
    </w:pPr>
    <w:r>
      <w:t>%s</w:t>
    </w:r>
  </w:p>
`
	//XMLTitle2WithGrayBg == 灰色背景的标题2
	XMLTitle2WithGrayBg = `<w:p>
    <w:pPr>
      <w:pStyle w:val="a3"/>
			<w:shd w:val="clear" w:color="auto" w:fill="D4D8DA"/>
    </w:pPr>
    <w:r>
      <w:t>%s</w:t>
    </w:r>
  </w:p>
`
	//XMLTitle3 == 标题3
	XMLTitle3 = `<w:p>
    <w:pPr>
      <w:pStyle w:val="a4"/>
    </w:pPr>
    <w:r>
      <w:t>%s</w:t>
    </w:r>
  </w:p>
`
	//XMLText == 正文
	XMLText = `<w:p>
    <w:r>
      <w:t>%s</w:t>
    </w:r>
  </w:p>
`
	//XMLInlineText == 不换行的正文
	XMLInlineText = `<w:pPr>
		<w:pStyle w:val="a2"/>
	</w:pPr>
	<w:r>
		<w:t>%s</w:t>
	</w:r>
`
	//XMLTableHead ...
	XMLTableHead = `<w:tbl>
				<w:tblPr>
					<w:tblW w:w="9138" w:type="dxa"/>
					<w:tblBorders>
						<w:top w:val="thick-thin-medium-gap" w:sz="24" wx:bdrwidth="120" w:space="0" w:color="A5A5A5"/>
					</w:tblBorders>
				</w:tblPr>
`
	//XMLTableTR ...
	XMLTableTR = `<w:tr wsp:rsidR="00AF5A68" wsp:rsidTr="00AF5A68">
`
	//XMLTableTD ...
	XMLTableTD = `<w:tc>
  <w:tcPr>
    <w:tcW w:w="3046" w:type="dxa"/>
    <w:shd w:val="clear" w:color="auto" w:fill="auto"/>
  </w:tcPr>
  <w:p wsp:rsidR="00AF5A68" wsp:rsidRDefault="00AF5A68" wsp:rsidP="00AF5A68">
    <w:pPr>
      <w:tabs>
        <w:tab w:val="center" w:pos="1312"/>
      </w:tabs>
    </w:pPr>
`
	//XMLHeadTableTDBegin ...
	XMLHeadTableTDBegin = `<w:tc>
	<w:tcPr>
		<w:tcW w:w="3046" w:type="dxa"/>
		<w:shd w:val="clear" w:color="auto" w:fill="D4D8DA"/>
	</w:tcPr>
	<w:p wsp:rsidR="00AF5A68" wsp:rsidRDefault="00AF5A68" wsp:rsidP="00AF5A68">
		<w:pPr>
			<w:tabs>
				<w:tab w:val="center" w:pos="1312"/>
			</w:tabs>
		</w:pPr>
`
	//XMLHeadtableTDText ..
	XMLHeadtableTDText = `<w:r>
			<w:t>%s</w:t>
		</w:r>
`

	//XMLHeadTableTDEnd ...
	XMLHeadTableTDEnd = `
	</w:p>
</w:tc>
`
	//XMLTableEndTR ...
	XMLTableEndTR = `</w:tr>`
	//XMLTableFooter ...
	XMLTableFooter = `</w:tbl>`
	//XMLSectBegin ...
	XMLSectBegin = `<w:sectPr>`
	//XMLSectEnd ..
	XMLSectEnd = `<w:pgSz w:w="11906" w:h="16838"/>
	<w:pgMar w:top="1440" w:right="1800" w:bottom="1440" w:left="1800" w:header="851" w:footer="992" w:gutter="0"/>
	<w:cols w:space="425"/>
	<w:docGrid w:type="lines" w:line-pitch="312"/>
	    </w:sectPr>
`
	//XMLhdr ==页眉
	XMLhdr = `<w:hdr w:type="odd">
		<w:p>
			<w:pPr>
				<w:pStyle w:val="a4"/>
				<w:tabs>
					<w:tab w:val="center" w:pos="4153"/>
					<w:tab w:val="center" w:pos="8306"/>
				</w:tabs>
				<w:jc w:val="center"/>
			</w:pPr>
			<w:r>
				<w:t>%s</w:t>
			</w:r>
		</w:p>
	</w:hdr>
`
	//XMLftr 加上页码
	XMLftr = `<w:ftr w:type="odd">
	<w:p>
		<w:pPr>
			<w:pStyle w:val="a3"/>
			<w:tabs>
				<w:tab w:val="center" w:pos="4153"/>
				<w:tab w:val="right" w:pos="8306"/>
			</w:tabs>
		</w:pPr>
		<w:r>
			<w:rPr>
				<w:sz w:val="18"/>
			</w:rPr>
			<w:pict>
				<v:shape id="_x0000_s1026" o:spt="202" type="#_x0000_t202" style="position:absolute;left:0pt;margin-top:0pt;height:144pt;width:144pt;mso-position-horizontal:center;mso-position-horizontal-relative:margin;mso-wrap-style:none;z-index:251658240;mso-width-relative:page;mso-height-relative:page;" filled="f" stroked="f" coordsize="21600,21600">
					<v:path/>
					<v:fill on="f" focussize="0,0"/>
					<v:stroke on="f"/>
					<v:imagedata o:title=""/>
					<o:lock v:ext="edit" aspectratio="f"/>
					<v:textbox inset="0.00pt,0.00pt,0.00pt,0.00pt" style="mso-fit-shape-to-text:t;">
						<w:txbxContent>
							<w:p>
								<w:pPr>
									<w:snapToGrid w:val="off"/>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
								</w:pPr>
								<w:r>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
									<w:fldChar w:fldCharType="begin"/>
								</w:r>
								<w:r>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
									<w:instrText> PAGE  \* MERGEFORMAT </w:instrText>
								</w:r>
								<w:r>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
									<w:fldChar w:fldCharType="separate"/>
								</w:r>
								<w:r>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
									<w:t>1</w:t>
								</w:r>
								<w:r>
									<w:rPr>
										<w:sz w:val="18"/>
									</w:rPr>
									<w:fldChar w:fldCharType="end"/>
								</w:r>
							</w:p>
						</w:txbxContent>
					</v:textbox>
				</v:shape>
			</w:pict>
		</w:r>
	</w:p>
</w:ftr>
`
	//XMLIMGTitle ..
	XMLIMGTitle = `<w:p>
`
	//XMLImage 图片XML格式
	XMLImage = `<w:r>
			<w:pict>
					<w:binData w:name="%s">%s</w:binData>
						<v:shape id="_x0000_s1026" o:spt="75" alt="%s" type="#_x0000_t75" style="height:%spt;width:%spt;"
						filled="f" o:preferrelative="t" stroked="f" coordsize="%s,%s">
							<v:fill on="f" focussize="0,0"/>
							<v:stroke on="f"/>
							<v:imagedata src="%s" o:title="%s"/>
							<o:lock v:ext="edit" aspectratio="t"/>
							<w10:wrap type="none"/>
							<w10:anchorlock/>
						</v:shape>
			</w:pict>
	</w:r>
`
	//XMLIcon always is used for table data
	XMLIcon = `<w:r>
		<w:pict>
				<w:binData w:name="%s">%s</w:binData>
					<v:shape id="_x0000_s1026" o:spt="75" alt="%s" type="#_x0000_t75" style="height:16pt;width:16pt;"
					filled="f" o:preferrelative="t" stroked="f" coordsize="21600,21600">
						<v:fill on="f" focussize="0,0"/>
						<v:stroke on="f"/>
						<v:imagedata src="%s" o:title="%s"/>
						<o:lock v:ext="edit" aspectratio="t"/>
						<w10:wrap type="none"/>
						<w10:anchorlock/>
					</v:shape>
		</w:pict>
</w:r>
`
	//XMLIMGtail ...
	XMLIMGtail = `</w:p>
`
	//XMLBr == 换行
	XMLBr = `<w:p/>
`
	//XMLEndHead ...
	XMLEndHead = `</wx:sect>
  	</w:body>
  </w:wordDocument>
`
)
