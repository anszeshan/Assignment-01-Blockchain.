package assignment01bca

// Imports the crypto/sha256 and fmt packages
import (
	"crypto/sha256" // Used to calculate the hash of a block
	"fmt"           // Used for formatting and printing output
	"math/rand"
	"time"
)

// Defines a Block type, which has the following fields
type Block struct {
	Transaction  string // The transaction data
	Nonce        int    // A random number that is used to generate a unique hash for the block
	PreviousHash string // The hash of the previous block in the blockchain
	CurrentHash  string // The hash of the current block
}

// Defines a Blockchain type, which is a collection of blocks
type Blockchain struct {
	Blocks []Block
}

// Creates a new Block with the given transaction data, nonce, and previous hash. It then calculates the hash of the block and sets the CurrentHash field accordingly.
func NewBlock(transaction string, nonce int, previousHash string) *Block {

	block := &Block{
		Transaction: transaction, // The transaction data.
		// Nonce:        rand, // A random number that is used to generate a unique hash for the block
		PreviousHash: previousHash, // The hash of the previous block in the blockchain
	}
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time

	block.Nonce = rand.Intn(1000000)         // Generate a random number between 0 and 999,999
	block.CurrentHash = CalculateHash(block) // Calculates the hash of the block.

	return block
}

// Calculates the hash of a block by using the SHA-256 algorithm. The hash is calculated using the transaction data, nonce, and previous hash of the block.
func CalculateHash(block *Block) string {
	hash := sha256.Sum256([]byte(block.Transaction + fmt.Sprint(block.Nonce) + block.PreviousHash)) // Calculates the hash of the block using the SHA-256 algorithm.
	fmt.Printf("")
	return fmt.Sprintf("%x", hash) // Returns the hash of the block as a hexadecimal string.
}

// Adds a block to the blockchain.
func (blockchain *Blockchain) AddBlock(block *Block) {
	blockchain.Blocks = append(blockchain.Blocks, *block) // Appends the block to the blockchain.
}

// Displays all of the blocks in the blockchain to the console.
func (blockchain *Blockchain) DisplayBlocks() {
	fmt.Println("Blockchain Output :") // Prints a message to the console indicating the start of the blockchain output.

	// Iterates over all of the blocks in the blockchain.
	for _, block := range blockchain.Blocks {
		fmt.Printf("")
		fmt.Printf("Block : %d\n", block.Nonce)               // Prints the nonce of the current block.
		fmt.Printf("Transaction: %s\n", block.Transaction)    // Prints the transaction data of the current block
		fmt.Printf("Previous hash: %s\n", block.PreviousHash) // Prints the previous hash of the current block.
		fmt.Printf("Current hash: %s\n", block.CurrentHash)   // Prints the current hash of the current block.
		fmt.Printf("")
	}
}

// Verifies the integrity of the blockchain by checking that the hash of each block matches the previous hash of the blockchain.
func (blockchain *Blockchain) VerifyChain() bool {
	fmt.Printf("")
	fmt.Printf("----- Verifying The Chain -----") // Prints a message to the console indicating the start of the chain verification process.
	fmt.Printf("\n\n")

	// Iterates over all of the blocks in the blockchain, starting from the second block.
	for i := 1; i < len(blockchain.Blocks); i++ {
		block := blockchain.Blocks[i]
		previousBlock := blockchain.Blocks[i-1]
		// Checks if the previous hash of the current block matches the current hash of the previous block.
		if block.PreviousHash != previousBlock.CurrentHash {
			return false // If not, then the blockchain is invalid and the function returns false.
		}
	}

	return true // If all of the previous hash checks pass, then the blockchain is valid and the function returns true.
}

// Changes the transaction data in a block and then recalculates the hash of the block. It then verifies that the blockchain is still valid and returns an error if it is not.
func ChangeBlock(blockchain *Blockchain, blockIndex int, newTransaction string) error {
	block := blockchain.Blocks[blockIndex]    // Gets the block that we want to change
	block.Transaction = newTransaction        // Changes the transaction data in the block.
	block.CurrentHash = CalculateHash(&block) // Recalculates the hash of the block.
	blockchain.Blocks[blockIndex] = block     // Replaces the old block with the new block in the blockchain.
	// Verifies the integrity of the blockchain.
	if !blockchain.VerifyChain() {
		fmt.Printf("")
		return fmt.Errorf("Blockchain is invalid after changing block") // If the blockchain is invalid, then the function returns an error
	}

	return nil // If all of the checks pass, then the function returns nil.
}

// Creates a new Blockchain object.
func main() {
	blockchain := &Blockchain{}

	genesisBlock := NewBlock("Genesis block", 0, "")
	blockchain.AddBlock(genesisBlock)

	blockchain.AddBlock(NewBlock("Bob sent 10 coins to Alice", 1, genesisBlock.CurrentHash))
	blockchain.AddBlock(NewBlock("Alice sent 5 coins to Carol", 2, blockchain.Blocks[1].CurrentHash))

	blockchain.DisplayBlocks()

	err := ChangeBlock(blockchain, 1, "Alice sent 7 coins to Carol")
	if err != nil {
		fmt.Println(err)
	}
	blockchain.DisplayBlocks()

	if blockchain.VerifyChain() {
		fmt.Println("Blockchain is valid")
	} else {
		fmt.Println("Blockchain is invalid")
	}
}

//package changed from main to assignment01bca, that's why causing error if you named it main package then all working fine
