import XCTest
import SwiftTreeSitter
import TreeSitterGettext

final class TreeSitterGettextTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_po())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading PO grammar")
    }
}
